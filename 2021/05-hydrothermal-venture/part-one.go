package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

type Point struct {
    x int
    y int
}

type Line struct {
    start Point
    end Point
}

func main() {
    var points []Point

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        p := strings.Split(strings.Replace(scanner.Text(), " -> ", ",", 1), ",")
        x1, _ := strconv.Atoi(p[0])
        y1, _ := strconv.Atoi(p[1])
        x2, _ := strconv.Atoi(p[2])
        y2, _ := strconv.Atoi(p[3])

        if (x1 == x2 || y1 == y2) {
            line := Line{ Point { x1, y1 }, Point { x2, y2 } }
            points = append(points, generateLinePoints(line)...)
        }
    }

    fmt.Println(countDuplicates(points))

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func generateLinePoints(line Line) []Point {
    var points []Point

    if (line.start.x == line.end.x) {
        points = generateVerticalLinePoints(line)
    } else {
        points = generateHorizontalLinePoints(line)
    }

    return points
}

func generateVerticalLinePoints(line Line) []Point {
    var points []Point

    if (line.start.y < line.end.y) {
        for i := line.start.y; i <= line.end.y; i++ {
            points = append(points, Point{ line.start.x, i })
        }
    } else {
        for i := line.start.y; i >= line.end.y; i-- {
            points = append(points, Point{ line.start.x, i })
        }
    }

    return points
}

func generateHorizontalLinePoints(line Line) []Point {
    var points []Point

    if (line.start.x < line.end.x) {
        for i := line.start.x; i <= line.end.x; i++ {
            points = append(points, Point{ i, line.start.y })
        }
    } else {
        for i := line.start.x; i >= line.end.x; i-- {
            points = append(points, Point{ i, line.start.y })
        }
    }

    return points
}

func countDuplicates(points []Point) int {
    var counter int
    var matrix = make(map[Point]int, 0)

    for i := 0; i <len(points); i++ {
        matrix[points[i]] += 1

        if (matrix[points[i]] == 2) {
            counter++
        }
    }

    return counter
}
