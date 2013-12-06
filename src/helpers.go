package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	//"log"
	"os"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		if len(part) > 0 && part[0] != '#' {
			buffer.Write(part)
			if !prefix {
				lines = append(lines, buffer.String())
				buffer.Reset()
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func parseCangjie5File(filename string, handlerName string, db *ConvertDB) {
	//tx, err := db.DB.Begin()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer tx.Commit()
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		fmt.Println(line)
		/*
			item, err := parseLine(line)
			if err != nil {
				panic(err)
			}
			err = db.Insert(handlerName, tx, item)
			if err != nil {
				panic(err)
			}
		*/
	}
}

type cangjieValue struct {
	Unicode   string
	Character string
	Code      string
	Serial    uint32
}
