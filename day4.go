package main

import (
	"fmt"
	"strings"
)

func day4_task1(){
	lines := readFileToStringArray("./data/day4.txt")

	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	for _, line:= range lines {
		game := strings.Split(line, ":")
		gameNumbers := strings.Split(game[1], "|")
		lucky, numbers := gameNumbers[0], gameNumbers[1]
		luck := strings.Split(lucky, " ")
		curr := strings.Split(numbers, " ")
		fmt.Println(luck, curr)
	}
}