package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  var result int
  work := [][]int{}
  
  input, _ := os.ReadFile(os.Args[1])
  numbers := strings.Split(string(input), "\n")

  for _, num := range numbers[:len(numbers) - 1] {
    var cn []int
    ss := strings.Split(num, " ")
    for _, s := range ss {
      cs, _ := strconv.Atoi(s)
      cn = append(cn, cs)
    }
    work = append(work, cn)
  }

  for _, line := range work {
    diff := []int{}
    adds := []int{}

    for i, w := range line {
      if i < len(line) - 1 {
        num := line[i + 1] - w
        diff = append(diff, num)
      }
    }
    fmt.Println("Line:   ", line)
    working := true
    for working {
      fmt.Println(diff)
      next := []int{}
      for i, w := range diff {
        if i < len(diff) - 1 {
          num := diff[i + 1] - w
          next = append(next, num)
        }
      }
      adds = append(adds, diff[len(diff) - 1])
      diff = next
  var grr bool = false

  for _, w := range diff {
    if w != 0 {
      grr = false
      break
    }
    grr = true
  }
      if grr || len(diff) == 1 {
        adds = append(adds, line[len(line) - 1])
        for _, add := range adds {
          result += add
          working = false
        }
        fmt.Println("   ", adds, result)
      }
    }
  }

  fmt.Println(result)
}
