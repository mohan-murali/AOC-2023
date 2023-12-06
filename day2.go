package main

import (
	"fmt"
	"strconv"
	"strings"
)

//only 12 red cubes, 13 green cubes, and 14 blue cubes

var red = 12
var green = 13
var blue = 14

func day2_task1(){
	lines := readFileToStringArray("./data/day2-1.txt")
	sum := 0
	//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	for _, line := range lines {
		gameInfo := strings.Split(line, ":")
		gameNumber, playedGames := gameInfo[0], gameInfo[1]
		gameId, err :=  strconv.Atoi(strings.Split(gameNumber, " ")[1])
		check(err)
		games := strings.Split(playedGames, ";")
		isValidGame := true
		for _, game := range games{
			if !isCubesValid(game){
				isValidGame = false
			}
		}

		if isValidGame {
			sum = sum + gameId
		}
	}

	fmt.Println(sum)

}

func isCubesValid(game string) bool{
	cubes := strings.Split(game, ",")

	for _, cube := range cubes {
		cubeInfo := strings.Split(cube, " ")
		n, color := cubeInfo[1], cubeInfo[2]
		number, err := strconv.Atoi(n)
		check(err)
		switch color {
		case "green":
			if number > green {
				return false
			}
		case "blue":
			if number > blue {
				return false
			}
		case "red":
			if number > red {
				return false
			}
		}
	}
	return true
}

func day2_task2(){
	lines:= readFileToStringArray("./data/day2-1.txt")
	result:= 0

	for _, line := range lines {
		gameInfo := strings.Split(line, ":")
		playedGames := gameInfo[1]
		games := strings.Split(playedGames, ";")
		records := make(map[string] int)
		for _, game := range games {
			cubes := strings.Split(game, ",")

			for _, cube := range cubes {
				cubeInfo := strings.Split(cube, " ")
				n, color := cubeInfo[1], cubeInfo[2]
				number, err := strconv.Atoi(n)
				check(err)

				value, exists := records[color]

				if exists {
					if number > value {
						records[color] = number
					}
				} else {
					records[color] = number
				}
				fmt.Println(records)
			}
		}

		power := 1

		for key, value := range records {
			fmt.Println(key, value)
			power = power * value
		}

		fmt.Println(power)

		result = result + power
	}

	fmt.Println(result)
}