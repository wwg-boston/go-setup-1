package main

func main() {
  mp := messagePrinter{"go"}
  mp.printMessage()
}

type messagePrinter struct {
  message string
}

// seperate data from the methods
func (myFunc *messagePrinter) printMessage(){
  println(myFunc.message)
}
