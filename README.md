
Code-generation for produce encode / decode methods (to avoid using reflection on custom binary protocols)

To use this:
- go install github.com/Nyarum/diho_bytes_generate

add "go:generate" comment to your struct
```
//go:generate diho_bytes_generate packet.go example
type Packet struct {
	ID   uint16
	Name string
	Bro  uint16
	Bro2 uint16
}
```

it will produce generated code in the same folder that your struct
```
func (p *Packet) Encode(endian binary.ByteOrder) ([]byte, error) {
	newBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(newBuf)
	var err error
	err = binary.Write(newBuf, endian, p.ID)
	if err != nil {
		return nil, err
	}
	err = utils.WriteStringNull(newBuf, p.Name)
	if err != nil {
		return nil, err
	}
	err = binary.Write(newBuf, endian, p.Bro)
	if err != nil {
		return nil, err
	}
	err = binary.Write(newBuf, endian, p.Bro2)
	if err != nil {
		return nil, err
	}
	return utils.Clone(newBuf), nil
}
```