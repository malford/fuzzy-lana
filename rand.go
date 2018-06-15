package main

import "time"
import "fmt"
import "math/rand"

func main() {
	fmt.Println(time.Now())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println()
}
