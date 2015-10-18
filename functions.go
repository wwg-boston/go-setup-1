package main

func main() {
  name := "James"
  printGreeting("one", "two", "three")
  println(name);
  total, result := add2(1, 2, 3)
  println("There are: ", total, " items, and the sum is: ", result)
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

// Return values in Go - returning multiple returns
// this function returns 2 integer terms
func add(items ...int) (int, int) {
  result := 0
  for _, item := range items{
      result += item
  }
  return len(items), result
}

// name return variables
func add2(items ...int) (total int, result int) {
  for _, item := range items{
      result += item
  }
  total = len(items)
  return
}
