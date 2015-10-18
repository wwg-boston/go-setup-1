package main

func main() {
  add := 1 + 2
  println(add)

  subtract := 10 - 5
  println(subtract)

  remainder := 5 % 2
  println(remainder)

  inc := 1
  inc++
  println(inc)
  inc++
  println(inc)
  inc *= 5 
  println(inc)
  // println(inc++) -> this is not allowed in Go, in order
  // to prevent confusion about the expected result
}
