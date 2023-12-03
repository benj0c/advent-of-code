package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    result := 0

    limit := map[string]int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        value := scanner.Text()

        game := strings.Split(value, ":")
        id, _ := strconv.Atoi(strings.Split(game[0], " ")[1]);
        sets:= strings.Split(game[1], ";")

        possible := true

        for _, set := range sets {
            cubes := strings.Split(set, ",")

            if (!possible) {
                break
            }

            for _, cube := range cubes {
                cube = strings.TrimSpace(cube)
                value, _ := strconv.Atoi(strings.Split(cube, " ")[0]);
                color := strings.Split(cube, " ")[1]

                if (value> limit[color]) {
                    possible = false
                    break
                }
            }
        }

        if (possible) {
            result += id
        }
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
