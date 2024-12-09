package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func safe(report []int) bool {
	var increasing = report[0] < report[len(report)-1]
	for i := 1; i < len(report); i++ {
		if increasing && (report[i] <= report[i-1] || report[i]-report[i-1] > 3) ||
			!increasing && (report[i] >= report[i-1] || report[i-1]-report[i] > 3) {
			return false
		}
	}
	return true
}

func convertStringArrayToIntArray(arr []string) []int {
	values := make([]int, len(arr))
	for i, field := range arr {
		convertedValue, _ := strconv.Atoi(field)
		values[i] = convertedValue
	}
	return values
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	var safes [2]int
	var scanner = bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		var report []int
		reportString := strings.Fields(scanner.Text())
		report = convertStringArrayToIntArray(reportString)
		if safe(report) {
			safes[0]++
			safes[1]++
		} else {
			for i, _ := range report {
				var r = make([]int, len(report))
				copy(r, report)
				if safe(append(r[0:i], r[i+1:]...)) {
					safes[1]++
					break
				}
			}
		}
	}
	fmt.Println(safes)
}
