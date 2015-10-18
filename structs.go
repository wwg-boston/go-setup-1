// struct -> arbitrary data fields

package main

import "fmt"

func main() {
  myVar := myStruct{}
  myVar.myField = "Hello"

  println(myVar.myField)

  // another way of doing this
  myVar2 := myStruct{"Go"}
  println(myVar2.myField)

  // creating values on the heap, vs the stack
  myVar3 := myStruct{}
  myVar3.myField = "!"
  fmt.Println(myVar3)

  // create a reference type instead
  myVar4 := new (myStruct{})
  myVar4.myField = "?"
  fmt.Println(myVar4)
}

type myStruct struct {
  myField string
}
