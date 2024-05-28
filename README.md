## Code Generation for Encode/Decode Methods of binary protocol

Generate encode and decode methods for your structs to avoid using reflection with custom binary protocols.

### How to Use

1. Install the code generation tool:
   ```sh
   go install github.com/Nyarum/diho_bytes_generate
   ```

2. Add a `go:generate` comment to your struct definition:
   ```go
   //go:generate diho_bytes_generate packet.go example
   type Packet struct {
       ID   uint16
       Name string
       Bro  uint16
       Bro2 uint16
   }
   ```

3. Run the `go generate` command in your terminal:
   ```sh
   go generate ./...
   ```

This will produce the generated code in the same folder as your struct.

### Example of Generated Code

The generated code will look something like this:

```go
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

This `Encode` method writes the struct fields into a byte buffer using the specified endianness and returns the encoded byte slice.
