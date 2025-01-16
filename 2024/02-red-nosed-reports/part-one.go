package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var result int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if (checkIfSafe(strings.Split(line, " "), "asc") || checkIfSafe(strings.Split(line, " "), "desc")) {
			result++
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func checkIfSafe(levels []string, order string) bool {
	for i := 1; i < len(levels); i++ {
		var diff float64
		prev, _ := strconv.Atoi(levels[i-1])
		current, _ := strconv.Atoi(levels[i])
		if (order == "desc") {
			diff = float64(prev - current)
		} else {
			diff = float64(current - prev)
		}
		if (diff < 1 || diff > 3) {
			return false
		}
	}

	return true
}
