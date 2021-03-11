// Steve Kalvi 3/11/21
// Exercise: Slices

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  //Slice a length of dy
  y := make([][]uint8, dy)
  for y_index, _ := range y {
      x := make([]uint8, dx)
      for x_index, _ := range x {
          x[x_index] = uint8(x_index * y_index) 
      }
      y[y_index] = x  // assign x to each y slice
  }
  return y
}

func main() {
  pic.Show(Pic)
}