package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
  var result int

  file, _ := os.ReadFile(os.Args[1])
  notes := strings.Split(string(file), "\n\n")

  for _, note := range notes {
    lines := strings.Split(note, "\n")
    if lines[len(lines) - 1] == "" { lines = lines[:len(lines) - 1] }

    vMap := make(map[int][]string)
    for _, line := range lines {
      //make vertical map
      for vDex, char := range line {
        vMap[vDex] = append(vMap[vDex], string(char))
      }
    }
    //switch vertical to horizontal
    var keys []int
    for key := range vMap {
      keys = append(keys, key)
    }
    sort.Ints(keys)
    vLines := []string{}
    for _, l := range keys {
      line := strings.Join(vMap[l], "")
      vLines = append(vLines, line)
    }

    fmt.Println(vLines)

  fmt.Println(result)
    result += findMirror(lines) * 100
  fmt.Println(result)
    result += findMirror(vLines)
  }

  fmt.Println(result)
  //fmt.Println(findMirror([]string{"d", "1","2","3","4", "4", "3", "2", "1"}))
}


func findMirror(grid []string) int {
  for dex := range grid {
    if dex == 0 { continue }
    
    above := grid[:dex]
    below := grid[dex:]


    var rev []string
    for dex := range above {
      rev = append(rev, above[len(above) - 1 - dex])
    }
    above = rev

    if len(above) >= len(below) {
      above = above[:len(below)]
    } else {
      below = below[:len(above)]
    }

    fmt.Println(above, "                                      ", below)
    for i, al := range above {
      if i > len(below) - 1 { continue }
      if al != below[i] { break }
      if i == len(above) - 1 { 
fmt.Println("wiener!", dex)
        return dex }
    }
  }
  return 0
}
