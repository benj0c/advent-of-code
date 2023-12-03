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

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        value := scanner.Text()

        game := strings.Split(value, ":")
        sets := strings.Split(game[1], ";")

        counter := map[string]int {
            "red": 0,
            "green": 0,
            "blue": 0,
        }

        for _, set := range sets {
            cubes := strings.Split(set, ",")

            for _, cube := range cubes {
                cube = strings.TrimSpace(cube)
                cubeValue, _ := strconv.Atoi(strings.Split(cube, " ")[0]);
                cubeColor := strings.Split(cube, " ")[1]

                if (counter[cubeColor] < cubeValue) {
                    counter[cubeColor] = cubeValue
                }
            }
        }

        result += counter["red"] * counter["green"] * counter["blue"]
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
