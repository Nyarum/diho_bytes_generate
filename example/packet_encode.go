package example

import (
	utils "bytes_generated/utils"
	"encoding/binary"
	bytebufferpool "github.com/valyala/bytebufferpool"
)

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
