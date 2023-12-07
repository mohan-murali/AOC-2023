package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func day3_task1(){
	lines:= readFileToStringArray("./data/day3.txt")
	sum:= 0
	for i, line := range lines {
		number := ""
		digitStarted := false
		for j, char := range line {
			left, right, isValidNumber := 0, 0, false

			if unicode.IsDigit(char){
				if !digitStarted {
					left = j
				}
				digitStarted = true
				number += string(char)
			}

			if !unicode.IsDigit(char) && digitStarted {
				num, err := strconv.Atoi(number)
				check(err)
				number = ""
				digitStarted = false
				right = j

				if left > 0 {
					sym := string(line[left-1])
					if !unicode.IsDigit(rune(line[left-1])) && sym != "."{
						isValidNumber = true
					}
				} else if right < len(line)-1 {
					sym := string(line[right + 1])
					fmt.Println(sym)
					if !unicode.IsDigit(rune(line[right + 1])) && sym != "."{
						isValidNumber = true
					}
				}

				if !isValidNumber && i > 0 {
					prevLine := lines[i-1]
					leftMost, rightMost := 0, 0
					if left > 0 {
						leftMost = left-1
					} else {
						leftMost = 0
					}

					if right < len(line) -1 {
						rightMost = right +1
					} else {
						rightMost = right
					}

					for l:= leftMost; l < rightMost; l++{
						sym := string(prevLine[l])
							if !unicode.IsDigit(rune(prevLine[l])) && sym != "."{
							isValidNumber = true
						}
					}
				}

				if !isValidNumber && i < len(lines) -1 {
					nextLine := lines[i+1]
					leftMost, rightMost := 0, 0
					if left > 0 {
						leftMost = left-1
					} else {
						leftMost = 0
					}

					if right < len(line) -1 {
						rightMost = right +1
					} else {
						rightMost = right
					}

					for l:= leftMost; l < rightMost; l++{
						sym := string(nextLine[l])
							if !unicode.IsDigit(rune(nextLine[l])) && sym != "."{
							isValidNumber = true
						}
					}
				}

				if isValidNumber {
					fmt.Println(num)
					sum += num
				}
			}
		}
	}

	fmt.Println(sum)
}