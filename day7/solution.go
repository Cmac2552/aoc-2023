package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type handAndWager struct {
	hand  string
	wager string
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	hands := make([][]handAndWager, 7)

	for scanner.Scan() {
		text := scanner.Text()
		currHand := handAndWager{hand: strings.Split(text, " ")[0], wager: strings.Split(text, " ")[1]}
		hand := map[string]int{}

		for c := range currHand.hand {

			hand[string([]rune(currHand.hand)[c])] = hand[string([]rune(currHand.hand)[c])] + 1
		}

		biggestMultiple := 0
		biggerThanOne := 0
		for _, v := range hand {
			if v > biggestMultiple {
				biggestMultiple = v
			}
			if v > 1 {
				biggerThanOne++
			}
		}

		switch biggestMultiple {
		case 5:
			hands[0] = append(hands[0], currHand)
		case 4:
			hands[1] = append(hands[1], currHand)
		case 3:
			if biggerThanOne > 1 {
				hands[2] = append(hands[2], currHand)
			} else {
				hands[3] = append(hands[3], currHand)
			}
		case 2:
			if biggerThanOne > 1 {
				hands[4] = append(hands[4], currHand)
			} else {
				hands[5] = append(hands[5], currHand)
			}
		case 1:
			hands[6] = append(hands[6], currHand)
		}

	}
	counter := 1
	for i := len(hands) - 1; i >= 0; i-- {
		if len(hands[i]) > 0 {
			hands[i] = quickSort(hands[i], 0, len(hands[i])-1)

			for j := len(hands[i]) - 1; j >= 0; j-- {
				wager, err := strconv.Atoi(hands[i][j].wager)
				if err != nil {
					fmt.Println(nil)
				}
				sum += wager * counter
				counter++
			}
		}

	}

	fmt.Println(sum)

}

func partition(arr []handAndWager, low, high int) ([]handAndWager, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if stronger(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []handAndWager, low, high int) []handAndWager {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func stronger(hand1 handAndWager, hand2 handAndWager) bool {
	cardMap := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	cards1 := []rune(hand1.hand)
	cards2 := []rune(hand2.hand)

	for i := 0; i < len(hand1.hand); i++ {
		if cardMap[string(cards1[i])] != cardMap[string(cards2[i])] {
			return cardMap[string(cards1[i])] > cardMap[string(cards2[i])]
		}
	}
	return false

}
