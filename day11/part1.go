package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	// "slices"
)

func main() {
  var result int
  var meep [][]string
  var colList, rowList []int
  distMap := make(map[[2]string]int)
  locMap := make(map[string][2]int) 
  galMultiplier := 1 

  file, _ := os.ReadFile(os.Args[1])
  rows := strings.Split(string(file), "\n")
  //add galaxy and expand rows
  for x, row := range rows {
    if !strings.Contains(row, "#") {
      rowList = append(rowList, x)
    }
    meep = append(meep, strings.Split(row, ""))
  }

  fmt.Println("rows expanded", rowList)
  //expand columns
  for _, row := range meep {
    for y, pos := range row {
      if pos == "#" {
        colList = append(colList, y)
      }
    }
  }

  fmt.Println("colList made", colList)
//  var offset int
//  for x, row := range meep {
//    offset = 0
//    for y, dot := range row {
//      if slices.Contains(colList, y) { continue }
//      for i := 1; i < galMultiplier; i++ {
//	      meep[x] = append(meep[x], dot)
//	      copy(meep[x][y + 1 + offset:], meep[x][y + offset:])
//	      meep[x][y + offset] = dot
//	      offset++
//      }
//    }
//  }
//  fmt.Println("cols expanded")

  //number galaxies and add their location to map
  var count int
  for x, row := range meep {
    for y, pos := range row {
      if pos == "#" {
        strCount := strconv.Itoa(count)
        meep[x][y] = strCount 
        locMap[strCount] = [2]int{x, y}
        count++
      }
    }
  }

  //map shortest distances
  for fGal, fxy := range locMap {
    for sGal, sxy := range locMap {
      var xD, yD int
      if fGal == sGal { continue }
      fmt.Println("[", fGal, ", ", sGal, "]")
      if fxy[0] > sxy[0] {
        for i := sxy[0] + 1; i < fxy[0]; i++ {
          fmt.Println(i)
          if slices.Contains(rowList, i) {
            xD += galMultiplier
          }
          xD++
          fmt.Println("xD", xD)
        }
      } else {
        for i := fxy[0] + 1; i < sxy[0] - 1; i++ {
          if slices.Contains(rowList, i) {
            xD += galMultiplier
          }
          xD++
        }
      }
      //for y coords
      if fxy[1] > sxy[1] {
        for i := sxy[1] + 1; i < fxy[1] - 1; i++ {
          if slices.Contains(colList, i) {
            yD += galMultiplier
          }
          yD++
        }
      } else {
        for i := fxy[1] + 1; i < sxy[1] - 1; i++ {
          if slices.Contains(colList, i) {
            yD += galMultiplier
          }
          yD++
        }
      }
      if distMap[[2]string{fGal, sGal}] == 0 && distMap[[2]string{sGal, fGal}] == 0 {
        distMap[[2]string{fGal, sGal}] = xD + yD
      }
    }
  }

  fmt.Println(distMap)
  for _, row := range meep {
    fmt.Println(row)
  }

  for _, val := range distMap {
    result += val
  }

  fmt.Println(result)
}
