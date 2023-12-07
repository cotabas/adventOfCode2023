package main

import (
  "fmt"
  "strings"
  "strconv"
)

var races [][2]int

func parseRaces(mets []string, set int) {
  i := 0
  for _, met := range mets {
    t, _ := strconv.Atoi(met)
    if t == 0 { continue }
    races[i][set] = t 
    i += 1
  }
}

func main() {
  var result int = 1
  races = make([][2]int, 4) // not good

  td := strings.Split(input, "\n")
  times := strings.Split(strings.Split(td[0], ":")[1], " ")
  dists := strings.Split(strings.Split(td[1], ":")[1], " ")

  parseRaces(times, 0)
  parseRaces(dists, 1)

  fmt.Println(races)

 for _, race := range races {
   count := 0
   for i := 0; i < race[0]; i++ {
     if i * (race[0] - i) > race[1] { count += 1 }
   }
   fmt.Println(count)
   result *= count
 }

  fmt.Println(result)
}


var test string = `Time:      7  15   30
Distance:  9  40  200`

var input string = `Time:        52     94     75     94
Distance:   426   1374   1279   1216`
