package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	firstLetterMap := map[string][]string{
		"o": {"one"},
		"t": {"two", "three"},
		"f": {"four", "five"},
		"s": {"six", "seven"},
		"e": {"eight"},
		"n": {"nine"},
	}
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := []rune(text)
		firstNumber := -1
		secondNumber := -1
		for i := 0; i < len(line); i++ {
			letter, err := strconv.Atoi(string(line[i]))
			if err != nil {
				if len(firstLetterMap[string(line[i])]) == 0 {
					continue
				}
				for j := 0; j < len(firstLetterMap[string(line[i])]); j++ {
					upperbound := i + len(firstLetterMap[string(line[i])][j])

					if upperbound >= len(text) {
						upperbound = len(text)
					}
					subString := text[i:upperbound]
					if numberMap[subString] != 0 {
						letter = numberMap[subString]
						break
					} else {
						letter = -1
					}
				}
			}
			if firstNumber == -1 && letter != -1 {
				firstNumber = letter
				secondNumber = letter

			} else if letter != -1 {
				secondNumber = letter
			}

		}
		numberForLine, err := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(secondNumber))
		if err != nil {
			fmt.Println(err)
		}
		sum += numberForLine

	}
	fmt.Println(sum)

}
