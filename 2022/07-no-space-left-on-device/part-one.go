package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type Node struct {
  name string
  isDir bool
  size int
  parent *Node
  children map[string]*Node
}

func main() {
  var root = Node{ name: "/", isDir: true, children: make(map[string]*Node) }
  var currentDir *Node

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()

    if strings.HasPrefix(line, "$ cd") {
      path := strings.Split(line, " ")[2]

      if (path == "/") {
        currentDir = &root
      } else if (path == "..") {
        currentDir = (*currentDir).parent
      } else {
        if (*currentDir).children[path] != nil {
          currentDir = (*currentDir).children[path]
        } else {
          newDir := Node{ name: path, isDir: true, parent: currentDir, children: make(map[string]*Node) }
          (*currentDir).children[path] = &newDir
          currentDir = &newDir
        }
      }
    } else if line == "$ ls" {
      continue
    } else {
      params := strings.Split(line, " ")
      itemName := params[1]

      var newNode Node

      if params[0] == "dir" {
        newNode = Node{ name: itemName, isDir: true, parent: currentDir, children: make(map[string]*Node) }
      } else {
        itemSize, _ := strconv.Atoi(params[0])
        newNode = Node{ name: itemName, isDir: false, size: itemSize }
      }

      (*currentDir).children[params[1]] = &newNode
    }
  }

  calculateDirSize(&root)

  fmt.Println(findResult(root, 100000))

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}

func findResult(node Node, threshold int) int {
  var result = 0

  for _, item := range node.children {
    if item.isDir {
      result += findResult(*item, threshold)
    }
  }

  if (node.size <= threshold) {
    result += node.size
  }

  return result
}

func calculateDirSize(node *Node) int {
  for _, item := range node.children {
    if item.isDir {
      node.size += calculateDirSize(item)
    } else {
      node.size += item.size
    }
  }

  return node.size
}
