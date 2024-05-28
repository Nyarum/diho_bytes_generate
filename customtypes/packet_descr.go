package customtypes

import "github.com/elliotchance/orderedmap/v2"

type PacketDescr struct {
	StructName      string
	FieldsWithTypes *orderedmap.OrderedMap[string, string]
}
