// Concurrency
// Thread like constructs - go routines (required for synchronization)
// Channels - sharing data between threads
// Concurrency vs Parallelism
// Concurrency : multiple things to do at the same time
// Concurrency: only one thing is done at one time
// Parallel : each thread is assigned to a seperate thread
package main

import "time"
import "runtime"

func main() {
  runtime.GOMAXPROCS(8)
  go alphabeth() // adding 'go' makes it run it concurrently

  println("This is before") // will finish current function
  // the main application can exit before the previous function is run,
  // so we have to put the application to sleep to let the function run
  time.Sleep(100 * time.Millisecond)
}

func alphabeth() {
  for i := byte('a'); i <= byte('z'); i++ {
    go println(string(i)) //optimize application internally
  }
}

// Go routines: are very small
