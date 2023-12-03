package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main()  {
	task1()
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