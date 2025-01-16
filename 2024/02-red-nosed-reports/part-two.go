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

		if (checkIfSafe(strings.Split(line, " "), "asc", false) || checkIfSafe(strings.Split(line, " "), "desc", false)) {
			result++
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func checkIfSafe(levels []string, order string, inside bool) bool {
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
			if (!inside) {
				safe := false
				for j := 0; j < len(levels); j++ {
					withoutCurrent := append([]string{}, levels[:j]...)
					withoutCurrent = append(withoutCurrent, levels[j+1:]...)
					if (checkIfSafe(withoutCurrent, order, true)) {
						safe = true
					}
				}
				return safe
			} else {
				return false
			}
		}
	}

	return true
}
