package main

import (
	"fmt"
	"os"
	helper "project1/helpers"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Args")
		return
	}

	inputFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%s File Not Found", os.Args[1])
		return
	}

	if !helper.CheckExtension(os.Args[2]) {
		panic("Invalid Output File")
	}

	output := helper.ParseString(string(inputFile))
	file, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	file.WriteString(output)
}
