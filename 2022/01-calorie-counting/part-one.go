package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var result int
    var counter = 0

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if (scanner.Text() == "") {
            if (result < counter) {
                result = counter
            }
            counter = 0
        }
        value, _ := strconv.Atoi(scanner.Text())
        counter += value
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
