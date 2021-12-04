package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    var binaryList []string

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        binary := scanner.Text()
        binaryList = append(binaryList, binary)
    }

    fmt.Println(findPowerConsumption(binaryList))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findPowerConsumption(binaryList []string) int64 {
    var(
        gamaStr string
        epsilonStr string
    )

    for _, item := range parseList(binaryList) {
        bits0 := strings.Count(item, "0")
        bits1 := strings.Count(item, "1")

        if (bits0 > bits1) {
            gamaStr = gamaStr + "0"
            epsilonStr = epsilonStr + "1"
        } else {
            gamaStr = gamaStr + "1"
            epsilonStr = epsilonStr + "0"
        }
    }

    gama, _ := strconv.ParseInt(gamaStr, 2, 64)
    epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)

    return gama * epsilon
}

func parseList(binaryList []string) []string {
    list := make([]string, len(binaryList[0]))

    for i := 0; i < len(binaryList); i++ {
        for j := 0; j < len(binaryList[i]); j++ {
            list[j] = list[j] + string(binaryList[i][j])
        }
    }

    return list
}
