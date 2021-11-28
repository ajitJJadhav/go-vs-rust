package main
import (
	"fmt"
	"time"
	"math/rand"
	"runtime"
	"sync"
)

func init() {
    numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

// Convolve computes w = u * v, where w[k] = Σ u[i]*v[j], i + j = k.
// Precondition: len(u) > 0, len(v) > 0.
// func Convolve(u, v []uint64) []uint64 {
//     n := len(u) + len(v) - 1
//     w := make([]uint64, n)
//     for k := 0; k < n; k++ {
//         w[k] = mul(u, v, k)
//     }
//     return w
// }

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
func mul(u, v []uint64, k int) uint64 {
    var res uint64
    n := min(k+1, len(u))
    j := min(k, len(v)-1)
    for i := k - j; i < n; i, j = i+1, j-1 {
        res += u[i] * v[j]
    }
    return res
}
func Convolve(u, v []uint64) []uint64 {
    n := len(u) + len(v) - 1
    w := make([]uint64, n)

    // Divide w into work units that take ~100μs-1ms to compute.
    size := max(1, 1000000/n)

    var wg sync.WaitGroup
    for i, j := 0, size; i < n; i, j = j, j+size {
        if j > n {
            j = n
        }
        // The goroutines share memory, but only for reading.
        wg.Add(1)
        go func(i, j int) {
            for k := i; k < j; k++ {
                w[k] = mul(u, v, k)
            }
            wg.Done()
        }(i, j)
    }
    wg.Wait()
    return w
}

func main() {
	start := time.Now()
	
	var s [20000]uint64
	n := 20000
    for i := 0; i < n; i++ {
        s[i] = uint64(rand.Intn(100))
    }
	ans := Convolve(s[:], s[:])

	
	elapsed := time.Since(start)
	fmt.Println("Execution time: %s", elapsed)
	fmt.Println("Total: ", ans[0])
}