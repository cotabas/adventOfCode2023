package main

import (
	"fmt"
	"os"
	"strings"
)


func main () {
  coords := make(map[string][2]string)
  
  input, _ := os.ReadFile(os.Args[1])

  lines := strings.Split(string(input), "\n")
  ins := lines[0]

  for _, line := range lines[1:] {
    if line == "" { continue }
    coord := strings.Split(line, " ")[0]
    coords[coord] = [2]string{strings.Split(line, "(")[1][0:3], strings.Split(line, ",")[1][1:4]}
  }

  var starts []string
  for k := range coords {
    if k[2] == 'A' { 
      starts = append(starts, k)
    }
  }

  var dist []int

  for _, start := range starts {
    steps := 0
    sc := 0
    for sc == 0 {
      for _, i := range ins {
        switch i {
        case 'R':
          start = coords[start][1]
        case 'L':
          start = coords[start][0]
        }
        steps += 1
        if start[2] == 'Z' { sc = steps }
      }
    }
    dist = append(dist, sc)
  }

  lcm := dist[0]

  for _, num := range dist[1:] {
    lcm = lcm * num / GCD(lcm, num)
  }
  fmt.Println(dist)
  fmt.Println(lcm)
}

func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

