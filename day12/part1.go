package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  var result int

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")

  for _, line := range input[:len(input) - 1] {
    both := strings.Split(line, " ")
    record := both[0]
    
    conditions := strings.Split(both[1], ",")
    var cond []int
    for _, condition := range conditions {
      c, _ := strconv.Atoi(condition)
      cond = append(cond, c)
    }

    fmt.Println(record, "  -  ", cond)
  }


  fmt.Println(result)
}

func readLine(line string) []int {

  var charCount int
  var work []int
  for _, char := range line {
    if char == '#' { 
      charCount++
    } else {
      if charCount > 0 { 
        work = append(work, charCount)
        charCount = 0
      }
    }
  }
  if charCount > 0 { work = append(work, charCount) }

  return work
}
