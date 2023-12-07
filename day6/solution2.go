package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 1
	scanner.Scan()
	times := strings.Fields(scanner.Text())
	scanner.Scan()
	distances := strings.Fields(scanner.Text())
	noKerningTime := ""
	noKerningDistance := ""
	for i := 1; i < len(times); i++ {
		noKerningTime += times[i]
		noKerningDistance += distances[i]
	}
	counter := 0
	currTime, err := strconv.Atoi(noKerningTime)
	if err != nil {
		fmt.Println(err)
	}

	currDistance, err := strconv.Atoi(noKerningDistance)
	if err != nil {
		fmt.Println(err)
	}

	for (currTime-counter)*counter <= currDistance {

		counter++
	}

	sum = currTime - counter - counter + 1

	fmt.Println(sum)

}
