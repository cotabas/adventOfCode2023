package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  var result int

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  for _, line := range input { fmt.Println(line) }

  fmt.Println(result)
}
