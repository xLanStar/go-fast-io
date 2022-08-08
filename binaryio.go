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
	file   *os.File
	reader *bufio.Reader
	buf_2  []byte
	buf_4  []byte
}

func (fileReader *FileReader) Init() {
	fileReader.buf_2 = make([]byte, 2)
	fileReader.buf_4 = make([]byte, 4)
}

func (fileReader *FileReader) OpenFile(name string, flat int, perm fs.FileMode) error {
	var err error

	fileReader.file, err = os.OpenFile(name, flat, perm)

	if err != nil {
		return err
	}

	fileReader.reader = bufio.NewReader(fileReader.file)
	return err
}

func (fileReader *FileReader) Available() bool {
	_, err := fileReader.reader.Peek(1)
	return err == nil
}

// 1 bytes 8 bits
func (fileReader *FileReader) Read() byte {
	data, _ := fileReader.reader.ReadByte()

	return data
}

func (fileReader *FileReader) ReadUint8() uint8 {
	data, _ := fileReader.reader.ReadByte()

	return data
}

// 2 bytes 16 bits
func (fileReader *FileReader) ReadUint16() uint16 {

	fileReader.reader.Read(fileReader.buf_2)

	return uint16(fileReader.buf_2[0])<<8 + uint16(fileReader.buf_2[1])
}

