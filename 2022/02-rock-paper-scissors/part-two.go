package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var points = map[string]int {
        "A X": 3,
        "A Y": 4,
        "A Z": 8,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 2,
        "C Y": 6,
        "C Z": 7,
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
