package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main3() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		score := 0
		line := scanner.Text()
		idAndGame := strings.Split(line, ":")
		winningAndRegularNumbers := strings.Split(idAndGame[1], " | ")
		winningNumbers := strings.Fields(winningAndRegularNumbers[0])
		regularNumbers := strings.Fields(winningAndRegularNumbers[1])
		for i := 0; i < len(regularNumbers); i++ {
			if slices.Contains(winningNumbers, regularNumbers[i]) {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}

		}
		sum += score

	}

	fmt.Println(sum)
}
