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
        aim int = 0
        horizontal int = 0
        depth int = 0
    )

    for _, command := range commands {
        direction := strings.Split(command, " ")[0]
        value, _ := strconv.Atoi(strings.Split(command, " ")[1])

        switch direction {
            case "forward": {
                horizontal += value
                depth += aim * value
            }
            case "down": aim += value
            case "up":  aim -= value
        }
    }

    return horizontal * depth
}
