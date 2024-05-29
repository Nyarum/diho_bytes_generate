package example

type Header struct {
}

//go:generate diho_bytes_generate packet.go example
type Packet struct {
	Header `dbg:"ignore"`
	ID     uint16
	Name   string
	Level  uint32
	Health uint8
	Mana   uint16
	Bro    int8
	Bro2   int16
	Bro3   int32
	Bro4   int64
}
