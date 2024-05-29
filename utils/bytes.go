package utils

import (
	"encoding/binary"
	"io"

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

func WriteBytes(buf *bytebufferpool.ByteBuffer, data []byte) error {
	err := binary.Write(buf, binary.LittleEndian, uint16(len(data))+1)
	if err != nil {
		return err
	}

	_, err = buf.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadStringNull(reader io.Reader) (string, error) {
	var ln uint16
	err := binary.Read(reader, binary.LittleEndian, &ln)
	if err != nil {
		return "", err
	}

	buf := make([]byte, ln)
	_, err = reader.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:ln-1]), nil
}

func ReadBytes(reader io.Reader) ([]byte, error) {
	var ln uint16
	err := binary.Read(reader, binary.LittleEndian, &ln)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, ln)
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:ln], nil
}

func Clone(buf *bytebufferpool.ByteBuffer) []byte {
	return append([]byte{}, buf.B...)
}
