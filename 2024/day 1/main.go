package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var leftList, rightList []int

	for scanner.Scan() {
		num1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Invalid number: %v", err)
		}

		if !scanner.Scan() {
			log.Fatal("Incomplete pair of numbers")
		}

		num2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Invalid number: %v", err)
		}

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	sum := 0
	similarityScore := 0

	rightListMap := make(map[int]int)
	for _, v := range rightList {
		rightListMap[v]++
	}

	for i := 0; i < len(leftList); i++ {
		occurrences := rightListMap[leftList[i]]
		similarityScore += leftList[i] * occurrences

		diff := leftList[i] - rightList[i]
		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}

	fmt.Println("Similarity Score: ", similarityScore)
	fmt.Println("Sum of Differences: ", sum)
}
