package customtypes

import "github.com/elliotchance/orderedmap/v2"

type Field struct {
	IsArray     bool
	TypeName    string
	IsLittle    bool
	CompositeIf map[string]string
}

type PacketDescr struct {
	PackageName     string
	StructName      string
	FieldsWithTypes *orderedmap.OrderedMap[string, Field]
	IsFilterMethod  bool
}
