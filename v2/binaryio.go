package binaryio

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
)

// SECTION_TYPE:FileReader
type FileReader struct {
	File   *os.File
	Reader *bufio.Reader
	buf_2  []byte
	buf_4  []byte
	buf_8  []byte
}

func (FileReader *FileReader) Init() {
	FileReader.buf_2 = make([]byte, 2)
	FileReader.buf_4 = make([]byte, 4)
	FileReader.buf_8 = make([]byte, 8)
}

func (FileReader *FileReader) OpenFile(name string, flat int, perm fs.FileMode) error {
	var err error

	FileReader.File, err = os.OpenFile(name, flat, perm)

	if err != nil {
		return err
	}

	FileReader.Reader = bufio.NewReader(FileReader.File)
	return err
}

func (FileReader *FileReader) Available() bool {
	_, err := FileReader.Reader.Peek(1)
	return err == nil
}

// 1 bytes 8 bits
func (FileReader *FileReader) Read() byte {
	data, _ := FileReader.Reader.ReadByte()

	return data
}

func (FileReader *FileReader) ReadUint8() uint8 {
	data, _ := FileReader.Reader.ReadByte()

	return data
}

// 2 bytes 16 bits
func (FileReader *FileReader) ReadUint16() uint16 {

	FileReader.Reader.Read(FileReader.buf_2)

	return uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])
}

func (FileReader *FileReader) ReadUint16Array() []uint16 {
	FileReader.Reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*2)

	FileReader.Reader.Read(buf)

	array := make([]uint16, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+2 {
		array[i] = uint16(buf[j])<<8 + uint16(buf[j+1])
	}

	return array
}

// 4 bytes 32 bits
func (FileReader *FileReader) ReadInt() int {

	FileReader.Reader.Read(FileReader.buf_4)

	return int(FileReader.buf_4[0])<<24 + int(FileReader.buf_4[1])<<16 + int(FileReader.buf_4[2])<<8 + int(FileReader.buf_4[3])
}

func (FileReader *FileReader) ReadUint32() uint32 {

	FileReader.Reader.Read(FileReader.buf_4)

	return uint32(FileReader.buf_4[0])<<24 + uint32(FileReader.buf_4[1])<<16 + uint32(FileReader.buf_4[2])<<8 + uint32(FileReader.buf_4[3])
}

func (FileReader *FileReader) ReadUint64() uint64 {

	FileReader.Reader.Read(FileReader.buf_8)

	return uint64(FileReader.buf_8[0])<<56 + uint64(FileReader.buf_8[1])<<48 + uint64(FileReader.buf_8[2])<<40 + uint64(FileReader.buf_8[3])<<32 + uint64(FileReader.buf_8[4])<<24 + uint64(FileReader.buf_8[5])<<16 + uint64(FileReader.buf_8[6])<<8 + uint64(FileReader.buf_8[7])
}

// string
func (FileReader *FileReader) ReadString() string {
	data, _ := FileReader.Reader.ReadBytes(0)

	return string(data[:len(data)-1])
}

// array
func (FileReader *FileReader) ReadIntArray() []int {
	data, _ := FileReader.Reader.ReadBytes(0)

	fmt.Println(data)

	FileReader.Reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*4)

	FileReader.Reader.Read(buf)

	array := make([]int, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = int(buf[j])<<24 + int(buf[j+1])<<16 + int(buf[j+2])<<8 + int(buf[j+3])
	}

	return array
}

func (FileReader *FileReader) ReadUint32Array() []uint32 {
	FileReader.Reader.Read(FileReader.buf_2)

	size := uint16(FileReader.buf_2[0])<<8 + uint16(FileReader.buf_2[1])

	buf := make([]byte, size*4)

	FileReader.Reader.Read(buf)

	array := make([]uint32, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = uint32(buf[j])<<24 + uint32(buf[j+1])<<16 + uint32(buf[j+2])<<8 + uint32(buf[j+3])
	}

	return array
}

func (FileReader *FileReader) Close() {
	// Release
	FileReader.buf_2 = nil
	FileReader.buf_4 = nil
	FileReader.buf_8 = nil

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
	// binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	// FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint16(FileWriter.Buffer_p, k)

		FileWriter.Buffer_p = FileWriter.Buffer_p[2:]
	}

	FileWriter.Buffer_p[0] = 0x000
	FileWriter.Buffer_p = FileWriter.Buffer_p[1:]
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

func (FileWriter *FileWriter) WriteUint64(data uint64) {
	binary.BigEndian.PutUint64(FileWriter.Buffer_p, data)

	FileWriter.Buffer_p = FileWriter.Buffer_p[8:]
}

// string
func (FileWriter *FileWriter) WriteString(data string) {
	tmp := append(FileWriter.Buffer_p[0:0], data...)

	FileWriter.Buffer_p[len(tmp)] = 0x000

	FileWriter.Buffer_p = FileWriter.Buffer_p[len(tmp)+1:]
}

// array
func (FileWriter *FileWriter) WriteIntArray(data []int) {
	// binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	// FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(FileWriter.Buffer_p, uint32(k))

		FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
	}

	FileWriter.Buffer_p[0] = 0x000
	FileWriter.Buffer_p = FileWriter.Buffer_p[1:]
}

func (FileWriter *FileWriter) WriteUint32Array(data []uint32) {
	// binary.BigEndian.PutUint16(FileWriter.Buffer_p, uint16(len(data)))

	// FileWriter.Buffer_p = FileWriter.Buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(FileWriter.Buffer_p, k)

		FileWriter.Buffer_p = FileWriter.Buffer_p[4:]
	}

	FileWriter.Buffer_p[0] = 0x000
	FileWriter.Buffer_p = FileWriter.Buffer_p[1:]
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

	// Release
	FileWriter.Buffer_m = nil
	FileWriter.Buffer_p = nil

	FileWriter.File.Close()
}
