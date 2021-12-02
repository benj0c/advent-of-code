package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    var commands []string

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value := scanner.Text()
        commands = append(commands, value)
    }

    fmt.Println(parseCommands(commands))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func parseCommands(commands []string) int {
    var(
        horizontal int = 0
        depth int = 0
    )

    for _, command := range commands {
        direction := strings.Split(command, " ")[0]
        value, _ := strconv.Atoi(strings.Split(command, " ")[1])

        switch direction {
            case "forward": horizontal += value
            case "down": depth += value
            case "up": depth -= value
        }
    }

    return horizontal * depth
}
