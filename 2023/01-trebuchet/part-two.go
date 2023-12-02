package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    var tokens = map[string]int{
        "1":        1,
        "2":        2,
        "3":        3,
        "4":        4,
        "5":        5,
        "6":        6,
        "7":        7,
        "8":        8,
        "9":        9,
        "one":      1,
        "two":      2,
        "three":    3,
        "four":     4,
        "five":     5,
        "six":      6,
        "seven":    7,
        "eight":    8,
        "nine":     9,
    }
    result := 0

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        value := scanner.Text()

        firstDigit := tokens[findFirstTokenInString(value, tokens)]
        secondDigit := tokens[reverse(findFirstTokenInString(reverse(value), map2(tokens, func(s string) string { return reverse(s) })))]

        result += firstDigit * 10 + secondDigit
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findFirstTokenInString(str string, tokens map[string]int) string {
    var result string = ""

    for i:= 0; i < len(str) && result == ""; i++ {
        line := str[i:]

        for token := range tokens {
            if strings.HasPrefix(line, token) {
                result = token
                break
            }
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

func map2(data map[string]int, f func(string) string) map[string]int {
    mapped := make(map[string]int, len(data))

    for k, v := range data {
        mapped[f(k)] = v
    }

    return mapped
}
