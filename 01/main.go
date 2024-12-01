package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
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

	total := 0.0

	leftCol := []float64{}
	rightCol := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")

		left, err := strconv.ParseFloat(numbers[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.ParseFloat(numbers[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		leftCol = append(leftCol, left)
		rightCol = append(rightCol, right)
	}

	if len(rightCol) != len(leftCol) {
		log.Fatalln("columns aren't the same size")
	}

	log.Println("Sorting leftCol and rightCol slices")
	sort.Float64s(leftCol)
	sort.Float64s(rightCol)
	log.Println("Slices sorted successfully")

	for i := range leftCol {
		distance := math.Abs(leftCol[i] - rightCol[i])
		log.Println(distance)
		total += distance
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total distance: %f", total)

	similarity := 0.0

	for i := range leftCol {
		initial := leftCol[i]
		numOfOccurences := 0.0
		for j := range rightCol {
			if rightCol[j] == initial {
				numOfOccurences++
			}
		}

		similarity = similarity + (initial * numOfOccurences)
	}

	log.Printf("Similarity: %f", similarity)
}
