package main

import (
  "fmt"
)

func main(){
  db := map[string]int{
    "Alex": 12,
    "Max": 34,
    "John": 18,
    "David": 23,
  }

  if val, c := db["Max"]; c{
    fmt.Printf("%v", val)
  }

  var number int = 10
  count := 10

  if number < count && number > 0{
    fmt.Printf("\nLess than count")
  } else if number > count{
    fmt.Printf("\nGreater than count\n")
  } else {
    fmt.Printf("\nEqual\n")
  }

  switch i:=5-2;i {
  case 1, 3:
    fmt.Printf("\nThe number 1 or 3\n")
    fallthrough // go to the next condition
  case 2: fmt.Printf("\nThe number 2\n")
  default: fmt.Printf("\nAnother number\n")
  }

  // fmt.Println(number>=count, number!=count)
}
