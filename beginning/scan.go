package main

import (
  "fmt"
  "math"
)

func main(){
  db := database{}
  var a float64 = 13.14
  db.id = createID(a)
  db.addElement()
  fmt.Println(db)
}

type database struct{
  name string
  id int
  age int
}

func (db *database) addElement(){
  var name string
  var age int
  fmt.Scanln(&name) //read by line from the console
  fmt.Scanln(&age)
  db.name = name
  db.age = age
}

func createID (a float64) int{
  var value int = int(math.Round(math.Sin(a)))
  return value
}
