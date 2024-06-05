package generate

import (
	"log"
	"strings"

	"github.com/Nyarum/diho_bytes_generate/customtypes"

	"github.com/dave/jennifer/jen"
)

func GenerateEncodeForStruct(filename, pkg string, packetDescrs []customtypes.PacketDescr) {
	f := jen.NewFilePathName("", pkg)

	f.HeaderComment("Code generated by diho_bytes_generate " + filename + "; DO NOT EDIT.")

	for _, packetDescr := range packetDescrs {
		body := []jen.Code{
			jen.Id("newBuf").Op(":=").Qual("github.com/valyala/bytebufferpool", "Get").Call(),
			jen.Defer().Qual("github.com/valyala/bytebufferpool", "Put").Call(jen.Id("newBuf")),
			jen.Var().Id("err").Error(),
		}

		for _, field := range packetDescr.FieldsWithTypes.Keys() {
			fieldInfo, _ := packetDescr.FieldsWithTypes.Get(field)

			if !fieldInfo.IsArray {
				switch fieldInfo.TypeName {
				case "uint16", "uint32", "uint64", "uint8", "int16", "int32", "int64", "int8", "bool":
					body = append(body, []jen.Code{
						jen.Err().Op("=").Qual("encoding/binary", "Write").Call(jen.Id("newBuf"), jen.Id("endian"), jen.Id("p").Dot(field)),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						),
					}...)
				case "string":
					body = append(body, []jen.Code{
						jen.Err().Op("=").Qual("github.com/Nyarum/diho_bytes_generate/utils", "WriteStringNull").Call(jen.Id("newBuf"), jen.Id("p").Dot(field)),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						),
					}...)
				default:
					endianSwitch := jen.Id("endian")
					if fieldInfo.IsLittle {
						endianSwitch = jen.Qual("encoding/binary", "LittleEndian")
					}

					body = append(body, []jen.Code{
						jen.If(jen.List(jen.Id("encodeBuf"), jen.Err()).Op(":=").Id("p").Dot(field).Dot("Encode").Call(jen.Id("ctx"), endianSwitch),
							jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						).Else().Block(
							jen.Id("newBuf").Dot("Write").Call(jen.Id("encodeBuf")),
						),
					}...)
				}
			} else {

				switch fieldInfo.TypeName {
				case "uint16", "uint32", "uint64", "uint8", "int16", "int32", "int64", "int8", "bool":
					body = append(body, []jen.Code{
						jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("p").Dot(field)).Block(
							jen.If(jen.Err().Op("=").Qual("encoding/binary", "Write").Call(jen.Id("newBuf"), jen.Id("endian"), jen.Id("v")),
								jen.Err().Op("!=").Nil()).Block(
								jen.Return(jen.Nil(), jen.Err()),
							),
						),
					}...)
				case "byte":
					body = append(body, []jen.Code{
						jen.Err().Op("=").Qual("github.com/Nyarum/diho_bytes_generate/utils", "WriteBytes").Call(jen.Id("newBuf"), jen.Id("p").Dot(field)),
						jen.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Err()),
						),
					}...)
				default:
					body = append(body, []jen.Code{
						jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("p").Dot(field)).Block(
							jen.If(jen.List(jen.Id("encodeBuf"), jen.Err()).Op(":=").Id("v").Dot("Encode").Call(jen.Id("ctx"), jen.Id("endian")),
								jen.Err().Op("!=").Nil()).Block(
								jen.Return(jen.Nil(), jen.Err()),
							).Else().Block(
								jen.Id("newBuf").Dot("Write").Call(jen.Id("encodeBuf")),
							),
						),
					}...)
				}
			}

			if packetDescr.IsFilterMethod {
				body = append(body, []jen.Code{
					jen.If(jen.Id("p").Dot("Filter").Call(jen.Id("ctx"))).Op("==").Id("true").Block(
						jen.Return(
							jen.Qual("github.com/Nyarum/diho_bytes_generate/utils", "Clone").Call(jen.Id("newBuf")),
							jen.Nil(),
						),
					),
				}...)
			}
		}

		body = append(body, jen.Return(
			jen.Qual("github.com/Nyarum/diho_bytes_generate/utils", "Clone").Call(jen.Id("newBuf")),
			jen.Nil(),
		))

		f.Func().Params(jen.Id("p").Op("*").Id(packetDescr.StructName)).Id("Encode").Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("endian").Qual("encoding/binary", "ByteOrder"),
		).Params(
			jen.Index().Byte(), jen.Error(),
		).Block(body...)
	}

	outputFilename := strings.TrimSuffix(filename, ".go") + "_encode.gen.go"
	if err := f.Save(outputFilename); err != nil {
		log.Fatalf("Failed to save file: %s", err)
	}
}
