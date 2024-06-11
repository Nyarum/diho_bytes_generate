package customtypes

import "github.com/elliotchance/orderedmap/v2"

type CompositeIf struct {
	Field string
	Eq    string
}

type Field struct {
	IsArray     bool
	TypeName    string
	IsLittle    bool
	CompositeIf map[string]CompositeIf
}

type PacketDescr struct {
	PackageName     string
	StructName      string
	FieldsWithTypes *orderedmap.OrderedMap[string, Field]
	IsFilterMethod  bool
}
