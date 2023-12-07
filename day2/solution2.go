package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		letterMap := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		text := scanner.Text()
		idAndGames := strings.Split(text, ":")

		individualGames := strings.Split(idAndGames[1], ";")

		for i := 0; i < len(individualGames); i++ {
			handfuls := strings.Split(individualGames[i], ",")
			for j := 0; j < len(handfuls); j++ {
				handful := strings.Split(handfuls[j], " ")

				cubes, err := strconv.Atoi(handful[1])
				if err != nil {
					fmt.Println(err)
				}
				///////////////
				///////////////
				///////////////
				///////////////

				if cubes > letterMap[handful[2]] || letterMap[handful[2]] == 0 {
					fmt.Println("test")
					letterMap[handful[2]] = cubes
				}

			}
		}

		// gameId, err := strconv.Atoi(strings.Split(idAndGames[0], " ")[1])

		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Println(letterMap)
		sum += letterMap["red"] * letterMap["blue"] * letterMap["green"]

	}
	fmt.Println(sum)

}
