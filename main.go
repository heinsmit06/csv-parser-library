package main

import (
	"a-library-for-others/parsers"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	/*var csvparser parsers.CSVParser = &parsers.CSVStruct{}
	describe(csvparser)

	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Println(line)
	}*/
	p := make([]byte, 1)
	file.Read(p)
	fmt.Print(p)
	file.Read(p)
	fmt.Print(p)
	file.Read(p)
	fmt.Print(p)
	file.Read(p)
	fmt.Print(p)
	file.Read(p)
	fmt.Println(p)

	// r := strings.NewReader("Hello, Reader!")

	// b := make([]byte, 3)
	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}

func describe(i parsers.CSVParser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
