package main

import (
  "fmt"
)

func main(){
  var a int = 33
  var b *int = &a
  fmt.Println(a, *b)
  a = 12
  fmt.Println(a, *b)
  *b = 24
  fmt.Println(a, *b)

  
}
