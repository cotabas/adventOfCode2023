package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const STEPS int = 26501365 
var xLen, yLen int
var been [][2]int
var meep map[[2]int][]int

func main() {
  var result int
  var dots [][2]int
  var start [2]int

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]
  xLen = len(input)

  for _, line := range input { fmt.Println(line) }

  for x, line := range input { 
    yLen = len(line)
    for y, char := range line {
      switch char {
      case 'S':
        start = [2]int{x, y}
        dots = append(dots, [2]int{x, y})
      case '.':
        dots = append(dots, [2]int{x, y})
      }
    }
  }

  meep = make(map[[2]int][]int)
  walk(dots, start, 0)
  result = len(been)
  fmt.Println(result)
}

func walk(dots [][2]int, start [2]int, steps int) {
  if steps == STEPS { return }
  if slices.Contains(meep[start], steps) { return }
  meep[start] = append(meep[start], steps)

  for i := -1; i <= 1; i += 2 {
    next := [2]int{start[0] + i, start[1]} 
    if slices.Contains(dots, oobCheck(next)) {
      if !slices.Contains(been, next) && steps == STEPS - 1 {
        fmt.Println(" Steps: ", steps, " x,y: ", next, " total: ", len(been))
        been = append(been, next) 
      }
      walk(dots, next, steps + 1)
    }
    next = [2]int{start[0], start[1] + i} 
    if slices.Contains(dots, oobCheck(next)) {
      if !slices.Contains(been, next) && steps == STEPS - 1 {
        fmt.Println(" Steps: ", steps, " x,y: ", next, " total: ", len(been))
        been = append(been, next) 
      }
      walk(dots, next, steps + 1)
    }
  }
}

func oobCheck(next [2]int) [2]int {
  check := next
  if next[0] < 0 { check[0] = next[0] % xLen + xLen }
  if next[0] > xLen { check[0] = next[0] % xLen }
  if check[0] == xLen { check[0] = 0 }
  if next[1] < 0 { check[1] = next[1] % yLen + yLen }
  if next[1] > yLen { check[1] = next[1] % yLen }
  if check[1] == yLen { check[1] = 0 }

  return check
}
