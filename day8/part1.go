package main

import (
  "fmt"
  "strings"
  "os"
)

func main () {
  var result, steps int
  coords := make(map[string][2]string)
  
  input, _ := os.ReadFile(os.Args[1])

  lines := strings.Split(string(input), "\n")
  ins := lines[0]

  start := "AAA"

  for _, line := range lines[1:] {
    if line == "" { continue }
    coord := strings.Split(line, " ")[0]
    coords[coord] = [2]string{strings.Split(line, "(")[1][0:3], strings.Split(line, ",")[1][1:4]}
  }

  for result == 0 {
    for _, i := range ins {
      switch i {
      case 'R':
        start = coords[start][1]
      case 'L':
        start = coords[start][0]
      }
      steps += 1
      if start == "ZZZ" { result = steps }
    }
  }
  fmt.Println(result)

}

