// go doesn't have constructors
// instead, object functions serve that purpose
// constructors can accept parameters for setting initial
// values

package main

import "fmt"

func main() {
  // items := &myStruct{}
  items := newMyStruct()
  // maps have to be initialized before they're used
  items.myMap["apples"] = 32
  fmt.Println(items)
}

type myStruct struct {
  myMap map[string]int
}

// create a properly initialized object: constructor function
// return it as a pointer
// custom name of naming is to start with 'new'
func newMyStruct() *myStruct {
  result := myStruct{}
  result.myMap = map[string]int{}
  return &result
}
