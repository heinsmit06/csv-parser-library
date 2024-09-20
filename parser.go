package main

import (
	"errors"
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type CSVStruct struct {
	fieldInBytes []byte
	line         []string
}

func (p *CSVStruct) ReadLine(r io.Reader) (string, error) {
	b := make([]byte, 1) // to read one byte at a time
	var err error
	p.line = []string{}

	for {
		_, err = r.Read(b)
		if err == io.EOF {
			return "", err
		}

		if lineIsTerminated(b[0]) {
			p.line = append(p.line, string(p.fieldInBytes))
			p.fieldInBytes = []byte{}
			return sliceToStr(p.line), err
		}

		if b[0] == byte(44) {
			p.line = append(p.line, string(p.fieldInBytes))
			p.fieldInBytes = []byte{}
			continue
		}

		p.fieldInBytes = append(p.fieldInBytes, b[0])
	}
}

func (p CSVStruct) GetField(n int) (string, error) {
	if n < 0 || n > len(p.line) {
		return "", ErrFieldCount
	}

	return p.line[n], nil
}

func (p CSVStruct) GetNumberOfFields() int {
	return 1
}

func sliceToStr(line []string) string {
	lineString := ""
	for i, v := range line {
		if i == len(line)-1 {
			lineString += v
		} else {
			lineString += v + ","
		}
	}
	return lineString
}

func lineIsTerminated(b byte) bool {
	if b == byte(10) || b == byte(13) {
		return true
	}
	return false
}
