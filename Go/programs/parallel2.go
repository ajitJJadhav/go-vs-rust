package main
import (
	"fmt"
	"time"
	"math/rand"
	"runtime"
	"sync"
)

func init() {
    // numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(8) // Try to use all available CPUs.
}

// Convolve computes w = u * v, where w[k] = Σ u[i]*v[j], i + j = k.
// Precondition: len(u) > 0, len(v) > 0.
func Convolve1(u []uint64) uint64 {
    var ans uint64
	ans = 0
    for i := 0; i < len(u); i++ {
        ans += u[i]*u[i]
    }
    return ans
}

func min(x, y int) int {
	if x < y {
	  return x
	}
	return y
   }
   
func max(x, y int) int {
if x > y {
	return x
}
return y
}

// mul returns Σ u[i]*v[j], i + j = k.
// func mul(u, v []uint64, k int) uint64 {
//     var res uint64
//     for i := 0; i < len(u); {
//         res += u[i] * v[j]
//     }
//     return res
// }

func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
   }

func Convolve(u []uint64) uint64 {
	var ans uint64
	ans = 0
    n := len(u)

    // Divide w into work units that take ~100μs-1ms to compute.
    // size := max(1, 1000000/n)
	size := 200000

    var wg sync.WaitGroup
    for i, j := 0, size; i < n; i, j = j, j+size {
        if j > n {
            j = n
        }
        // The goroutines share memory, but only for reading.
        wg.Add(1)
        go func(i, j int) {
            for k := i; k < j; k++ {
                ans += u[k]*u[k]
            }
            wg.Done()
        }(i, j)
    }
    wg.Wait()
    return ans
}

func main() {
	
	
	var s [2000000]uint64
	n := 2000000
    for i := 0; i < n; i++ {
        s[i] = uint64(rand.Intn(100))
    }

	start := time.Now()
	ans := Convolve(s[:])
	elapsed := time.Since(start)
	fmt.Println("Execution time: %s", elapsed)
	fmt.Println("Total: ", ans)
}