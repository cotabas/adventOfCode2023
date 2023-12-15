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
    fmt.Println(note)

    lines := strings.Split(note, "\n")

    start := 0
    if lines[len(lines) - 1] == "" { lines = lines[:len(lines) - 1] }
    vMap := make(map[int][]string)
    for dex, line := range lines {
      //make vertical map
      for vDex, char := range line {
        vMap[vDex] = append(vMap[vDex], string(char))
      }

      if dex == 0 { continue }
      //find horizontal mirrors
      if line == lines[dex - 1] { 
        count := 0
        for i := 0; dex + i <= len(lines) - 1; i++ {
          if lines[dex - 1 - i] == lines[dex + i] { 
            count++
            if start == 0 { start = dex - 1 - i }
            fmt.Println("mirror line horizontal: [", dex - 1 - i, ", ", dex + i, "]")
            if (dex - 1 - i == 0 || dex + i == len(lines) - 1) && count >= (len(lines) - 2) / 2 { 
              fmt.Println("wiiner!")
              result += 100 * (start + 1)
            } 
          }
          if dex - 1 - i == 0 { break }
        }
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
      //fmt.Println(vMap[l], dex)
      line := strings.Join(vMap[l], "")
      vLines = append(vLines, line)
    }

    //find vertical mirros
    start = 0
    for dex, line := range vLines {
      count := 0
      if dex == 0 { continue }
      if line == vLines[dex - 1] { 
        for i := 0; dex + i <= len(vLines) - 1; i++ {
          if vLines[dex - 1 - i] == vLines[dex + i] { 
            count++
            if start == 0 { start = dex + i }
            fmt.Println("mirror line vertical: [", dex - 1 - i, ", ", dex + i, "]")
            if (dex - 1 - i == 0 || dex + i == len(lines) - 1) && count >= (len(lines) - 2) / 2 { 
              fmt.Println("wiiner!")
              result += start
            } 
          }
          if dex - 1 - i == 0 { break }
        }
      }
    }

  }

  fmt.Println(result)
}
