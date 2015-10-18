package main

// iota gives us an auto-incrementing number that determines
// the value of the integers (basic type of enumeration)
// first = 1
// second = 1 * 2 ^ 10 = 1024
// third = second * 2 ^ 10 = 1024 * 1024 = 1048576
const (
    first =  1 << (10 * iota)
    second
    third
)

var (
  canChange = 1
)

func main() {
    println(first)
    println(second)
    println(third)
}

// if you try to set a const to something else in the program,
// it will result in an error
// func main() {
//   //first += 1 --> will result in an error
//   canChange += 1
//   println(canChange)
// }
