package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  var result int

  file, _ := os.ReadFile(os.Args[1])
  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  dirs := []string{ "N", "S", "E", "W" }

  count := 0
  for count <= 1000000000 {
    if count % 1000000 == 0 { fmt.Println(count) }
    for _, dir := range dirs {
      input = tilt(input, dir)
    }
    count++
  }

  for dex, line := range input {
    for _, char := range line {
      if char == 'O' { result += len(input) - dex }
    }
  }

  fmt.Println(result)
}

func tilt(input []string, dir string) []string {
  d := 0
  ew := 0
  switch dir {
  case "N":
    d = 1
  case "S":
    d = -1
  case "E":
    ew = -1
  case "W":
    ew = 1
  }
  repeat := true
  for repeat {
    repeat = false
    for lineDex, line := range input {
      if dir == "N" && lineDex == len(input) - 1 { continue }
      if dir == "S" && lineDex == 0 { continue }
      lineB := []byte(line)
      for charDex, char := range line {
        if dir == "E" && charDex == 0 { continue }
        if dir == "W" && charDex == len(line) - 1 { continue }
        if char == '.' && input[lineDex + d][charDex + ew] == 'O' {
          repeat = true
          lineB[charDex] = 'O'
          if dir == "N" || dir == "S" {
            inputB := []byte(input[lineDex + d])
            inputB[charDex] = '.'
            input[lineDex + d] = string(inputB)
          } else { 
            lineB[charDex + ew] = '.'
          }

        }
      }
      input[lineDex] = string(lineB)
    }
  }
  return input
}
