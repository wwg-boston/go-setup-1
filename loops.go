package main

func main() {
  for i:=0; i < 5; i++ {
    println(i);
  }

  i := 2
  for {
    i++
    println(i)
    if(i < 5) {
      break
    }
  }

  // iterating through slices
  s := []string{"one", "two", "three"}

  for idx, v:= range s {
    println(idx, v)
  }

  // iterating through maps
  map1 := make(map[int]string)
  map1[32] = "Hello"
  map1[21] = "Go"

  for k,v := range map1 {
    println(k, v)
  }
}
