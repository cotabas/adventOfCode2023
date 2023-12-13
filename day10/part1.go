package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  var result int
  var meep [][]string
  var start [2]int
  locations := make(map[[2]int]string)
  key := make(map[string][2]string)
	  key["|"] = [2]string{"N", "S"}
	  key["-"] = [2]string{"E", "W"}
	  key["L"] = [2]string{"N", "E"}
	  key["J"] = [2]string{"N", "W"}
	  key["7"] = [2]string{"S", "W"}
	  key["F"] = [2]string{"S", "E"}
	  key["S"] = [2]string{"S", "N"}

  file, _ := os.ReadFile(os.Args[1])

  xAxis := strings.Split(string(file), "\n")

  for _, x := range xAxis[:len(xAxis) - 1] {
    meep = append(meep, strings.Split(x, ""))
  }

  for x, row := range meep {
    for y, pipe := range row {
      if pipe == "S" { start = [2]int{x, y} }
    }
  }

  locations[start] = "S"

  var next [2]int = start
  var count int
  var last [2]int

  for next != [2]int{-1, -1} {
      count++
      last = start
      start = next
      next = travel(key, locations, meep, start, last)
  }
  result = count / 2
  fmt.Println(locations)
  fmt.Println(wtf)
  fmt.Println(count)
  fmt.Println(result)
}
var wtf [][2]int
var firstCheck bool = false
func travel(key map[string][2]string, locations map[[2]int]string, meep [][]string, start [2]int, last [2]int) [2]int {
  var next [2]int

  if locations[start] == "S" && firstCheck { return [2]int{-1, -1} }
  firstCheck = true 
  for pipe, loc := range key {
    fmt.Print(pipe, "   ")
    for i := 0; i < 2; i++ {
      switch key[locations[start]][i] {
      case "N":
        if start[0] > 0 && string(meep[start[0] - 1][start[1]]) == pipe && (loc[0] == "S" || loc[1] == "S") && last != [2]int{start[0] - 1, start[1]} {
          next = [2]int{start[0] - 1, start[1]}
          locations[next] = pipe
          break
        }
      case "S":
        if start[0] < len(meep) - 1 && string(meep[start[0] + 1][start[1]]) == pipe && (loc[0] == "N" || loc[1] == "N") && last != [2]int{start[0] + 1, start[1]} {
          next = [2]int{start[0] + 1, start[1]}
          locations[next] = pipe
          break
        }
      case "W":
        if start[1] > 0 && string(meep[start[0]][start[1] - 1]) == pipe && (loc[0] == "E" || loc[1] == "E") && last != [2]int{start[0], start[1] - 1} {
          next = [2]int{start[0], start[1] - 1}
          locations[next] = pipe
          break
        }
      case "E":
        if start[1] < len(meep[start[0]]) - 1 && string(meep[start[0]][start[1] + 1]) == pipe && (loc[0] == "W" || loc[1] == "W") && last != [2]int{start[0], start[1] + 1} {
          next = [2]int{start[0], start[1] + 1}
          locations[next] = pipe
          break
        }
      }
    }
  }
  //fmt.Println("\nwarn~~!", next, locations[next])
  wtf = append(wtf, next)
//  for i := -1; i < 2; i += 2 {
//    if i == -1 && (start[0] == 0 || start[1] == 0) { continue }
//    if i == 1 && (start[0] >= len(meep) - 1 || start[1] >= len(meep[start[0]]) - 1) { continue }
//    if meep[start[0] + i][start[1]] == "S" {
//      next = [2]int{start[0] + i, start[1]}
//    }
//    if meep[start[0]][start[1] + i] == "S" {
//      next = [2]int{start[0], start[1] + i}
//    }
//  }
  return next
}

