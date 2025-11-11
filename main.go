package main

import (
	"os"
	helper "project1/helpers"
)

func main() {
	if len(os.Args) != 2 {
		panic("No Args")
	}
	inputF := os.Args[1]

	data, err := os.ReadFile(inputF)
	if err != nil {
		panic(err)
	}
	outputF, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	_, err = outputF.WriteString(helper.ToUpper(string(data)))
	if err != nil {
		panic(err)
	}
}
