package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
  part string
  op string
  val int
  t string
  f string
}

func parseRule(r string) rule {
  var result rule
  work := strings.Split(r, ",")
  valT := strings.Split(work[0][2:], ":")

  result.part = string(work[0][0])
  result.op = string(work[0][1])
  result.val, _ = strconv.Atoi(valT[0])
  result.t = valT[1]
  result.f = strings.Join(work[1:], ",")
  
  return result
}

func flow(meep map[string]rule, name rule, vals map[string]int) int {
  var total int
  for _, v := range vals { total += v }
    if (name.op == ">" && vals[name.part] > name.val) || (name.op == "<" && vals[name.part] < name.val) {
      if strings.Contains(name.t, ":") {
        next := parseRule(name.t)
        return flow(meep, next, vals)
      }
      if name.t == "A" {
        return total
      }
      if name.t == "R" { return 0 }
      return flow(meep, meep[name.t], vals)
    } else {
      if strings.Contains(name.f, ":") {
        next := parseRule(name.f)
        return flow(meep, next, vals)
      }
      if name.f == "A" {
        return total
      }
      if name.f == "R" { return 0 }
      return flow(meep, meep[name.f], vals)
    }
}

func main() {
  var result int
  meep := make(map[string]rule)
  vals := make(map[string]int)

  file, _ := os.ReadFile(os.Args[1])
  input := strings.Split(string(file), "\n\n")
  workflows := input[0]
  values := strings.Split(input[1], "\n")
  values = values[:len(values) - 1]

  for _, wf := range strings.Split(workflows, "\n") {
    wfSplit := strings.Split(wf, "{")
    meep[wfSplit[0]] = parseRule(strings.Split(wfSplit[1], "}")[0])
  }
  for _, val := range values {
    val = strings.ReplaceAll(val, "{", "")
    val = strings.ReplaceAll(val, "}", "")
    singles := strings.Split(val, ",")
    
    for _, single := range singles {
      vals[string(single[0])], _ = strconv.Atoi(single[2:])
    }
    result += flow(meep, meep["in"], vals)
    fmt.Println(val, vals)
  }

  for k, v := range meep {
    fmt.Println(k, v)
  }


  fmt.Println(result)
}

