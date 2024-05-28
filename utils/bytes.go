package utils

import (
	"encoding/binary"

	"github.com/valyala/bytebufferpool"
)

func WriteStringNull(buf *bytebufferpool.ByteBuffer, str string) error {
	err := binary.Write(buf, binary.LittleEndian, uint16(len(str))+1)
	if err != nil {
		return err
	}

	_, err = buf.WriteString(str)
	if err != nil {
		return err
	}

	err = buf.WriteByte(0x00)
	if err != nil {
		return err
	}

	return nil
}

func Clone(buf *bytebufferpool.ByteBuffer) []byte {
	return append([]byte{}, buf.B...)
}
