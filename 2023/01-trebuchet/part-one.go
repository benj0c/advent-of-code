package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
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

        firstDigit := findFirstDigitInString(value)
        secondDigit := findFirstDigitInString(reverse(value))

        result += firstDigit * 10 + secondDigit
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findFirstDigitInString(str string) int {
    var result int

    found := false
    i := 0

    for !found {
        digit, err := strconv.Atoi(string(str[i]))
        if err == nil {
            found = true
            result = digit
        } else {
            i++
        }
    }

    return result
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
