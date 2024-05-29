package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/Nyarum/diho_bytes_generate/customtypes"

	"github.com/elliotchance/orderedmap/v2"
)

func ParseBinaryFile(filename string) customtypes.PacketDescr {
	// Parse the Go source file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse file: %s", err)
	}

	packetDescr := customtypes.PacketDescr{
		FieldsWithTypes: orderedmap.NewOrderedMap[string, string](),
	}

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

			packetDescr.StructName = typeSpec.Name.Name

			if v := typeSpec.Type.(*ast.StructType); v != nil {
				packetDescr.StructName = typeSpec.Name.Name

				for _, field := range v.Fields.List {
					if field.Tag != nil && strings.Contains(field.Tag.Value, "ignore") {
						continue
					}

					if v, ok := field.Type.(*ast.Ident); ok {
						packetDescr.FieldsWithTypes.Set(field.Names[0].Name, v.Name)
					}

					if v, ok := field.Type.(*ast.ArrayType); ok {
						if v, ok := v.Elt.(*ast.Ident); ok {
							if v.Name == "byte" {
								packetDescr.FieldsWithTypes.Set(field.Names[0].Name, "[]byte")
							}
						}
					}
				}
			}
		}
	}

	return packetDescr
}
