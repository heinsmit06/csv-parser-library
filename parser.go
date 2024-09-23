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
	previousByte byte
}

func (p *CSVStruct) ReadLine(r io.Reader) (string, error) {
	b := make([]byte, 1) // to read one byte at a time
	var err error
	p.line = []string{}
	firstByteIsQuote := false
	illegalLine := false
	EOFflag := false

	for {
		_, err = r.Read(b)
		if err == io.EOF {
			EOFflag = true
		}

		if firstByteIsQuote {
			// the field starting with a quote is finalized only if
			// the previous byte is '"' and the current byte is ',' and there is even number of '"'
			if b[0] == ',' && p.previousByte == '"' && countQuotesInField(p.fieldInBytes)%2 == 0 {
				p.line = append(p.line, string(p.fieldInBytes[1:len(p.fieldInBytes)-1]))
				p.fieldInBytes = []byte{}
				firstByteIsQuote = false
				continue
			}
			if b[0] == ',' && p.previousByte == '"' && countQuotesInField(p.fieldInBytes)%2 == 1 {
				p.line = append(p.line, string(p.fieldInBytes[1:len(p.fieldInBytes)-1]))
				p.fieldInBytes = []byte{}
				firstByteIsQuote = false
				illegalLine = true
				continue
			}
			if lineIsTerminated(b[0]) || err == io.EOF {
				if p.previousByte == '"' {
					if countQuotesInField(p.fieldInBytes)%2 == 1 {
						p.line = append(p.line, string(p.fieldInBytes[1:len(p.fieldInBytes)-1]))
						p.fieldInBytes = []byte{}
						firstByteIsQuote = false
						illegalLine = true
					}
				} else if EOFflag {
				} else {
					continue
				}
			}
		} else {
			if b[0] == ',' {
				if countQuotesInField(p.fieldInBytes) > 0 {
					p.fieldInBytes = []byte{}
					illegalLine = true
					continue
				}
				p.line = append(p.line, string(p.fieldInBytes))
				p.fieldInBytes = []byte{}
				continue
			} else if lineIsTerminated(b[0]) || err == io.EOF {
				if countQuotesInField(p.fieldInBytes) > 0 {
					illegalLine = true
				}
			}
		}

		if lineIsTerminated(b[0]) || EOFflag {
			if illegalLine {
				p.fieldInBytes = []byte{}
				if EOFflag {
					return "", ErrQuote
				}
				return "", ErrQuote
			}

			if firstByteIsQuote && p.previousByte == '"' && countQuotesInField(p.fieldInBytes)%2 == 0 {
				p.line = append(p.line, string(p.fieldInBytes[1:len(p.fieldInBytes)-1]))
				p.fieldInBytes = []byte{}
			} else {
				p.line = append(p.line, string(p.fieldInBytes))
				p.fieldInBytes = []byte{}
			}

			// If EOF is true, return EOF after processing the line
			if EOFflag {
				return sliceToStr(p.line), io.EOF
			}

			return sliceToStr(p.line), err
		}

		p.fieldInBytes = append(p.fieldInBytes, b[0])
		p.previousByte = b[0]

		if p.fieldInBytes[0] == '"' {
			firstByteIsQuote = true
		}
	}
}

func (p CSVStruct) GetField(n int) (string, error) {
	if n < 0 || n > len(p.line) {
		return "", ErrFieldCount
	}
	stringWithoutDoubleQuotes := ""

	for i := range p.line[n] {
		if i < len(p.line[n])-1 {
			if p.line[n][i] == '"' && p.line[n][i+1] == '"' {
				stringWithoutDoubleQuotes = p.line[n][:i] + p.line[n][i+1:]
				p.line[n] = stringWithoutDoubleQuotes
			}
		}
	}

	return p.line[n], nil
}

func (p CSVStruct) GetNumberOfFields() int {
	return len(p.line)
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

func countQuotesInField(field []byte) int {
	count := 0
	for _, v := range field {
		if v == '"' {
			count++
		}
	}
	return count
}

func countCommas(field []byte) int {
	count := 0
	for _, v := range field {
		if v == ',' {
			count++
		}
	}
	return count
}

func indexOfLastQuote(field []byte) int {
	idx := 0
	for i, v := range field {
		if v == '"' {
			idx = i
		}
	}
	return idx
}
