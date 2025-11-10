package main

import (
	"fmt"
	"os"
	"project1/helpers"
)

func main() {
	if len(os.Args) != 2 {
		panic("No Args")
	}
	input := os.Args[1]
	//output := os.Args[2]

	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	fmt.Print(helper.ToUpper(string(file)))
}