func (fileReader *FileReader) ReadUint16Array() []uint16 {
	fileReader.reader.Read(fileReader.buf_2)

	size := uint16(fileReader.buf_2[0])<<8 + uint16(fileReader.buf_2[1])

	buf := make([]byte, size*2)

	fileReader.reader.Read(buf)

	array := make([]uint16, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+2 {
		array[i] = uint16(buf[j])<<8 + uint16(buf[j+1])
	}

	return array
}

// 4 bytes 32 bits
func (fileReader *FileReader) ReadInt() int {

	fileReader.reader.Read(fileReader.buf_4)

	return int(fileReader.buf_4[0])<<24 + int(fileReader.buf_4[1])<<16 + int(fileReader.buf_4[2])<<8 + int(fileReader.buf_4[3])
}

func (fileReader *FileReader) ReadUint32() uint32 {

	fileReader.reader.Read(fileReader.buf_4)

	return uint32(fileReader.buf_4[0])<<24 + uint32(fileReader.buf_4[1])<<16 + uint32(fileReader.buf_4[2])<<8 + uint32(fileReader.buf_4[3])
}

// string
func (fileReader *FileReader) ReadString() string {
	data, _ := fileReader.reader.ReadBytes(0)

	return string(data[:len(data)-1])
}

// array
func (fileReader *FileReader) ReadIntArray() []int {
	fileReader.reader.Read(fileReader.buf_2)

	size := uint16(fileReader.buf_2[0])<<8 + uint16(fileReader.buf_2[1])

	buf := make([]byte, size*4)

	fileReader.reader.Read(buf)

	array := make([]int, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = int(buf[j])<<24 + int(buf[j+1])<<16 + int(buf[j+2])<<8 + int(buf[j+3])
	}

	return array
}

func (fileReader *FileReader) ReadUint32Array() []uint32 {
	fileReader.reader.Read(fileReader.buf_2)

	size := uint16(fileReader.buf_2[0])<<8 + uint16(fileReader.buf_2[1])

	buf := make([]byte, size*4)

	fileReader.reader.Read(buf)

	array := make([]uint32, size)

	for i, j := uint16(0), 0; i != size; i, j = i+1, j+4 {
		array[i] = uint32(buf[j])<<24 + uint32(buf[j+1])<<16 + uint32(buf[j+2])<<8 + uint32(buf[j+3])
	}

	return array
}

// /*
// 讀取一個 Media 物件，可能是 Anime 或 Manga 型別

// 僅從緩衝區讀取，需自行開啟、關閉檔案
// */
// func (fileReader *FileReader) ReadMedia() Media.IMedia {

// 	mediaType := Media.MediaType(fileReader.ReadUint8())

// 	if mediaType == Media.ANIME {
// 		anime := &Media.Anime{Type: mediaType}

// 		anime.Episodes = fileReader.ReadUint16()

// 		if anime.Episodes != 0 {
// 			anime.Videos = make([]string, anime.Episodes&32767)

// 			for i := 0; i != len(anime.Videos); i++ {
// 				anime.Videos[i] = fileReader.ReadString()
// 			}

// 			anime.ExEpisodes = fileReader.ReadUint32Array()
// 		}

// 		anime.Id_if101 = fileReader.ReadUint32()

// 		anime.Title = fileReader.ReadString()

// 		anime.Description = fileReader.ReadString()

// 		return anime
// 	} else if mediaType == Media.NOVEL {
// 		novel := &Media.Novel{Type: mediaType}

// 		novel.Volumes = fileReader.ReadUint16()

// 		novel.Title = fileReader.ReadString()

// 		novel.Description = fileReader.ReadString()

// 		return novel
// 	} else {
// 		manga := &Media.Manga{Type: mediaType}

// 		manga.Volumes = fileReader.ReadUint32Array()

// 		manga.Id_cartoonmad = fileReader.ReadUint32()

// 		manga.Title = fileReader.ReadString()

// 		manga.Description = fileReader.ReadString()

// 		return manga
// 	}
// }
// func (fileReader *FileReader) ReadMedia_MIN() Media.IMedia {

// 	mediaType := Media.MediaType(fileReader.ReadUint8())

// 	if mediaType == Media.ANIME {
// 		anime := &Media.Anime{Type: mediaType}

// 		anime.Episodes = fileReader.ReadUint16()

// 		if anime.Episodes != 0 {
// 			anime.Videos = make([]string, anime.Episodes&32767)

// 			buffer := make([]byte, (anime.Episodes&32767)*64)

// 			fileReader.reader.Read(buffer)

// 			for i := 0; i != len(anime.Videos); i++ {
// 				anime.Videos[i] = string(buffer[i*64 : (i+1)*64])
// 			}

// 			anime.ExEpisodes = fileReader.ReadUint32Array()
// 		}

// 		anime.Id_if101 = fileReader.ReadUint32()

// 		anime.Title = fileReader.ReadString()

// 		anime.Description = fileReader.ReadString()

// 		return anime
// 	} else if mediaType == Media.NOVEL {
// 		novel := &Media.Novel{Type: mediaType}

// 		novel.Volumes = fileReader.ReadUint16()

// 		novel.Title = fileReader.ReadString()

// 		novel.Description = fileReader.ReadString()

// 		return novel
// 	} else {
// 		manga := &Media.Manga{Type: mediaType}

// 		manga.Volumes = fileReader.ReadUint32Array()

// 		manga.Id_cartoonmad = fileReader.ReadUint32()

// 		manga.Title = fileReader.ReadString()

// 		manga.Description = fileReader.ReadString()

// 		return manga
// 	}
// }
func (fileReader *FileReader) Close() {
	fileReader.file.Close()
}

// SECTION_TYPE:FileWriter
type FileWriter struct {
	file     *os.File
	writer   *bufio.Writer
	buffer_m []byte
	buffer_p []byte
}

func (fileWriter *FileWriter) Init() {
	fileWriter.buffer_m = make([]byte, 1<<20)
	fileWriter.buffer_p = fileWriter.buffer_m[:]
}

func (fileWriter *FileWriter) OpenFile(name string, flat int, perm fs.FileMode) error {
	var err error
	fileWriter.file, err = os.OpenFile(name, flat, perm)

	if err != nil {
		return err
	}

	fileWriter.writer = bufio.NewWriter(fileWriter.file)

	fileWriter.buffer_p = fileWriter.buffer_m[:]
	return nil

}

// 1 byte 8 bits
func (fileWriter *FileWriter) Write(data byte) {
	fileWriter.buffer_p[0] = data

	fileWriter.buffer_p = fileWriter.buffer_p[1:]
}
func (fileWriter *FileWriter) WriteUint8(data uint8) {
	fileWriter.buffer_p[0] = data

	fileWriter.buffer_p = fileWriter.buffer_p[1:]
}

// 2 bytes 16 bits
func (fileWriter *FileWriter) WriteUint16(data uint16) {
	binary.BigEndian.PutUint16(fileWriter.buffer_p, data)

	fileWriter.buffer_p = fileWriter.buffer_p[2:]
}

func (fileWriter *FileWriter) WriteUint16Array(data []uint16) {
	binary.BigEndian.PutUint16(fileWriter.buffer_p, uint16(len(data)))

	fileWriter.buffer_p = fileWriter.buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint16(fileWriter.buffer_p, k)

		fileWriter.buffer_p = fileWriter.buffer_p[2:]
	}
}

// 4 bytes 32 bits
func (fileWriter *FileWriter) WriteInt(data int) {
	binary.BigEndian.PutUint32(fileWriter.buffer_p, uint32(data))

	fileWriter.buffer_p = fileWriter.buffer_p[4:]
}

