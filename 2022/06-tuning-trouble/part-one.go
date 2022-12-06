package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
  var result int
  var sequence = 4

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    characters := strings.Split(line, "")

    for i := 0; i < len(characters) - sequence - 1; i++ {
      if allDifferent(characters[i:i+sequence]) {
        result = i + sequence
        break
      }
    }
  }

  fmt.Println(result)

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}

func allDifferent(elements []string) bool {
  var set = make(map[string]bool)

  for _, item := range elements {
    set[item] = true
  }

  return len(elements) == len(set)
}
