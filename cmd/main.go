package main

import (
	"adventofcode/helper"
	problemDayOne "adventofcode/problem_dayone"
	"log"
)

func main() {
	data, err := helper.FileReader()
	if err != nil {
		log.Fatal(err)
	}
	problemDayOne.ProblemTwo(data)
}
