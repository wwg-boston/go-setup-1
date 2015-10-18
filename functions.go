package main

func main() {
  name := "James"
  printGreeting("one", "two", "three")
  println(name);
}

// the value passed in is a copy, not the original
func printGreeting2(name string) {
  println("Hi " + name)
}

func printGreeting3(name *string) {
  println("Hi " + *name)
  *name = "Jessica"
}

// variadic functions
// passed in as a slice of strings, passed into
// the messages parameter
func printGreeting(messages ...string) {
  for _, message := range messages {
    println(message)
  }
}
