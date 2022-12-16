package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    var treeMap [][]int

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        treeMap = append(treeMap, sliceAtoi(strings.Split(scanner.Text(), "")))
    }

    fmt.Println(findHighestScore(treeMap))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findHighestScore(treeMap [][]int) int {
    var scores = []int{}

    for i:=1; i<len(treeMap)-1; i++ {
        for j:=1; j<len(treeMap[i])-1; j++ {

            var visible = [4]int{0, 0, 0, 0}

            for k:=i-1; k>=0; k-- {
                visible[0]++

                if treeMap[i][j] <= treeMap[k][j] {
                    break
                }
            }

            for k:=i+1; k<len(treeMap); k++ {
                visible[1]++

                if treeMap[i][j] <= treeMap[k][j] {
                    break
                }
            }

            for k:=j+1; k<len(treeMap[i]); k++ {
                visible[2]++

                if treeMap[i][j] <= treeMap[i][k] {
                    break
                }
            }

            for k:=j-1; k>=0; k-- {
                visible[3]++

                if treeMap[i][j] <= treeMap[i][k] {
                    break
                }
            }

            scores = append(scores, visible[0] * visible[1] * visible[2] * visible[3])

        }
    }

    return findMax(scores)
}

func findMax(elements []int) int {
    max := elements[0]
    for _, e := range elements {
        if (e > max) {
            max = e
        }
    }

    return max
}

func sliceAtoi(sa []string) []int {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, _:= strconv.Atoi(a)
        si = append(si, i)
    }
    return si
}
