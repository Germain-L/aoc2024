package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Opening file input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("File opened successfully")

	safeReports := 0
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count += 1
		log.Println("Processing line :", count)
		numbersString := strings.Split(scanner.Text(), " ")

		numbers := []float64{}
		isSafe := true

		for _, n := range numbersString {
			converted, err := strconv.ParseFloat(n, 64)
			if err != nil {
				log.Println(err)
			}

			numbers = append(numbers, converted)
		}

		// first 2 are equal, not safe
		if numbers[0] == numbers[1] {
			isSafe = false
			log.Println("first 2 are eaqual")
			continue
		}

		isIncreasing := numbers[1] > numbers[0]

		for i := range numbers {
			if i == len(numbers)-1 {
				continue
			}

			if isIncreasing {
				// next number should be higher since increasing
				if numbers[i+1] < numbers[i] {
					isSafe = false
					log.Printf("Increasing list, decreasing next : %f < %f", numbers[i+1], numbers[i])
					break
				}
			} else {
				// next number should be smaller since decreasing
				if numbers[i+1] > numbers[i] {
					isSafe = false
					log.Printf("Decreasing list, increasing next : %f > %f", numbers[i+1], numbers[i])
					break
				}
			}

			// Calculate the difference between current and next numbers
			// absolute so we can work with decreasing or increasing differences
			dif := math.Abs(numbers[i] - numbers[i+1])
			if dif < 1 || dif > 3 {
				isSafe = false
				log.Printf("Dif not in range %f", dif)
				break
			}
		}

		if isSafe {
			safeReports += 1
		}

		log.Println(isSafe)
		log.Println(numbers)
		log.Println("")
	}

	log.Println(safeReports)
}
