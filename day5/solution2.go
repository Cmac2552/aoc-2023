package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type infoHolder struct {
	destStart   int
	sourceStart int
	sourceEnd   int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mapsMap := map[string][]infoHolder{}
	seeds := make([]int, 0)
	currentMap := ""

	for scanner.Scan() {
		text := scanner.Text()
		numberMap := strings.Split(text, " ")
		if len(text) > 5 && text[0:6] == "seeds:" {
			seedStrings := strings.Split(text, " ")
			for i := 1; i < len(seedStrings); i++ {
				seedNumber, err := strconv.Atoi(seedStrings[i])
				if err != nil {
					fmt.Println(err)
				}
				seeds = append(seeds, seedNumber)
			}

		} else if len(strings.Split(text, "-")) > 1 {
			currentMap = text

		} else if text != "" {
			mapRange, err := strconv.Atoi(numberMap[2])
			if err != nil {
				fmt.Println(err)
			}

			destinationStart, err := strconv.Atoi(numberMap[0])
			if err != nil {
				fmt.Println(err)
			}

			sourceStart, err := strconv.Atoi(numberMap[1])
			if err != nil {
				fmt.Println(err)
			}

			mapsMap[currentMap] = append(mapsMap[currentMap], infoHolder{destStart: destinationStart, sourceStart: sourceStart, sourceEnd: sourceStart + mapRange})

		}
	}

	lowestLocation := -1
	for j := 0; j < len(seeds); j += 2 {
		for i := seeds[j]; i < seeds[j]+seeds[j+1]; i++ {
			location :=
				mapConversion(mapsMap["humidity-to-location map:"],
					mapConversion(mapsMap["temperature-to-humidity map:"],
						mapConversion(mapsMap["light-to-temperature map:"],
							mapConversion(mapsMap["water-to-light map:"],
								mapConversion(mapsMap["fertilizer-to-water map:"],
									mapConversion(mapsMap["soil-to-fertilizer map:"],
										mapConversion(mapsMap["seed-to-soil map:"], i)))))))
			if lowestLocation > location || lowestLocation == -1 {
				lowestLocation = location
			}
		}
	}

	fmt.Println(lowestLocation)

}

func mapConversion(myMap []infoHolder, source int) int {
	for i := 0; i < len(myMap); i++ {
		if source >= myMap[i].sourceStart && source < myMap[i].sourceEnd {
			return myMap[i].destStart - myMap[i].sourceStart + source
		}
	}
	return source
}
