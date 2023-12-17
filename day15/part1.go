package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  var result int
  file, _ := os.ReadFile(os.Args[1])
  steps := strings.Split(string(file), ",")

  for _, step := range steps {
    step = strings.ReplaceAll(step, "\n", "")
    result += hash(step)
  }

  fmt.Println(result)
}

func hash(step string) int {
  var val int
  for _, char := range step {
    val += int(char)
    val *= 17
    val %= 256
  }
  return val
}