func (fileWriter *FileWriter) WriteUint32(data uint32) {
	binary.BigEndian.PutUint32(fileWriter.buffer_p, data)

	fileWriter.buffer_p = fileWriter.buffer_p[4:]
}

// string
func (fileWriter *FileWriter) WriteString(data string) {
	tmp := append(fileWriter.buffer_p[0:0], data...)

	fileWriter.buffer_p[len(tmp)] = '\x00'

	fileWriter.buffer_p = fileWriter.buffer_p[len(tmp)+1:]
}

// array
func (fileWriter *FileWriter) WriteIntArray(data []int) {
	binary.BigEndian.PutUint16(fileWriter.buffer_p, uint16(len(data)))

	fileWriter.buffer_p = fileWriter.buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(fileWriter.buffer_p, uint32(k))

		fileWriter.buffer_p = fileWriter.buffer_p[4:]
	}
}

func (fileWriter *FileWriter) WriteUint32Array(data []uint32) {
	binary.BigEndian.PutUint16(fileWriter.buffer_p, uint16(len(data)))

	fileWriter.buffer_p = fileWriter.buffer_p[2:]

	for _, k := range data {
		binary.BigEndian.PutUint32(fileWriter.buffer_p, k)

		fileWriter.buffer_p = fileWriter.buffer_p[4:]
	}
}

// /*
// 寫入一個 Media 物件，可接受 Anime 及 Manga 型別

// 僅寫入至緩衝區，需自行開啟、關閉檔案
// */
// func (fileWriter *FileWriter) WriteMedia_MIN(media Media.IMedia) {
// 	fileWriter.Write(byte(media.GetType()))

// 	if media.GetType() == Media.ANIME {
// 		fileWriter.WriteUint16(media.(*Media.Anime).GetEpisodes())

// 		if media.(*Media.Anime).Episodes != 0 {
// 			temp := fileWriter.buffer_p[0:0]

// 			for _, video := range media.(*Media.Anime).Videos {
// 				temp = append(temp, video...)
// 			}

// 			fileWriter.buffer_p = fileWriter.buffer_p[len(temp):]

// 			fileWriter.WriteUint32Array(media.(*Media.Anime).GetExEpisodes())
// 		}

// 		fileWriter.WriteUint32(media.(*Media.Anime).GetId_if101())
// 	} else if media.GetType() == Media.NOVEL {
// 		fileWriter.WriteUint16(media.(*Media.Novel).GetVolumes())
// 	} else {
// 		fileWriter.WriteUint32Array(media.(*Media.Manga).GetVolumes())

// 		fileWriter.WriteUint32(media.(*Media.Manga).GetId_cartoonmad())
// 	}

// 	fileWriter.WriteString(media.GetTitle())

// 	fileWriter.WriteString(media.GetDescription())
// }

// func (fileWriter *FileWriter) WriteMedia(media Media.IMedia) {
// 	fileWriter.Write(byte(media.GetType()))

// 	if media.GetType() == Media.ANIME {
// 		fileWriter.WriteUint16(media.(*Media.Anime).GetEpisodes())

// 		if media.(*Media.Anime).Episodes != 0 {
// 			for _, video := range media.(*Media.Anime).Videos {
// 				fileWriter.WriteString(video)
// 			}

// 			fileWriter.WriteUint32Array(media.(*Media.Anime).GetExEpisodes())
// 		}

// 		fileWriter.WriteUint32(media.(*Media.Anime).GetId_if101())
// 	} else if media.GetType() == Media.NOVEL {
// 		fileWriter.WriteUint16(media.(*Media.Novel).GetVolumes())
// 	} else {
// 		fileWriter.WriteUint32Array(media.(*Media.Manga).GetVolumes())

// 		fileWriter.WriteUint32(media.(*Media.Manga).GetId_cartoonmad())
// 	}

// 	fileWriter.WriteString(media.GetTitle())

// 	fileWriter.WriteString(media.GetDescription())
// }

func (fileWriter *FileWriter) Flush() {
	fileWriter.writer.Write(fileWriter.buffer_m[:cap(fileWriter.buffer_m)-cap(fileWriter.buffer_p)])

	fileWriter.writer.Flush()

	fileWriter.buffer_p = fileWriter.buffer_m[:]
}

func (fileWriter *FileWriter) Close() {
	if cap(fileWriter.buffer_m) != cap(fileWriter.buffer_p) {
		fileWriter.Flush()
	}

	fileWriter.file.Close()
}
