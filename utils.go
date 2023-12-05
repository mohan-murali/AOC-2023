package main

import (
	"bufio"
	"os"
)

func ReadFileToStringArray(path string) [] string{
	fileReader, err := os.Open(path)
	check(err)

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan(){
		lines = append(lines, fileScanner.Text())
	}

	return lines
}