package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  var result, boxNum int
  var boxes [256][]map[string]int

  file, _ := os.ReadFile(os.Args[1])
  steps := strings.Split(string(file), ",")

  noAdd:
  for _, step := range steps {
    step = strings.ReplaceAll(step, "\n", "")
    if strings.Contains(step, "=") {
      parts := strings.Split(step, "=")
      boxNum = hash(parts[0])
      val, _ := strconv.Atoi(parts[1])
      for _, v := range boxes[boxNum] {
        if v[parts[0]] != 0 {
          v[parts[0]] = val
          continue noAdd
        }
      }
      boxes[boxNum] = append(boxes[boxNum], map[string]int{parts[0]: val})
    } else {
      step = strings.ReplaceAll(step, "-", "")
      boxNum = hash(step)
      for key, val := range boxes[boxNum] {
        if val[step] != 0 {
          boxes[boxNum] = append(boxes[boxNum][:key], boxes[boxNum][key + 1:]...)
        }
      }
    }
  }

  for boxNum, box := range boxes {
    if box == nil { continue }
    for slot, slots := range box {
      for _, focalLength := range slots {
        result += (boxNum + 1) * (slot + 1) * focalLength
      }
    }
    fmt.Println("Box ", boxNum, ": ", box)
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
