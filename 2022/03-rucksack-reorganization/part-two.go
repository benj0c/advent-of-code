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

  var compartments [][]string

  for scanner.Scan() {
    rucksack := scanner.Text()

    compartments = append(compartments, strings.Split(rucksack, ""))

    if (len(compartments) == 3) {
      commonItems = commonItems + strings.Join(intersection(intersection(compartments[0], compartments[1]), compartments[2]), "")

      compartments = [][]string(nil)
    }

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

func getItemValue(char byte) int {
  if (char < 97) {
    return int(char) - 38
  } else {
    return int(char) - 96
  }
}
