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

    draw(cycles)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func draw(cycles []int) {
    for i:=0; i<6; i++ {
        for j:=1; j<41; j++ {
            value := cycles[40*i+j-1]
            sprite := []int{value, value+1, value+2}
            if contains(sprite, j) {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println("")
    }
}

func contains(elems []int, v int) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}
