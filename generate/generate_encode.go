package generate

import (
	"bytes_generated/customtypes"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
)

func GenerateEncodeForStruct(filename, pkg string, packetDescr customtypes.PacketDescr) {
	f := jen.NewFilePathName("", pkg)

	body := []jen.Code{
		jen.Id("newBuf").Op(":=").Qual("github.com/valyala/bytebufferpool", "Get").Call(),
		jen.Defer().Qual("github.com/valyala/bytebufferpool", "Put").Call(jen.Id("newBuf")),
		jen.Var().Id("err").Error(),
	}

	for _, field := range packetDescr.FieldsWithTypes.Keys() {
		fieldType, _ := packetDescr.FieldsWithTypes.Get(field)

		switch fieldType {
		case "uint16":
			body = append(body, []jen.Code{
				jen.Err().Op("=").Qual("encoding/binary", "Write").Call(jen.Id("newBuf"), jen.Id("endian"), jen.Id("p").Dot(field)),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Return(jen.Nil(), jen.Err()),
				),
			}...)
		case "string":
			body = append(body, []jen.Code{
				jen.Err().Op("=").Qual("bytes_generated/utils", "WriteStringNull").Call(jen.Id("newBuf"), jen.Id("p").Dot(field)),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Return(jen.Nil(), jen.Err()),
				),
			}...)
		}
	}

	body = append(body, jen.Return(
		jen.Qual("bytes_generated/utils", "Clone").Call(jen.Id("newBuf")),
		jen.Nil(),
	))

	f.Func().Params(jen.Id("p").Op("*").Id(packetDescr.StructName)).Id("Encode").Params(
		jen.Id("endian").Qual("encoding/binary", "ByteOrder"),
	).Params(
		jen.Index().Byte(), jen.Error(),
	).Block(body...)

	outputFilename := strings.TrimSuffix(filename, ".go") + "_encode.go"
	if err := f.Save(outputFilename); err != nil {
		log.Fatalf("Failed to save file: %s", err)
	}
}
