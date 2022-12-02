package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var points = map[string]int {
        "A X": 4,
        "A Y": 8,
        "A Z": 3,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 7,
        "C Y": 2,
        "C Z": 6,
    }

    var result int

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result += points[scanner.Text()]
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
