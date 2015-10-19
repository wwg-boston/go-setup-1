// in traditional OOP, classes inherit from a parent class
// this can be confusing as you try to figure out
// what function is actually being used; for this reason, in Go,
// inheritance not supported
// composition is used instead.
// source structs can include the object that they want to
// use the methods from, and this makes it more clear what function is

package main

func main() {
  mp := advancedMessagePrinter{}
  mp.message = "go"
  mp.printMessage()
}

type messagePrinter struct {
  message string
}

type advancedMessagePrinter struct {
  messagePrinter
}

// seperate data from the methods
func (myFunc *messagePrinter) printMessage(){
  println(myFunc.message)
}
