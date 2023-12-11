package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  var result int
  var intLine []int
  input, _ := os.ReadFile(os.Args[1])
  lines := strings.Split(string(input), "\n")

  for _, line := range lines[:len(lines) - 1] {
    intLine = nil 

    fmt.Println("Line:   ", line)
    s := strings.Split(line, " ")
    for _, num := range s {
      add, _ := strconv.Atoi(num)
      intLine = append(intLine, add)
    }

    fmt.Println(intLine, "   ", result)
    result += work(intLine) + intLine[len(intLine) - 1]
    fmt.Println(intLine, "   ", result)

  }
  fmt.Println(result)
}

func work(intLine []int) int {
  var grr bool = false

  for _, w := range intLine {
    if w != 0 {
      grr = false
      break
    }
    grr = true
  }

  if grr || len(intLine) == 1 { return 0 }

  var next []int
  for dex, i := range intLine {
    if dex < len(intLine) - 1 {
      fmt.Println(intLine[dex + 1], "  -  ", i)
      next = append(next, intLine[dex + 1] - i)
    }
  }
  

  fmt.Println(next, "  --  ", " + ", next[len(next) - 1])
  return work(next) + next[len(next) - 1]
}

