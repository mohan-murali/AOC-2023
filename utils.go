package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileToStringArray(path string) [] string{
	fmt.Print(path)
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