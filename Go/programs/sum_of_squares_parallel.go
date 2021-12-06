// package main

// import (
//     "fmt"
//     "math/rand"
// 	"time"
// )

// // function to add an array of numbers.
// func sum(s []uint64, c chan uint64) {
// 	var sum uint64
// 	sum = 0
// 	for _, v := range s {
// 		sum += v*v
// 	}
// 	// writes the sum to the go routines.
// 	c <- sum // send sum to c
// }

// func main() {

// 	start := time.Now()

// 	var ans uint64
// 	var s [2000000]uint64
// 	n := 2000000
// 	chunk_size := 25000
// 	buffer_size := 1000
//     for i := 0; i < n; i++ {
//         s[i] = uint64(rand.Intn(100))
//     }

// 	c := make(chan uint64, buffer_size)
// 	for i := 0; i < n; i += chunk_size {
// 		if i + chunk_size < n {
// 			go sum(s[i:i+chunk_size], c)
// 		} else {
// 			go sum(s[i:], c)
// 		}
		
//     }

// 	ans = 0
// 	for i := 0; i < n; i+=4 {
// 		x := <-c
// 		ans += x
// 	}
	
// 	elapsed := time.Since(start)
//     fmt.Println("Execution time: %s", elapsed)
// 	fmt.Println(ans)
// }





package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)
// return channel for input numbers
func getInputChan() <-chan int {
	// make return channel
	input := make(chan int, 100)

	// sample numbers
	var numbers [2000000]int
	n := 2000000
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(10)
	}

	// run goroutine
	go func() {
		for num := range numbers {
			input <- num
		}
		// close channel once all numbers are sent to channel
		close(input)
	}()

	return input
}

// returns a channel which returns square of numbers
func getSquareChan(input <-chan int) <-chan int {
	// make return channel
	output := make(chan int, 100)

	// run goroutine
	go func() {
		// push squares until input channel closes
		for num := range input {
			output <- num * num
		}

		// close output channel once for loop finishesh
		close(output)
	}()

	return output
}

// returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this is veriadic function
func merge(outputsChan ...<-chan int) <-chan int {
	// create a WaitGroup
	var wg sync.WaitGroup
	
	// make return channel
	merged := make(chan int, 100)
	
	// increase counter to number of channels `len(outputsChan)`
	// as we will spawn number of goroutines equal to number of channels received to merge
	wg.Add(len(outputsChan))
	
	// function that accept a channel (which sends square numbers)
	// to push numbers to merged channel
	output := func(sc <-chan int) {
		// run until channel (square numbers sender) closes
		for sqr := range sc {
			merged <- sqr
		}
		// once channel (square numbers sender) closes,
		// call `Done` on `WaitGroup` to decrement counter
		wg.Done()
	}
	
	// run above `output` function as groutines, `n` number of times
	// where n is equal to number of channels received as argument the function
	// here we are using `for range` loop on `outputsChan` hence no need to manually tell `n`
	for _, optChan := range outputsChan {
		go output(optChan)
	}
	
	// run goroutine to close merged channel once done
	go func() {
		// wait until WaitGroup finishesh
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	start := time.Now()
	// step 1: get input numbers channel
	// by calling `getInputChan` function, it runs a goroutine which sends number to returned channel
	chanInputNums := getInputChan()
	
	// step 2: `fan-out` square operations to multiple goroutines
	// this can be done by calling `getSquareChan` function multiple times where individual function call returns a channel which sends square of numbers provided by `chanInputNums` channel
	// `getSquareChan` function runs goroutines internally where squaring operation is ran concurrently
	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)
	
	// step 3: fan-in (combine) `chanOptSqr1` and `chanOptSqr2` output to merged channel
	// this is achieved by calling `merge` function which takes multiple channels as arguments
	// and using `WaitGroup` and multiple goroutines to receive square number, we can send square numbers
	// to `merged` channel and close it
	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)
	
	// step 4: let's sum all the squares from 0 to 9 which should be about `285`
	// this is done by using `for range` loop on `chanMergedSqr`
	sqrSum := 0
	
	// run until `chanMergedSqr` or merged channel closes
	// that happens in `merge` function when all goroutines pushing to merged channel finishes
	// check line no. 86 and 87
	for num := range chanMergedSqr {
		sqrSum += num
	}
	
	elapsed := time.Since(start)
	fmt.Println("Execution time: %s", elapsed)
	// step 5: print sum when above `for loop` is done executing which is after `chanMergedSqr` channel closes
	fmt.Println("Total: ", sqrSum)
}