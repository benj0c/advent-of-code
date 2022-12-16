package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func main() {
    var positions = [][2]int{{0, 0}}
    var head = [2]int{0, 0}
    var tail = [2]int{0 ,0}

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        params := strings.Split(scanner.Text(), " ")

        dir := params[0]
        steps, _ := strconv.Atoi(params[1])

        for i:=1; i<=steps; i++ {
            if dir == "R" {
                head = [2]int{head[0]+1, head[1]}
            } else if dir == "U" {
                head = [2]int{head[0], head[1]+1}
            } else if dir == "L" {
                head = [2]int{head[0]-1, head[1]}
            } else if dir == "D" {
                head = [2]int{head[0], head[1]-1}
            }

            if hasToMoveTail(head, tail) {
                tail = moveTail(head, tail)
                positions = append(positions, tail)
            }
        }

    }

    fmt.Println(len(removeDups(positions)))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func hasToMoveTail(head, tail [2]int) bool {
    return (math.Abs(float64(head[1]) - float64(tail[1])) > 1) || (math.Abs(float64(head[0]) - float64(tail[0])) > 1)
}

func moveTail(head, tail [2]int) [2]int {
    xd := head[0] - tail[0]
    yd := head[1] - tail[1]

    if (xd == 0 && yd > 1) {
        return [2]int{tail[0], tail[1]+1}
    } else if (xd == 0 && yd < -1) {
        return [2]int{tail[0], tail[1]-1}
    } else if (yd == 0 && xd > 1) {
        return [2]int{tail[0]+1, tail[1]}
    } else if (yd == 0 && xd < -1) {
        return [2]int{tail[0]-1, tail[1]}

    } else if (xd > 0 && yd > 0) {
        return [2]int{tail[0]+1, tail[1]+1}
    } else if (xd > 0 && yd < 0) {
        return [2]int{tail[0]+1, tail[1]-1}
    } else if (xd < 0 && yd > 0) {
        return [2]int{tail[0]-1, tail[1]+1}
    } else if (xd < 0 && yd < 0) {
        return [2]int{tail[0]-1, tail[1]-1}
    }

    return tail
}

func removeDups(elements [][2]int) (nodups [][2]int) {
    encountered := make(map[[2]int]bool)

    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}
