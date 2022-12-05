package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

func main() {
    var result int

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ",")

        sections := [][]int {
            sliceAtoi(strings.Split(line[0], "-")),
            sliceAtoi(strings.Split(line[1], "-")),
        }

        sort.Slice(sections, func(i, j int) bool {
            return sections[i][0] < sections[j][0]
        })

        if fullyContain(sections[0], sections[1]) {
            result++
        }
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func fullyContain(s1, s2 []int) bool {
    return (s1[0] <= s2[0] && s1[1] >= s2[1]) || s1[0] == s2[0]
}

func sliceAtoi(sa []string) []int {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, _:= strconv.Atoi(a)
        si = append(si, i)
    }
    return si
}
