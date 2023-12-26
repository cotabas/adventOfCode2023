package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const STEPS int = 64
var been [][2]int

func main() {
  var result int
  var dots [][2]int
  var start [2]int

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  for _, line := range input { fmt.Println(line) }

  for x, line := range input { 
    for y, char := range line {
      switch char {
      case 'S':
        start = [2]int{x, y}
      case '.':
        dots = append(dots, [2]int{x, y})
      }
    }
  }

  walk(dots, start, 0)
  result = len(been) + 1 //for the start
  fmt.Println(result)
}

func walk(dots [][2]int, start [2]int, steps int) {
  if steps == STEPS { return }

  for i := -1; i <= 1; i += 2 {
    next := [2]int{start[0] + i, start[1]} 
    if slices.Contains(dots, next) {
      if !slices.Contains(been, next) && steps == STEPS - 1 {
        fmt.Println(" Steps: ", steps, " x,y: ", next, " total: ", len(been) + 1)
        been = append(been, next) 
      }
      walk(dots, next, steps + 1)
    }
    next = [2]int{start[0], start[1] + i} 
    if slices.Contains(dots, next) {
      if !slices.Contains(been, next) && steps == STEPS - 1 {
        fmt.Println(" Steps:", steps, " x,y:", next, " total:", len(been) + 1)
        been = append(been, next) 
      }
      walk(dots, next, steps + 1)
    }
  }
}

