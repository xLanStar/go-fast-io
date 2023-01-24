package binaryi_test

import (
	"fmt"
	"os"
	"testing"

	fastio "github.com/xLanStar/go-fast-io/v2"
)

func TestF(t *testing.T) {
	// var fileWriter fastio.FileWriter

	// fileWriter.Init()

	// fileWriter.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)

	// fileWriter.WriteInt(144)
	// a := []int{1, 2, 3}
	// fileWriter.WriteIntArray(a)

	// fileWriter.WriteString("123456")
	// fileWriter.Flush()
	// fileWriter.Close()

	var fileReader fastio.FileReader

	fileReader.Init()

	fileReader.OpenFile("test.txt", os.O_RDONLY, 0666)

	fmt.Println(fileReader.ReadInt())
	fmt.Println(fileReader.ReadIntArray())

	fmt.Println(fileReader.ReadString())
	fileReader.Close()
}
