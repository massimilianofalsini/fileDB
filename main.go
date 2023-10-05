package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
)

type row struct {
	id      string
	name    string
	size    int64
	content []byte
}

func createRow() row {
	file, err := os.Open("examplefile.mov")
	if err != nil {
		fmt.Println(err)
		return row{}
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return row{}
	}

	tmp := row{uuid.New().String(), file.Name(), stat.Size(), make([]byte, stat.Size())}
	_, err = bufio.NewReader(file).Read(tmp.content)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return row{}
	}

	return tmp
}

func appendRow(m row) {
	if _, err := os.Stat("./test/" + m.name); os.IsNotExist(err) {
		os.MkdirAll("./test/", 0700)
	}
	f, err := os.Create("./test/" + m.name)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	_, err2 := f.Write(m.content)

	if err2 != nil {
		log.Fatal(err2)
		return
	}
}

func main() {
	tmp := createRow()
	appendRow(tmp)
	fmt.Println("done")
}
