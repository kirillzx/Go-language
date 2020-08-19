package main

import (
  "fmt"
)

func main(){
  func (){
    fmt.Printf("Function inside another function")
  }()

  object1 := object{
    name: "John",
    id: 1,
  }
  fmt.Printf("\n")
  object1.method()

  greet := "Hello"
  name := "bunny"
  greetings(&greet, &name)
  s := sum(2, 12, 3)
  fmt.Println("Your age is ", *s)
  res, err := divide(2.1, 0.0)
  if err!=nil{
    fmt.Println(err)
    return
  }
  fmt.Println(res)

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

func divide(a, b float64) (float64, error){
  if b == 0.0{
    return 0.0, fmt.Errorf("Can not divide by zero")
  }
  return a / b, nil
}

type object struct{
  name string
  id int
}

//create a method
func (o object) method(){
  fmt.Printf("We have the user with the name %v and id %v\n", o.name, o.id)
}
