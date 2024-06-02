package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/Nyarum/diho_bytes_generate/customtypes"
	"github.com/davecgh/go-spew/spew"

	"github.com/elliotchance/orderedmap/v2"
)

func isNormalType(t string) bool {
	switch t {
	case "uint16", "uint32", "uint64", "uint8", "int16", "int32", "int64", "int8", "string", "byte":
		return true
	default:
		return false
	}
}

func ParseBinaryFile(filename string) (pkgName string, packetsDescrs []customtypes.PacketDescr) {
	// Parse the Go source file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse file: %s", err)
	}

	pkgName = node.Name.Name

	packetsDescrs = make([]customtypes.PacketDescr, 0)

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			packetDescr := customtypes.PacketDescr{
				FieldsWithTypes: orderedmap.NewOrderedMap[string, customtypes.Field](),
				PackageName:     node.Name.Name,
			}

			packetDescr.StructName = typeSpec.Name.Name

			if v := typeSpec.Type.(*ast.StructType); v != nil {
				packetDescr.StructName = typeSpec.Name.Name

				for _, field := range v.Fields.List {
					spew.Dump(field)

					if field.Tag != nil && strings.Contains(field.Tag.Value, "ignore") {
						continue
					}

					if v, ok := field.Type.(*ast.Ident); ok {
						packetDescr.FieldsWithTypes.Set(field.Names[0].Name, customtypes.Field{
							TypeName: v.Name,
						})
					}

					if v, ok := field.Type.(*ast.ArrayType); ok {
						if v, ok := v.Elt.(*ast.Ident); ok {
							packetDescr.FieldsWithTypes.Set(field.Names[0].Name, customtypes.Field{
								IsArray:  true,
								TypeName: v.Name,
							})
						}
					}
				}
			}

			packetsDescrs = append(packetsDescrs, packetDescr)
		}
	}

	return
}
