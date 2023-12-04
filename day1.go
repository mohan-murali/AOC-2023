package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main()  {
	// task1()
	task2()
}

func task2(){
	fileReader, err := os.Open("./data/day1-2.txt")
	check(err)

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan(){
		lines = append(lines, fileScanner.Text())
	}

	fileReader.Close()

	source := map [string] int {
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	sum:= 0

	for _, line := range lines {
		firstIndex, lastIndex := math.MaxInt, math.MinInt
		firstVal, secondVal := 0,0

			for k, v := range source {
				if index := strings.Index(line, k); index >=0 {
					if(index < firstIndex){
						firstIndex = index
						firstVal = v
					}

					if(index > lastIndex){
						lastIndex = index
						secondVal = v
					}
				}
			}
		fmt.Println(firstVal, secondVal)
		sum = sum + (firstVal * 10) + secondVal
	}

	fmt.Println(sum)
}

func task1(){
	dat, err := os.Open("./data/day1-1.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	sum:= 0

	for fileScanner.Scan(){
		var numberArray [] int
		line := fileScanner.Text()
		for _, c := range line {
			if unicode.IsDigit(c){
				num , e := strconv.Atoi(string(c))
				check(e)
				numberArray = append(numberArray, num)
			}
		}
		length := len(numberArray)
		first := numberArray[0]
		last := numberArray[length -1]
		fmt.Println(first, last, sum)
		sum = sum + (first * 10) + last
	}

	dat.Close()

	fmt.Println(sum)
}

func check(e error){
	if e!= nil {
		panic(e)
	}
}