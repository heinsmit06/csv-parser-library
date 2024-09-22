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
				fmt.Printf("Line %v: %v\n", lineCount, line)
				for i := 0; i < csvparser.GetNumberOfFields(); i++ {
					field, err := csvparser.GetField(i)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Printf("  Field %v: %v\n", i, field)
				}
				break
			}
			fmt.Println("Error reading line:", err)
			fmt.Println("  ---------------  ")
			lineCount++
			continue
			// return
		}

		fmt.Printf("Line %v: %v\n", lineCount, line)

		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			field, err := csvparser.GetField(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("  Field %v: %v\n", i, field)
		}

		fmt.Println("  ---------------  ")
		lineCount++
	}
}

func describe(i CSVParser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
