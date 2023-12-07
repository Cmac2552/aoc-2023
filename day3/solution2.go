package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	counter := 0
	idCounter := 1
	lines := make([][]rune, 0)
	locationMap := map[string]int{}
	idMap := map[int]string{}
	for scanner.Scan() {
		text := scanner.Text()
		line := []rune(text)
		lines = append(lines, line)
	}

	for r := 0; r < len(lines); r++ {
		currLine := lines[r]
		currNum := ""
		lastCharNumber := false
		for c := 0; c < len(currLine); c++ {
			if isNumber2(currLine[c]) && !lastCharNumber {
				currNum += string(currLine[c])
				lastCharNumber = true
			} else if isNumber2(currLine[c]) {
				currNum += string(currLine[c])
				if c+1 == len(currLine) {
					lastCharNumber = false

					for i := 0; i < len(currNum); i++ {
						locationMap[strconv.Itoa(r)+","+strconv.Itoa(c-i)] = idCounter
					}
					idMap[idCounter] = currNum
					idCounter++
					currNum = ""
				}

			} else {
				if lastCharNumber {
					lastCharNumber = false

					for i := 1; i < len(currNum)+1; i++ {
						locationMap[strconv.Itoa(r)+","+strconv.Itoa(c-i)] = idCounter
					}
					idMap[idCounter] = currNum
					idCounter++

				}
				currNum = ""
			}

		}
	}
	for r := 0; r < len(lines); r++ {
		currLine := lines[r]
		for c := 0; c < len(currLine); c++ {
			if isStar(currLine[c]) {
				ints := make([]int, 0)
				for r2 := -1; r2 <= 1; r2++ {
					for c2 := -1; c2 <= 1; c2++ {
						if locationMap[strconv.Itoa(r+r2)+","+strconv.Itoa(c+c2)] != 0 && !slices.Contains(ints, locationMap[strconv.Itoa(r+r2)+","+strconv.Itoa(c+c2)]) {
							ints = append(ints, locationMap[strconv.Itoa(r+r2)+","+strconv.Itoa(c+c2)])
						}
					}
				}
				if len(ints) == 2 {
					number1, err := strconv.Atoi(idMap[ints[0]])
					if err != nil {
						fmt.Println(err)
					}
					number2, err := strconv.Atoi(idMap[ints[1]])
					if err != nil {
						fmt.Println(err)
					}
					counter++
					fmt.Println(number1, number2)
					sum += number1 * number2
				} else {
					fmt.Println(strconv.Itoa(r+1)+","+strconv.Itoa(c+1), "here")
				}
			}

		}
	}
	fmt.Println(locationMap)

	fmt.Println(sum)
}

func isNumber2(char rune) bool {
	return unicode.IsDigit(char)
}

func isStar(char rune) bool {
	return string(char) == "*"
}
