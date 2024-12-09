package main

import (
	"bufio"
	"fmt"
	"os"
)

// [X  M  A  S ]
// [88 77 65 83]
func main() {
	file, _ := os.Open("sample.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	//wordCheck := ""
	lineArray := make([][]byte, 0)
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineArray = append(lineArray, append([]byte(nil), line...))
	}
	for i := 0; i < len(lineArray); i++ {
		for j := 1; j < len(lineArray[i]); j++ {
			if lineArray[i][j] == 88 {
				wordCount += checkStraightForward(lineArray[i], j)
				wordCount += checkStraightDownwards(lineArray, j, lineNumber)
				if j >= 3 {
					wordCount += checkStraightBackward(lineArray[i], j)
				}
			}
		}
		lineNumber++
	}
	fmt.Println(lineNumber)
	fmt.Println(wordCount)
}

func checkStraightForward(line []byte, position int) int {
	if position <= 137 {
		if line[position+1] == 77 && line[position+2] == 65 && line[position+3] == 83 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func checkStraightBackward(line []byte, position int) int {
	if line[position-1] == 77 && line[position-2] == 65 && line[position-3] == 83 {
		return 1
	} else {
		return 0
	}
}

func checkStraightUpwards() {}
func checkStraightDownwards(line [][]byte, position int, lineNumber int) int {
	if lineNumber < 137 {
		if line[lineNumber+1][position] == 77 && line[lineNumber+2][position] == 65 && line[lineNumber+3][position] == 83 {
			fmt.Println(lineNumber, position)
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func checkTopLeft()     {}
func checkTopRight()    {}
func checkBottomLeft()  {}
func checkBottomRight() {}
