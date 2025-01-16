package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
    "sort"
)

type Monkey struct {
    items []int
    inspects int
    operation string
    test int
    moveTo map[bool]int
}

func main() {
    var monkeys []Monkey
    var index int

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if (strings.HasPrefix(line, "Monkey")) {
            index, _ = strconv.Atoi(strings.Replace(strings.Split(line, " ")[1], ":", "", 1))
            newMonkey := Monkey{items: []int{}, inspects: 0, operation: "", test: 0, moveTo: make(map[bool]int)}
            monkeys = append(monkeys, newMonkey)
        } else if (strings.HasPrefix(line, "  Starting items:")) {
            items := strings.Split(strings.Split(line, "  Starting items: ")[1], ", ")
            monkeys[index].items = sliceAtoi(items)
        } else if (strings.HasPrefix(line, "  Operation:")) {
            operation := strings.Split(strings.Split(line, "  Operation: ")[1], " = ")[1]
            monkeys[index].operation = operation
        } else if (strings.HasPrefix(line, "  Test:")) {
            test, _ := strconv.Atoi(strings.Split(line, "divisible by ")[1])
            monkeys[index].test = test
        } else if (strings.HasPrefix(line, "    If true:")) {
            moveTo, _ := strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
            monkeys[index].moveTo[true] = moveTo
        } else if (strings.HasPrefix(line, "    If false:")) {
            moveTo, _ := strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
            monkeys[index].moveTo[false] = moveTo
        }

    }

    for round:=1; round<21; round++ {
        for i:=0; i<len(monkeys); i++ {
            for len(monkeys[i].items) > 0 {
                item := monkeys[i].items[0]
                monkeys[i].items = monkeys[i].items[1:]
                monkeys[i].inspects++

                item = eval(monkeys[i].operation, item)
                item = int(math.Trunc(float64(item) / 3))
                target := monkeys[i].moveTo[item % monkeys[i].test == 0]
                monkeys[target].items = append(monkeys[target].items, item)
            }
        }
    }

    inspects := getInspects(monkeys)

    fmt.Println(inspects[0] * inspects[1])

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func getInspects(monkeys []Monkey) []int {
    list := []int{}

    for _, monkey := range monkeys {
        list = append(list, monkey.inspects)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(list)))
    return list
}

func sliceAtoi(sa []string) []int {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, _:= strconv.Atoi(a)
        si = append(si, i)
    }
    return si
}

func eval(expression string, old int) int {
    params := strings.Split(strings.Replace(expression, "old", strconv.Itoa(old), -1), " ")

    operation := params[1]
    p1, _ := strconv.Atoi(params[0])
    p2, _ := strconv.Atoi(params[2])

    if operation == "+" {
        return p1 + p2
    } else {
        return p1 * p2
    }
}
