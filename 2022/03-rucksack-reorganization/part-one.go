package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
  var commonItems string
  var result int = 0

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    rucksack := scanner.Text()
    compartments := [][]string {
      strings.Split(rucksack[0:(len(rucksack) / 2)], ""),
      strings.Split(rucksack[(len(rucksack) / 2):len(rucksack)], ""),
    }

    commonItems = commonItems + strings.Join(intersection(compartments[0], compartments[1]), "")
  }

  for _, item := range commonItems {
    result += getItemValue(byte(item))
  }

  fmt.Println(result)

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}

func intersection(s1, s2 []string) (inter []string) {
    hash := make(map[string]bool)

    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        if hash[e] {
            inter = append(inter, e)
        }
    }

    inter = removeDups(inter)
    return
}

func removeDups(elements []string) (nodups []string) {
    encountered := make(map[string]bool)

    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}

func getItemValue(item byte) int {
  if (item < 97) {
    return int(item) - 38
  } else {
    return int(item) - 96
  }
}
