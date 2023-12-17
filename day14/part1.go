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


  for _, line := range input { fmt.Println("og: ", line) }
  repeat := true
  for repeat {
    repeat = false
    for lineDex, line := range input {
      if lineDex == len(input) - 1 { continue }
      lineB := []byte(line)
      for charDex, char := range line {
        if char == '.' && input[lineDex + 1][charDex] == 'O' {
          repeat = true
          lineB[charDex] = 'O'
          inputB := []byte(input[lineDex + 1])
          inputB[charDex] = '.'
          input[lineDex + 1] = string(inputB)
        }
      }
      input[lineDex] = string(lineB)
    }
  }
  for _, line := range input { fmt.Println(line) }

  for dex, line := range input {
    for _, char := range line {
      if char == 'O' { result += len(input) - dex }
    }
  }

  fmt.Println(result)
}
