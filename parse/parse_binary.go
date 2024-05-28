package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"

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
					packetDescr.FieldsWithTypes.Set(field.Names[0].Name, field.Type.(*ast.Ident).Name)
				}
			}
		}
	}

	return packetDescr
}
