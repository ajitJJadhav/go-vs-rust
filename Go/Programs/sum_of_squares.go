package main

import (
    "fmt"
    "math/rand"
	"time"
)

// function to add an array of numbers.
func sum(s []uint64) uint64{
	var sum uint64
	sum = 0
	for _, v := range s {
		sum += v*v
	}
	return sum
}

func main() {
	start := time.Now()

	var s [2000000]uint64
	n := 2000000
    for i := 0; i < n; i++ {
        s[i] = uint64(rand.Intn(100))
    }
	ans := sum(s[:])

	elapsed := time.Since(start)
    fmt.Println("Execution time: %s", elapsed)
	fmt.Println(ans)
}