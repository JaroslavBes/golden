package packet

import (
	"io"
//	"log"
	"encoding/binary"
)

type BinaryWriter struct {
	writer io.Writer
	offset int64
}

func NewBinaryWriter(writer io.Writer) (*BinaryWriter, error) {
	bw := new(BinaryWriter)
	bw.writer = writer
	bw.offset = 0
	return bw, nil
}

func (self *BinaryWriter) Offset() int64 {
	return self.offset
}

func (self *BinaryWriter) WriteUINT8(value uint8) (error) {
	err := binary.Write(self.writer, binary.LittleEndian, &value)
	self.offset += 1
	return err
}

func (self *BinaryWriter) WriteUINT16(value uint16) (error) {
	err := binary.Write(self.writer, binary.LittleEndian, value)
	self.offset += 2
	return err
}

func (self *BinaryWriter) WriteBytes(msg []byte) (error) {
	size := len(msg)
	for j := 0; j < size; j++ {
		var i byte = msg[j]
		err := binary.Write(self.writer, binary.LittleEndian, &i)
		if err != nil {
			return err
		}
		self.offset += 1
	}
	return nil
}

func (self *BinaryWriter) WriteZString(msg []byte) (error) {
	/* ... */
	size := len(msg)
	for j := 0; j < size; j++ {
		var i1 byte = msg[j]
		err1 := binary.Write(self.writer, binary.LittleEndian, &i1)
		if err1 != nil {
			return err1
		}
		self.offset += 1
	}
	/* Write ZERO byte */
	var i2 byte = '\x00'
	err2 := binary.Write(self.writer, binary.LittleEndian, &i2)
	if err2 != nil {
		return err2
	}
	self.offset += 1
	/* Done */
	return nil
}

func (self *BinaryWriter) Close() {
}