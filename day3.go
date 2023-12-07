package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func day3_task1(){
	lines:= readFileToStringArray("./data/day3.txt")
	for _, line := range lines {
		number := ""
		digitStarted := false
		for _, char := range line {

			if unicode.IsDigit(char){
				digitStarted = true
				number += string(char)
			}

			if !unicode.IsDigit(char) && digitStarted {
				num, err := strconv.Atoi(number)
				check(err)
				fmt.Println(num)
				number = ""
				digitStarted = false
			}
		}

	}
}