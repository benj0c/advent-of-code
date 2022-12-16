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

    fmt.Println(findVisibleTrees(treeMap))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func findVisibleTrees(treeMap [][]int) int {
    var result = len(treeMap) * 2 + (len(treeMap[0]) - 2) * 2

    for i:=1; i<len(treeMap)-1; i++ {
        for j:=1; j<len(treeMap[i])-1; j++ {

            visible := false

            for k:=i-1; k>=0; k-- {
                visible = treeMap[i][j] > treeMap[k][j]

                if (!visible) {
                    break
                }
            }

            if visible {
                result++
                continue
            }

            for k:=i+1; k<len(treeMap); k++ {
                visible = treeMap[i][j] > treeMap[k][j]

                if (!visible) {
                    break
                }
            }

            if visible {
                result++
                continue
            }

            for k:=j+1; k<len(treeMap[i]); k++ {
                visible = treeMap[i][j] > treeMap[i][k]

                if (!visible) {
                    break
                }
            }

            if visible {
                result++
                continue
            }

            for k:=j-1; k>=0; k-- {
                visible = treeMap[i][j] > treeMap[i][k]

                if (!visible) {
                    break
                }
            }

            if visible {
                result++
            }

        }
    }

    return result
}

func sliceAtoi(sa []string) []int {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, _:= strconv.Atoi(a)
        si = append(si, i)
    }
    return si
}
