package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
)

func main() {
    var calories []int
    var counter = 0

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if (scanner.Text() == "") {
            calories = append(calories, counter)
            counter = 0
        }
        value, _ := strconv.Atoi(scanner.Text())
        counter += value
    }

    fmt.Println(calculateTopThreeSum(calories))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func calculateTopThreeSum(elements []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(elements)))

    return elements[0] + elements[1] + elements[2]
}
