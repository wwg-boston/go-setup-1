package main

import "fmt"

func main() {
  arr := [3]int{}
  arr[0] = 16
  arr[1] = 56
  arr[2] = 23

  fmt.Println(arr) // prints out a "pretty" version of the array

  // declaring an array in this way lets you not have to
  // worry about specifying the size ahead of time
  arr2 := [...]int{12, 45, 76, 75}
  fmt.Println(arr2)
  fmt.Println(len(arr2))

  // slices are more flexible because you can manipulate the size
  slice := arr2[:]  // create a slice of all the values in the array
  slice = append(slice, 12)
  fmt.Println(slice)

  slice2 := []float32{14., 15., 19.}
  fmt.Println(slice2)
  fmt.Println(len(slice2))

  // build a large initial slice
  slice3 := make([]float32, 100)
  slice3[0] = 1
  slice3[1] = 2
  slice3[2] = 3
  fmt.Println(slice3)

  map1 := make(map[int]string)
  fmt.Println(map1)
  map1[32] = "Hello"
  map1[21] = "Go"

  fmt.Println(map1)
  fmt.Println(map1[34])

}
