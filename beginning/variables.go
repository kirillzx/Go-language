package main

import (
  "fmt"
  "math"
)

var (
  name string = "Kirill"
  age int = 19
)

const (
  isHungry = 1 << iota
  isTired
  isSleep
)

func main(){
  const count float32 = 12.2
  n := math.Sqrt(12)
  var id float32 = count + 12
  fmt.Printf("Name: %v\nAge: %v\nID: %v", name, age, float64(id)+n)
  fmt.Printf("\nIs hungry? %v", isHungry&isSleep == isHungry)
}
