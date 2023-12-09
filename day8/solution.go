package main

import (
	"bufio"
	"fmt"
	"os"
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
	sum := 0
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		tree[fields[0]] = node{left: fields[2][1:4], right: fields[3][0:3]}

	}

	orderCounter := 0
	order := []rune(directions)
	value := "AAA"
	for value != "ZZZ" {
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

	fmt.Println(sum)

}
