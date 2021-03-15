// Steve Kalvi 3/11/21
// Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

const limit = .000000000000001

func Sqrt(x float64) float64 {
    z := float64(1)
    for i:=1; i < 10; i++ { 
	   z -= (z*z - x)/ (2*z)
	   fmt.Println("z is",z)
	} 
	return z
}

// Use Absolute value in case negative difference when approaching limit
// p or Prev value can't be set to z first time through.
func SqrtLimit(x float64) float64 {
    z := float64(1)
    for p:=float64(2); math.Abs(p-z) > limit;{
	   p = z
	   z -= (z*z - x)/ (2*z)
	   fmt.Println("z is: ",z, "p is: ",p)
	} 
	return z
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(SqrtLimit(2))
}
