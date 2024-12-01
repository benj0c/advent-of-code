package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"math"
)

func main() {
	var left []int
	var right []int
	var result int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		locationIds := strings.Split(line, "   ")

		l, _ := strconv.Atoi(locationIds[0])
		r, _ := strconv.Atoi(locationIds[1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	for i := 0; i < len(left); i++ {
		distance := math.Abs(float64(left[i] - right[i]))
		result += int(distance)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
