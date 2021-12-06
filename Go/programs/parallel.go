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

func Convolve_parallel(u, v []uint32) uint32 {

    var ans uint32
    ans = 0
    n := len(u)
    
    size := 20000
    var wg sync.WaitGroup
    for i, j := 0, size; i < n; i, j = j, j+size {
        if j > n {
            j = n
        }
        
        wg.Add(1)
        go func(i, j int) {
            for k := i; k < j; k++ {
                u[k] = u[k]*u[k]; 
            }
            wg.Done()
        }(i, j)
    }
    wg.Wait()
    return ans
}

func Convolve(u, v []uint32) uint32 {

    var ans uint32
	ans = 0
    for i := 0; i < len(u); i++ {
        ans += u[i]*u[i]
    }
    return ans
}


func main() {
	
	
	var s [2000000]uint32
	n := 2000000
	for i := 0; i < n; i++ {
	 s[i] = uint32(rand.Intn(100))
	}

	start := time.Now()
	ans := Convolve_parallel(s[:], s[:])
	elapsed := time.Since(start)
	fmt.Println("Execution time parallel: %s", elapsed)
	fmt.Println("Total: ", ans)

	start1 := time.Now()
	ans1 := Convolve(s[:], s[:])
	elapsed1 := time.Since(start1)
	fmt.Println("Execution time: %s", elapsed1)
	fmt.Println("Total: ", ans1)
}
