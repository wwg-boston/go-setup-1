package main

// iota gives us an auto-incrementing number that determines
// the value of the integers 
const (
    first = iota
    second
    third
)

func main() {
    println(first)
    println(second)
    println(third)
}
