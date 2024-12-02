package main

import (
	code1 "adventofcode/code_1"
	"adventofcode/helper"
	"log"
)

func main() {
	data, err := helper.FileReader()
	if err != nil {
		log.Fatal(err)
	}
	code1.ProblemTwo(data)
}
