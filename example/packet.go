package example

//go:generate diho_bytes_generate packet.go example
type Packet struct {
	ID   uint16
	Name string
	Bro  uint16
	Bro2 uint16
}
