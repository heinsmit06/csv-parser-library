package parsers

import (
	"errors"
	"fmt"
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
	byteSlice []byte
}

func (p *CSVStruct) ReadLine(r io.Reader) (string, error) {
	p.byteSlice = make([]byte, 1)
	r.Read(p.byteSlice)
	fmt.Print(p.byteSlice)
	return "", ErrQuote
}

func (p *CSVStruct) GetField(n int) (string, error) {
	return "", ErrQuote
}

func (p *CSVStruct) GetNumberOfFields() int {
	return 1
}
