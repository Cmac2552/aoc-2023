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
	letterMap := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		idAndGames := strings.Split(text, ":")
		win := true
		individualGames := strings.Split(idAndGames[1], ";")
		for i := 0; i < len(individualGames); i++ {
			handfuls := strings.Split(individualGames[i], ",")
			for j := 0; j < len(handfuls); j++ {
				handful := strings.Split(handfuls[j], " ")

				cubes, err := strconv.Atoi(handful[1])
				if err != nil {
					fmt.Println(err)
				}
				if cubes > letterMap[handful[2]] {

					win = false
				}
			}
		}

		if win {
			gameId, err := strconv.Atoi(strings.Split(idAndGames[0], " ")[1])

			if err != nil {
				fmt.Println(err)
			}
			sum += gameId
		}

	}
	fmt.Println(sum)

}
