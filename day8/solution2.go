package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	tree := map[string]node{}
	scanner := bufio.NewScanner(file)

	as := make([]string, 0)
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		tree[fields[0]] = node{left: fields[2][1:4], right: fields[3][0:3]}
		if string([]rune(fields[0])[2]) == "A" {
			as = append(as, fields[0])
		}

	}
	sums := make([]int, 0)
	for i := 0; i < len(as); i++ {
		orderCounter := 0
		order := []rune(directions)
		value := as[i]
		sum := 0
		for string([]rune(value)[2]) != "Z" {
			if string(order[orderCounter]) == "L" {
				value = tree[value].left
			} else {
				value = tree[value].right
			}
			orderCounter++
			if orderCounter == len(directions) {
				orderCounter = 0
			}
			sum++
		}
		sums = append(sums, sum)
	}

	divisorsArray := make([][]int, len(sums))
	for i := 0; i < len(sums); i++ {
		divisorsArray[i] = divisors(sums[i])
		sort.Ints(divisorsArray[i])

	}

	lcm := (sums[0] * sums[1]) / gcd(divisorsArray[0], divisorsArray[1])

	for i := 2; i < len(sums); i++ {
		newLCMdivisors := divisors(lcm)
		sort.Ints(newLCMdivisors)
		lcm = (sums[i] * lcm) / gcd(divisorsArray[i], newLCMdivisors)
	}

	fmt.Println(lcm)

}

func divisors(number int) []int {
	counter := 1
	numbers := make([]int, 0)

	for number/counter > counter {
		if number%counter == 0 {
			numbers = append(numbers, counter)
			numbers = append(numbers, number/counter)
		}
		counter++
	}
	return numbers

}

func gcd(numbers1 []int, numbers2 []int) int {
	for i := len(numbers1) - 2; i > -1; i++ {
		if slices.Contains(numbers2, numbers1[i]) {
			return numbers1[i]
		}
	}
	return 1
}
