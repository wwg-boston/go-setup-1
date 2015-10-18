package main

func main() {
  myVar := 1
  if myVar == 1 {  // parens not necessary here
    println("true")
  } else { // else has to be on the same line as the end closing brace
    println("false")
  }

  // can declare a variable right before comparing it
  if myVar2 := 2; myVar2 == 1 {
    println("true")
  } else {
    println("false")
  }

  myVar3 := 1
  switch{
  case myVar3 == 1:
    println("one")
  case myVar3 == 2:
    println("two")
  }
}
