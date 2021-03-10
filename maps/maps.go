package main

import (
     "golang.org/x/tour/wc"
     "strings"
)

func WordCount(s string) map[string]int {
     m := make(map[string]int)
     z := strings.Split(s, " ")
     for _, elem := range z {
        m[elem] += 1
     }
     return m
}

func main() {
    wc.Test(WordCount)
}
