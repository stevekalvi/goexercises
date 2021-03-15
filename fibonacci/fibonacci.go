// Steve Kalvi 3/15/21
// Fibonacci Sequence Closure starting at zero
package main

import "fmt"

func fibonacci () func() int {
  a, b := 0, 1
  return func() int {
     a, b = a + b, a  // a = a + b, b = old a is precedence
     //fmt.Println(a,b)
     return b
  }
}

func main() {
     f := fibonacci()
     for i := 1; i < 10; i++ {
      	fmt.Print(f(), ",")
     }	
     fmt.Println(f())
}