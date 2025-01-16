package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
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

		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		expressions := re.FindAll([]byte(line), -1)

		for _, expression := range expressions {
			result += findResult(expression)
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func findResult(input []byte) int {
	pattern := `(\d+)\D+(\d+)`
	re := regexp.MustCompile(pattern)

	// FindStringSubmatch extracts matches including the whole match and groups
	matches := re.FindStringSubmatch(string(input[:]))

	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])

	return num1 * num2
}
