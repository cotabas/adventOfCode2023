package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var energized []string
var eSquares [][2]int

func main() {
  var result int
  file, _ := os.ReadFile(os.Args[1])
  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  for dex, line := range input {
    fmt.Println(dex, ": ", line)
  }

  for i := 0; i < len(input); i++ {
    eSquares = [][2]int{}
    energized = []string{}
    move("E", [2]int{i, -1}, input)
    if result < len(eSquares) { result = len(eSquares) }
    eSquares = [][2]int{}
    energized = []string{}
    move("W", [2]int{i, len(input[i])}, input)
    if result < len(eSquares) { result = len(eSquares) }
  }
  fmt.Println("ew done")
  for i := 0; i < len(input[0]); i++ {
    eSquares = [][2]int{}
    energized = []string{}
    move("S", [2]int{-1, i}, input)
    if result < len(eSquares) { result = len(eSquares) }
    eSquares = [][2]int{}
    energized = []string{}
    move("N", [2]int{len(input), i}, input)
    if result < len(eSquares) { result = len(eSquares) }
  }

  fmt.Println(result)
}

func move(dir string, location [2]int, input []string) {
  var ns, ew int
  x := location[0]
  y := location[1]

  switch dir {
  case "N":
    ns = -1
  case "S":
    ns = 1
  case "E":
    ew = 1
  case "W":
    ew = -1
  }

  if len(input) <= x + ns || x + ns < 0 { return }
  if len(input[0]) <= y + ew || y + ew < 0 { return }
  
  loc := [2]int{x + ns, y + ew}
  xyd := strconv.Itoa(loc[0]) + "," + strconv.Itoa(loc[1]) + dir

  if !slices.Contains(eSquares, loc) {
    eSquares = append(eSquares, loc)
  }

  if slices.Contains(energized, xyd) { return }
  energized = append(energized, xyd)
  

  //fmt.Println(string(input[loc[0]][loc[1]]), " at ", loc)
  switch input[x + ns][y + ew] {
  case '.':
    move(dir, loc, input) 
  case '|':
    if ns == 0 {
      move("N", loc, input)
      move("S", loc, input)
    } else {
      move(dir, loc, input)
    }
  case '-':
    if ew == 0 {
      move("E", loc, input)
      move("W", loc, input)
    } else {
      move(dir, loc, input)
    }
  case '/':
    switch dir {
    case "N":
      move("E", loc, input)
    case "S":
      move("W", loc, input)
    case "E":
      move("N", loc, input)
    case "W":
      move("S", loc, input)
    }
    case '\\':
    switch dir {
    case "N":
      move("W", loc, input)
    case "S":
      move("E", loc, input)
    case "E":
      move("S", loc, input)
    case "W":
      move("N", loc, input)
    }
  }
}
