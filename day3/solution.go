package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main2() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	lines := make([][]rune, 0)
	for scanner.Scan() {
		text := scanner.Text()
		line := []rune(text)
		lines = append(lines, line)
	}

	for r := 0; r < len(lines); r++ {
		currLine := lines[r]
		currNum := ""
		check := false
		for c := 0; c < len(currLine); c++ {
			if isNumber(currLine[c]) && !check {
				currNum += string(currLine[c])
				boundDigits := inBoundsDigits(r, len(lines), c, len(currLine))
				for r2 := r + boundDigits[0]; r2 <= r+boundDigits[1]; r2++ {
					for c2 := c + boundDigits[2]; c2 <= c+boundDigits[3]; c2++ {

						check = isSymbol(lines[r2][c2])
						fmt.Println(check, r, c)
						if check {
							goto out
						}

					}
				}
			out:
			} else if isNumber(currLine[c]) && check {
				currNum += string(currLine[c])
				if c+1 == len(currLine) {
					number, err := strconv.Atoi(currNum)
					if err != nil {
						fmt.Println(nil)
					}

					fmt.Println(number, r)

					sum += number
					currNum = ""
					check = false
				}

			} else if check {
				number, err := strconv.Atoi(currNum)
				if err != nil {
					fmt.Println(nil)
				}

				fmt.Println(number, r)

				sum += number
				currNum = ""
				check = false
			} else {
				currNum = ""
			}

		}
	}

	fmt.Println(sum)
}

func isNumber(char rune) bool {
	return unicode.IsDigit(char)
}

func isSymbol(char rune) bool {
	return !unicode.IsDigit(char) && string(char) != "."
}

func inBoundsDigits(rowYouAreIn int, totalRows int, columnYouAreIn int, totalCoulmun int) []int {
	returnArray := make([]int, 0)
	if rowYouAreIn == 0 {
		returnArray = append(returnArray, 0)
	} else {
		returnArray = append(returnArray, -1)
	}

	if rowYouAreIn+1 == totalRows {
		returnArray = append(returnArray, 0)
	} else {
		returnArray = append(returnArray, 1)
	}

	if columnYouAreIn == 0 {
		returnArray = append(returnArray, 0)
	} else {
		returnArray = append(returnArray, -1)
	}

	if columnYouAreIn+1 == totalCoulmun {
		returnArray = append(returnArray, 0)
	} else {
		returnArray = append(returnArray, 1)
	}

	return returnArray
}
