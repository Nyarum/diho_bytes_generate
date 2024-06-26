package example

import "context"

type Header struct {
	Pass uint8
}

type InternalStruct struct {
	Test uint16
}

//go:generate diho_bytes_generate packet.go
type Packet struct {
	Header                  `dbg:"ignore,little"`
	ID                      uint16
	Name                    string
	Level                   uint32
	Health                  uint8
	Mana                    uint16
	Bro                     int8
	Bro2                    int16
	Bro3                    int32
	Bro4                    int64
	BytesField              []byte
	InternalStruct          InternalStruct `dbg:"little"`
	ArrayTest               [5]uint16
	SliceTest               []uint16
	InternalStructs         [2]InternalStruct
	IsIncludeByAnotherField uint32 `dbg:"Bro4==0,Bro3==1"`
}

func (p *Packet) Filter(ctx context.Context, name string) bool {

	return false
}
