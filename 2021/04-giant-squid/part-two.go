package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    var (
        fileByLine []string
        numbers []int
        boards [][][]int
    )

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        binary := scanner.Text()
        fileByLine = append(fileByLine, binary)
    }

    numbers = parseNumbers(fileByLine[0])
    boards = parseBoards(fileByLine[2:])

    score := calculateScore(findLastBingo(numbers, boards))
    fmt.Println(score)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func calculateScore(numbers []int, board [][]int) int {
    sum := 0

    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            _, found := Find(numbers, board[i][j])
            if (!found) {
                sum += board[i][j]
            }
        }
    }

    return sum * numbers[len(numbers) - 1]
}

func findLastBingo(numbers []int, boards [][][]int) ([]int, [][]int) {
    var (
        drawnNumbers []int
        winBoard [][]int
        winBoardHistory []int
    )


    for i := 5; i <= len(numbers); i++ {
        tempNumbers := numbers[:i]

        Board:
        for i := 0; i < len(boards); i++ {
            _, alreadyWon := Find(winBoardHistory, i)
            if (alreadyWon) {
                continue
            }

            BoardRow:
            for j := 0; j < 5; j++ {
                for k := 0; k < 5; k++ {
                    _, found := Find(tempNumbers, boards[i][j][k])

                    if (!found) {
                       continue BoardRow
                    }
                }

                drawnNumbers = tempNumbers
                winBoard = boards[i]
                winBoardHistory = append(winBoardHistory, i)
                continue Board
            }

            BoardColumn:
            for j := 0; j < 5; j++ {
                for k := 0; k < 5; k++ {
                    _, found := Find(tempNumbers, boards[i][k][j])

                    if (!found) {
                       continue BoardColumn
                    }
                }

                drawnNumbers = tempNumbers
                winBoard = boards[i]
                winBoardHistory = append(winBoardHistory, i)
                continue Board
            }
        }
    }

    return drawnNumbers, winBoard
}

func parseBoards(fileByLine []string) [][][]int {
    var boards [][][]int

    for i := 0; i < len(fileByLine); i+=6 {
        var board [][]int

        for j := 0; j < 5; j++ {
            var numbers []int
            numbersStrArray := strings.Fields(fileByLine[i+j])

            for _, k := range numbersStrArray {
                number, _:= strconv.Atoi(string(k))
                numbers = append(numbers, number)
            }

            board = append(board, numbers)
        }

        boards = append(boards, board)
    }

    return boards
}

func parseNumbers(numbersStr string) []int {
    var numbers []int
    numbersStrArray := strings.Split(numbersStr, ",")

    for _, i := range numbersStrArray {
        number, _:= strconv.Atoi(string(i))
        numbers = append(numbers, number)
    }

    return numbers
}

func Find(slice []int, val int) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}
