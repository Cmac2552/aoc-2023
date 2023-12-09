package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		fields := stringArrayToInts(strings.Fields(text))
		slices.Reverse(fields)
		differences := make([][]int, 0)
		differences = append(differences, fields)

		for compare(differences[len(differences)-1]) {
			differences = append(differences, makeDifferences(differences[len(differences)-1]))

		}
		sum += calculateNextField(differences)

	}

	fmt.Println(sum)

}

func stringArrayToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, str := range strings {
		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		ints[i] = number

	}
	return ints
}

func makeDifferences(ints []int) []int {
	newDifferences := make([]int, len(ints)-1)
	for i := 0; i < len(ints)-1; i++ {
		newDifferences[i] = ints[i+1] - ints[i]
	}
	return newDifferences
}

func compare(ints []int) bool {
	for _, number := range ints {
		if number != 0 {
			return true
		}
	}
	return false
}

func calculateNextField(differences [][]int) int {
	field := 0
	for i := len(differences) - 1; i > -1; i-- {
		field += differences[i][len(differences[i])-1]
	}
	return field
}
