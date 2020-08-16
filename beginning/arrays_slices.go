package main

import (
  "fmt"
)

func main(){
  numbers := [...]int{1, 12, 34}
  cars := [2]string{"BMW", "Mercedes Benz"}
  var matrix [3][3]int
  matrix[0] = [3]int{1,2,3}
  matrix[1] = [3]int{4,5,6}
  matrix[2] = [3]int{7,8,9}
  fmt.Printf("%v\n%v", numbers, cars)
  fmt.Printf("\n%v", matrix)

  //using pointers for change elements in two arrays
  b := &numbers
  b[1] = 55
  fmt.Printf("\nAfter changing: %v\n%v", numbers, b)

  // slices
  slice := []float32{1.2, 3.4, 5.6, 2.3, 7.9, 12.4}
  slice1 := slice[:]
  slice2 := slice[2:]
  slice3 := slice[1:4]
  fmt.Printf("\nSlices: %v, %v, %v", slice1, slice2, slice3)

  slice = append(slice, 32.1)
  fmt.Printf("\nAdd element to slice: %v", slice)
}
