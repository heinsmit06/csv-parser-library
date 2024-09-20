package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser CSVParser = &CSVStruct{}

	lineCount := 1
	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Printf("Line %v: %v\n", lineCount, line)
		fmt.Println("  ---------------  ")
		lineCount++
	}
}

func describe(i CSVParser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
