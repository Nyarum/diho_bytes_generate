package customtypes

import "github.com/elliotchance/orderedmap/v2"

type Field struct {
	IsArray  bool
	TypeName string
	IsLittle bool
}

type PacketDescr struct {
	PackageName     string
	StructName      string
	FieldsWithTypes *orderedmap.OrderedMap[string, Field]
	IsFilterMethod  bool
}
