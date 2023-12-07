package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cardMap := map[int]int{1: 0}
	counter := 1

	for scanner.Scan() {

		line := scanner.Text()
		idAndGame := strings.Split(line, ":")
		winningAndRegularNumbers := strings.Split(idAndGame[1], " | ")
		winningNumbers := strings.Fields(winningAndRegularNumbers[0])
		regularNumbers := strings.Fields(winningAndRegularNumbers[1])
		for j := 0; j < cardMap[counter]+1; j++ {
			score := 0
			for i := 0; i < len(regularNumbers); i++ {
				if slices.Contains(winningNumbers, regularNumbers[i]) {
					score++
				}

			}
			for i := 1; i < score+1; i++ {
				cardMap[counter+i] = cardMap[counter+i] + 1
			}
		}
		cardMap[counter] = cardMap[counter] + 1
		counter++

	}
	totalScore := 0
	for _, v := range cardMap {
		totalScore += v
	}

	fmt.Println(totalScore)
}
