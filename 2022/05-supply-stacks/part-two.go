package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    var parseSteps bool
    var stacks = make(map[int][]string)
    var result string

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            parseSteps = true
            continue
        }

        if !parseSteps && string(line[1]) != "1" {
            cols := (len(line) + 1) / 4

            for i := 0; i < cols; i++ {
                crate := string(line[i * 4 + 1])

                if crate != " " {
                    stacks[i + 1] = append(stacks[i + 1], crate)
                }
            }
        }

        if parseSteps {
            params := strings.Split(line, " ")
            q, _ := strconv.Atoi(params[1])
            from, _ := strconv.Atoi(params[3])
            to, _ := strconv.Atoi(params[5])

            temp := append([]string{}, stacks[from][:q]...)
            stacks[from] = stacks[from][q:]
            stacks[to] = append(temp, stacks[to]...)
        }
    }

    for i :=  0; i < len(stacks); i++ {
        result += stacks[i+1][0]
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
