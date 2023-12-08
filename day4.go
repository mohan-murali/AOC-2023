package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day4_task1(){
	lines := readFileToStringArray("./data/day4.txt")
	sum:=0

	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	for _, line:= range lines {
		game := strings.Split(line, ":")
		gameNumbers := strings.Split(game[1], "|")
		lucky, numbers := strings.TrimSpace(gameNumbers[0]), strings.TrimSpace(gameNumbers[1])
		luck := strings.Split(lucky, " ")
		curr := strings.Split(numbers, " ")
		luckyNumbers := make(map[int] bool)
		for _, num:= range luck{
			number, err := strconv.Atoi(num)
			if(err == nil){
				luckyNumbers[number]=false
			}
		}
		fmt.Println(luck, curr)
		numbersExist:=false
		power :=1
		for _, num:= range curr {
			number, err:= strconv.Atoi(num)
			if(err == nil){
				_, exists := luckyNumbers[number]
				if exists {
					luckyNumbers[number] = true
					if numbersExist {
						power *=2
					}
					numbersExist = true
					fmt.Println(number)
				}
			}
		}

		if numbersExist {
			sum += power
		}
	}

	fmt.Println(sum)
}