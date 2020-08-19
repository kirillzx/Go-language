package main

import(
  "fmt"
)

func main(){
  object := Value(1)
  var val Incrementer = &object
  for i:=0;i<5;i++{
      fmt.Printf("The value: %v\n", val.Increase())
  }

  var i interface{} = true
  switch i.(type) {
  case int:
    fmt.Printf("i is an integer")
  case string:
    fmt.Printf("i is a string")
  default:
    fmt.Printf("i is a boolean")
  }

}

//initialize interface
type Incrementer interface{
  Increase() int //write the methods here
}

//create object
type Value int

//method
func (digit *Value) Increase() int{
  *digit++
  return int(*digit)
}
