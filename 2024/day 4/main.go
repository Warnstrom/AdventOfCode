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
	//wordCheck := ""
	lineArray := make([][]byte, 0)
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineArray = append(lineArray, append([]byte(nil), line...))
	}
	for i := 0; i < len(lineArray); i++ {
		for j := 0; j < len(lineArray[i]); j++ {
			if lineArray[i][j] == 88 {
				wordCount += checkStraightForward(lineArray[i], j)
				wordCount += checkStraightBackward(lineArray[i], j)
				wordCount += checkStraightDownwards(lineArray, j, i)
				wordCount += checkStraightUpwards(lineArray, j, i)
				wordCount += checkBottomRight(lineArray, j, i)
				wordCount += checkBottomLeft(lineArray, j, i)
				wordCount += checkTopLeft(lineArray, j, i)
				wordCount += checkTopRight(lineArray, j, i)
			}
		}
	}
	fmt.Println(wordCount)
}

func checkStraightForward(line []byte, position int) int {
	if position+3 < len(line) {
		if line[position+1] == 77 && line[position+2] == 65 && line[position+3] == 83 {
			return 1
		}
	}
	return 0
}

func checkStraightBackward(line []byte, position int) int {
	if position >= 3 {
		if line[position-1] == 77 && line[position-2] == 65 && line[position-3] == 83 {
			return 1
		}
	}
	return 0
}

func checkStraightUpwards(line [][]byte, position int, lineNumber int) int {
	if lineNumber >= 3 {
		fmt.Println(position, lineNumber, line[lineNumber][position])
		if line[lineNumber-1][position] == 77 && line[lineNumber-2][position] == 65 && line[lineNumber-3][position] == 83 {
			return 1
		}
	}
	return 0
}

func checkStraightDownwards(line [][]byte, position int, lineNumber int) int {
	if lineNumber <= 136 {
		if line[lineNumber+1][position] == 77 && line[lineNumber+2][position] == 65 && line[lineNumber+3][position] == 83 {
			return 1
		}
	}
	return 0
}

func checkTopLeft(line [][]byte, position int, lineNumber int) int {
	if position >= 3 && lineNumber >= 3 {
		if line[lineNumber-1][position-1] == 77 && line[lineNumber-2][position-2] == 65 && line[lineNumber-3][position-3] == 83 {
			return 1
		}
	}
	return 0
}
func checkTopRight(line [][]byte, position int, lineNumber int) int {
	if position < 137 && lineNumber >= 3 {
		if line[lineNumber-1][position+1] == 77 && line[lineNumber-2][position+2] == 65 && line[lineNumber-3][position+3] == 83 {
			return 1
		}
	}
	return 0
}

func checkBottomLeft(line [][]byte, position int, lineNumber int) int {
	if position >= 3 && lineNumber < 137 {
		if line[lineNumber+1][position-1] == 77 && line[lineNumber+2][position-2] == 65 && line[lineNumber+3][position-3] == 83 {
			return 1
		}
	}
	return 0
}

func checkBottomRight(line [][]byte, position int, lineNumber int) int {
	if position < 137 && lineNumber < 137 {
		if line[lineNumber+1][position+1] == 77 && line[lineNumber+2][position+2] == 65 && line[lineNumber+3][position+3] == 83 {
			return 1
		}
	}
	return 0
}
