package main

import (
  "fmt"
)

func main(){
  greet := "Hello"
  name := "bunny"
  greetings(&greet, &name)
  s := sum(2, 12, 3)
  fmt.Println("Your age is ", *s)
}

func greetings(greet, name *string){
  fmt.Println(*greet, *name)
}

func sum(val ...int) *int{ //returning a pointer
  iter := 0
  for _, v := range val{
    iter += v
  }
  return &iter //return adress of the memory cell becaues of pointer
}
