package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

type fn func([]string) string

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

    fmt.Println(findLifeSupportRating(binaryList))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findLifeSupportRating(binaryList []string) int64 {
    oxygenRating := getRating(binaryList, findMostCommon)
    CO2Rating := getRating(binaryList, findLeastCommon)

    return oxygenRating * CO2Rating
}

func getRating(binaryList []string, findCommon fn) int64 {
    numberOfBits := len(binaryList[0])
    resultList := binaryList

    common := findCommon(parseList(resultList))

    for index := 0; index < numberOfBits; index++ {
        var temp []string

        for _, item := range resultList {
            i, _ := strconv.Atoi(string(item[index]))
            j, _ := strconv.Atoi(string(common[index]))
            if (i == j) {
                temp = append(temp, item)
            }
        }

        resultList = temp
        common = findCommon(parseList(resultList))

        if (len(resultList) == 1) {
            break
        }

    }

    result, _ := strconv.ParseInt(resultList[0], 2, 64)

    return result
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

func findMostCommon(list []string) string {
    var result string

    for _, item := range list {
        bits0 := strings.Count(item, "0")
        bits1 := strings.Count(item, "1")

        if (bits0 > bits1) {
            result = result + "0"
        } else {
            result = result + "1"
        }
    }

    return result
}

func findLeastCommon(list []string) string {
    var result string

    for _, item := range list {
        bits0 := strings.Count(item, "0")
        bits1 := strings.Count(item, "1")

        if (bits0 > bits1) {
            result = result + "1"
        } else {
            result = result + "0"
        }
    }

    return result
}
