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

func main() {
    println(first)
    println(second)
    println(third)
}
