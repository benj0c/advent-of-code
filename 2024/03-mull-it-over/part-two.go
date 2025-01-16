package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
)

func main() {
	var input string
	var result int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text()
	}

	parts := findParts(input)

	for _, part := range parts {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		expressions := re.FindAll([]byte(part), -1)

		for _, expression := range expressions {
			result += findResult(expression)
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func findParts(line string) [][]byte {
	pattern := `(^.*?don't\(\))|(do\(\).*?don't\(\))|(do\(\).*?$)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAll([]byte(line), -1)

	return matches
}

func findResult(input []byte) int {
	pattern := `(\d+)\D+(\d+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(string(input[:]))

	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])

	return num1 * num2
}
