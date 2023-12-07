package main

import (
  "fmt"
  "strings"
  "strconv"
)

var race [2]int

func parseRaces(mets []string, set int) {
  for _, met := range mets {
    t, _ := strconv.Atoi(met)
    if t == 0 { continue }
    race[set] = t 
  }
}

func main() {
  var result int = 1

  input = strings.ReplaceAll(input, " ", "")
  td := strings.Split(input, "\n")
  times := strings.Split(strings.Split(td[0], ":")[1], " ")
  dists := strings.Split(strings.Split(td[1], ":")[1], " ")

  parseRaces(times, 0)
  parseRaces(dists, 1)

  fmt.Println(race)

   count := 0
   for i := 0; i < race[0]; i++ {
     if i * (race[0] - i) > race[1] { count += 1 }
   }
   fmt.Println(count)
   result *= count

  fmt.Println(result)
}


var test string = `Time:    71530
Distance:  940200`

var input string = `Time:        52     94     75     94
Distance:   426   1374   1279   1216`
