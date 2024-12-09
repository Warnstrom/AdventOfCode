package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// https://regex101.com/r/5Mbbo5/1
func main() {
	matchString := regexp.MustCompile(`(?m)(?:don't\(\)(?:.|\n)*?do\(\))|(mul\(\d+,\d+\))`)
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(bufio.NewReader(file))
	instructions := 0
	for scanner.Scan() {
		matches := matchString.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			if len(match) < 3 {
				continue
			}
			a, err1 := strconv.Atoi(match[1])
			b, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				continue
			}
			fmt.Printf("mul(%d,%d)\n", a, b)
			instructions += a * b
		}
	}
	fmt.Println(instructions)
}
