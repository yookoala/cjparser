package main

import (
	"bufio"
	"fmt"
	"io"
	//"log"
	"os"
	"strings"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file *os.File
		line string
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		if line, err = reader.ReadString('\r'); err != nil {
			break
		}
		if len(line) > 0 && line[0] != '#' {
			line = strings.TrimRight(line, '\r')
			lines = append(lines, line)
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
