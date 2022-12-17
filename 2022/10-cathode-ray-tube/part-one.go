package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    var cycles = []int{1}

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        lastValue := cycles[len(cycles)-1]

        if line == "noop" {
            cycles = append(cycles, lastValue)
        } else {
            value, _ := strconv.Atoi(strings.Split(line, " ")[1])
            cycles = append(cycles, []int{lastValue, lastValue+value}...)
        }
    }

    fmt.Println(findSumOfSignalStrengths(cycles))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findSumOfSignalStrengths(cycles []int) (result int) {
    result = 0

    for i:=20; i<len(cycles); i+=40 {
        result += i * cycles[i-1]
    }

    return
}
