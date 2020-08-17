package main

import (
  "fmt"
)

type People struct{
  name string
  age int
  friends []string
}

type Student struct{
  People
  bandScore int
}

func main(){
  //map
  db := map[string]int{
    "Alex": 12,
    "Max": 34,
    "John": 18,
    "David": 23,
  }
  fmt.Println(db)
  //add new element
  db["Rob"] = 51
  fmt.Printf("\nAdded new name: %v", db)
  //delete element in map
  delete(db, "Max")
  fmt.Printf("\nDeleted name Max: %v\n", db)
  //accessory check
  _, a := db["David"]
  fmt.Println(a)

  //structure
  people := People{
    name: "Lisa",
    age: 23,
    friends: []string{"Tom", "Wendy"},
  }
  fmt.Printf("\nStructure: %v", people)

  //structure in one line
  people1 := struct{name string}{name: "John"}
  //copy structure
  people2 := people1
  people2.name = "Sam"
  fmt.Printf("\nInitial structure: %v\nAfter coping: %v", people1, people2)
  //with rememeber adress
  people3 := &people1
  people3.name = "Sam"
  fmt.Printf("\nInitial structure: %v\nAfter coping: %v\n", people1, people3)
  //duplicate structure in antoher structure
  student := Student{}
  student.name = "Alex"
  student.age = 20
  student.friends = []string{"Felix", "Oliver"}
  student.bandScore = 8
  fmt.Println(student)

  //another way
  student2 := Student{
    People: People{name: "Rob", age: 24, friends: []string{"Wendy", "John"}},
    bandScore: 7,
  }
  fmt.Printf("\n%v", student2)
}
