package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
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

func parseCangjie3File(filename string, handlerName string,
	sepDist int, cat string, src string, db *ConvertDB) {

	var serial uint32

	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		line = strings.TrimRight(line, "\t ")
		if len(line) > sepDist && line[0] != ' ' {
			serial += 1
			partCode := strings.Trim(line[0:sepDist], "\t ")
			partChar := strings.Trim(line[sepDist:], "\t ")
			unicode, _ := utf8.DecodeRuneInString(partChar)
			item := cangjieValue{
				Unicode:   strings.ToUpper(fmt.Sprintf("U+%x", unicode)),
				Character: partChar,
				Version:   handlerName,
				Category:  cat,
				Code:      strings.ToLower(partCode),
				Source:    src,
				Serial:    serial,
			}
			err = db.Insert(handlerName, tx, item)
			if err != nil {
				panic(err)
			}

		}
	}
}

func parseCangjie5File(filename string, handlerName string, cat string, src string, db *ConvertDB) {

	var (
		sect   string
		serial uint32
	)

	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		line = strings.Trim(line, "\t ")
		if line[0] == '[' && line[len(line)-1] == ']' {
			sect = line[1 : len(line)-1]
		} else if sect == "Text" && len(line) != 0 {
			serial += 1
			splited := strings.SplitN(line, "", 2)
			unicode, _ := utf8.DecodeRuneInString(splited[0])
			item := cangjieValue{
				Unicode:   strings.ToUpper(fmt.Sprintf("U+%x", unicode)),
				Character: splited[0],
				Version:   handlerName,
				Category:  cat,
				Code:      strings.ToLower(splited[1]),
				Source:    src,
				Serial:    serial,
			}
			err = db.Insert(handlerName, tx, item)
			if err != nil {
				panic(err)
			}

		}
	}
}

type cangjieValue struct {
	Unicode   string
	Character string
	Version   string
	Category  string
	Code      string
	Source    string
	Serial    uint32
}
