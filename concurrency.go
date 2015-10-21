// Concurrency
// Thread like constructs - go routines (required for synchronization)
// Channels - sharing data between threads
// Concurrency vs Parallelism
// Concurrency : multiple things to do at the same time
// Concurrency: only one thing is done at one time
// Parallel : each thread is assigned to a seperate thread
package main

import "runtime"

func main() {
  runtime.GOMAXPROCS(8)

  ch := make(chan string) // add a new channel
  doneCh := make(chan bool)

  go alphabet(ch) // adding 'go' makes it run it concurrently
  go printer(ch, doneCh)

  println("This is before") // will finish current function
  // the main application can exit before the previous function is run,
  // so we have to put the application to sleep to let the function run
  <-doneCh
}

func alphabet(ch chan string) {
  for i := byte('a'); i <= byte('z'); i++ {
    ch <- string(i)   // passing data to the channel
  }
  close(ch)  // makes it unavaible to recevive new messages
}

// need a function that will accept the message that
// we pass to the channel
func printer(ch chan string, doneCh chan bool) {
  for i:= range ch {  // loop will exit when the channel is finished sending messages, which will happen when it's closed
    println(i)
  }

  doneCh <-true
}

// Go routines: are very small
