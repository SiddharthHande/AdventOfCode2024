package helper

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func FileReader() (*[]string, error) {
	readfile, err := os.Open("iostream/input.txt")
	if err != nil {
		return nil, errors.New("error reading file")
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readfile.Close()
	return &fileLines, nil
}

func QuickParseInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	} else {
		log.Fatal(err)
		return 0
	}
}
