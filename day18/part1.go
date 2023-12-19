package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
  var result int
  meep := make(map[int][]int)
  curLoc := []int{0, 0}

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  for _, line := range input {
    fmt.Println(line)
    dir := strings.Split(line, " ")
    var ud, rl int
    switch dir[0] {
    case "R":
      rl = 1
    case "L":
      rl = -1
    case "U":
      ud = -1
    case "D":
      ud = 1
    }
    steps, _ := strconv.Atoi(dir[1])

    for i := 1; i <= steps; i++ {
      curLoc = []int{curLoc[0] + ud, curLoc[1] + rl}
      meep[curLoc[0]] = append(meep[curLoc[0]], curLoc[1])
    }


  }

  // make a list of holes
  // sort and count differences between highest and next highest
  // map[xCoord][]int of y coords

  for d, m := range meep { 
    f := true
    firstSect := true
    first := 0
    sort.IntSlice.Sort(m)
    //
    pr := result
    //
    result += len(m)
    for dex, coord := range m {
      if f {
        if dex + 1 < len(m) && coord + 1 == m[dex + 1] { continue }
        first = coord
        f = false
      } else {
        if firstSect {
          fmt.Println(coord, " - ", first, " = ", coord - first)
          firstSect = false 
          result += coord - first - 1
          if dex + 1 < len(m) && coord + 1 != m[dex + 1] {
            first = coord
          } else { f = true }
          continue
        }
        if dex + 1 < len(m) && coord + 1 != m[dex + 1] {
          first = coord
        } else { f = true }
        firstSect = true
      }
    }
    fmt.Println(d, m, "total:", result - pr)
  }

  fmt.Println(meep)
  fmt.Println(result)
}
