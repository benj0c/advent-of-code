package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var report []int

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        report = append(report, value)
    }

    fmt.Println(countIncrease(prepareSums(report)))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func prepareSums(array []int) []int {
    var sums []int

    for i := 0; i < len(array) - 2; i++ {
        sum := array[i] + array[i + 1] + array[i + 2]
        sums = append(sums, sum)
    }

    return sums
}

func countIncrease(array []int) int {
    var counter = 0

    for i := 1; i < len(array); i++ {
        if (array[i] > array[i - 1]) {
            counter++;
        }
    }

    return counter
}
