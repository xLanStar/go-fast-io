package binaryio

import (
	"bufio"
	"encoding/binary"
	"io/fs"

	// Media "media"
	"os"
)

// SECTION_TYPE:FileReader
type FileReader struct {
	File   *os.File
	reader *bufio.Reader
	buf_2  []byte
	buf_4  []byte
}

func (FileReader *FileReader) Init() {
	FileReader.buf_2 = make([]byte, 2)
	FileReader.buf_4 = make([]byte, 4)
}

func (FileReader *FileReader) OpenFile(name string, flat int, perm fs.FileMode) error {
	var err error

	FileReader.File, err = os.OpenFile(name, flat, perm)

	if err != nil {
		return err
	}

	FileReader.reader = bufio.NewReader(FileReader.File)
	return err
}

func (FileReader *FileReader) Available() bool {
	_, err := FileReader.reader.Peek(1)
	return err == nil
}

// 1 bytes 8 bits
func (FileReader *FileReader) Read() byte {
	data, _ := FileReader.reader.ReadByte()

	return data
}

func (FileReader *FileReader) ReadUint8() uint8 {
	data, _ := FileReader.reader.ReadByte()

	return data
}

// 2 bytes 16 bits
func (FileReader *FileReader) ReadUint16() uint16 {

	FileReader.reader.Read(FileReader.buf_2)

	return uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])
}

func (FileReader *FileReader) ReadUint16Array() []uint16 {
	FileReader.reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*2)

	FileReader.reader.Read(buf)

	array := make([]uint16, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+2 {
		array[i] = uint16(buf[j])<<8 + uint16(buf[j+1])
	}

	return array
}

// 4 bytes 32 bits
func (FileReader *FileReader) ReadInt() int {

	FileReader.reader.Read(FileReader.buf_4)

	return int(FileReader.buf_4[0])<<24 + int(FileReader.buf_4[1])<<16 + int(FileReader.buf_4[2])<<8 + int(FileReader.buf_4[3])
}

func (FileReader *FileReader) ReadUint32() uint32 {

	FileReader.reader.Read(FileReader.buf_4)

	return uint32(FileReader.buf_4[0])<<24 + uint32(FileReader.buf_4[1])<<16 + uint32(FileReader.buf_4[2])<<8 + uint32(FileReader.buf_4[3])
}

// string
func (FileReader *FileReader) ReadString() string {
	data, _ := FileReader.reader.ReadBytes(0)

	return string(data[:len(data)-1])
}

// array
func (FileReader *FileReader) ReadIntArray() []int {
	FileReader.reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*4)

	FileReader.reader.Read(buf)

	array := make([]int, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = int(buf[j])<<24 + int(buf[j+1])<<16 + int(buf[j+2])<<8 + int(buf[j+3])
	}

	return array
}

func (FileReader *FileReader) ReadUint32Array() []uint32 {
	FileReader.reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*4)

	FileReader.reader.Read(buf)

	array := make([]uint32, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = uint32(buf[j])<<24 + uint32(buf[j+1])<<16 + uint32(buf[j+2])<<8 + uint32(buf[j+3])
	}

	return array
}

func (FileReader *FileReader) Close() {
	FileReader.File.Close()
}

// SECTION_TYPE:FileWriter
type FileWriter struct {
	File     *os.File
	Writer   *bufio.Writer
	Buffer_m []byte
	Buffer_p []byte
}

func (FileWriter *FileWriter) Init() {
	FileWriter.Buffer_m = make([]byte, 1<<20)
	FileWriter.Buffer_p = FileWriter.Buffer_m[:]
}

func (FileWriter *FileWriter) OpenFile(name string, flat int, perm fs.FileMode) error {
	var err error
	FileWriter.File, err = os.OpenFile(name, flat, perm)

	if err != nil {
		return err
	}

	FileWriter.Writer = bufio.NewWriter(FileWriter.File)

	FileWriter.Buffer_p = FileWriter.Buffer_m[:]
	return nil

}

// 1 byte 8 bits
func (FileWriter *FileWriter) Write(data byte) {
	FileWriter.Buffer_p[0] = data

	FileWriter.Buffer_p = FileWriter.Buffer_p[1:]
}
func (FileWriter *FileWriter) WriteUint8(data uint8) {
	FileWriter.Buffer_p[0] = data

	FileWriter.Buffer_p = FileWriter.Buffer_p[1:]
}

// 2 bytes 16 bits
func (FileWriter *FileWriter) WriteUint16(data uint16) {
	binary.BigEndian.PutUint16(FileWriter.Buffer_p, data)

	FileWriter.Buffer_p = FileWriter.Buffer_p[2:]
}

func (FileWriter *FileWriter) WriteUint16Array(data []uint16) {
	binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint16(FileWriter.Buffer_p, k)

		FileWriter.Buffer_p = FileWriter.Buffer_p[2:]
	}
}

// 4 bytes 32 bits
func (FileWriter *FileWriter) WriteInt(data int) {
	binary.BigEndian.PutUint32(FileWriter.Buffer_p, uint32(data))

	FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
}

func (FileWriter *FileWriter) WriteUint32(data uint32) {
	binary.BigEndian.PutUint32(FileWriter.Buffer_p, data)

	FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
}

// string
func (FileWriter *FileWriter) WriteString(data string) {
	tmp := append(FileWriter.Buffer_p[0:0], data...)

	FileWriter.Buffer_p[len(tmp)] = '\x00'

	FileWriter.Buffer_p = FileWriter.Buffer_p[len(tmp)+1:]
}

// array
func (FileWriter *FileWriter) WriteIntArray(data []int) {
	binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(FileWriter.Buffer_p, uint32(k))

		FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
	}
}

func (FileWriter *FileWriter) WriteUint32Array(data []uint32) {
	binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(FileWriter.Buffer_p, k)

		FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
	}
}

func (FileWriter *FileWriter) Flush() {
	FileWriter.Writer.Write(FileWriter.Buffer_m[:cap(FileWriter.Buffer_m)-cap(FileWriter.Buffer_p)])

	FileWriter.Writer.Flush()

	FileWriter.Buffer_p = FileWriter.Buffer_m[:]
}

func (FileWriter *FileWriter) Close() {
	if cap(FileWriter.Buffer_m) != cap(FileWriter.Buffer_p) {
		FileWriter.Flush()
	}

	FileWriter.File.Close()
}
