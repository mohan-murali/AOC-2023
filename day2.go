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
	lines := ReadFileToStringArray("./data/day2-1.txt")

	//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	for _, line := range lines {
		gameInfo := strings.Split(line, ":");
		gameNumber, playedGames := gameInfo[0], gameInfo[1]
		gameId, err :=  strconv.Atoi(strings.Split(gameNumber, " ")[1]);
		fmt.Println(gameId);
		check(err)
		games := strings.Split(playedGames, ";")
		for _, game := range games{
			fmt.Println(game);
		}
	}

}