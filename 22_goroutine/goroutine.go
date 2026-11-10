package goroutine

// A goroutine is a function running at the same time as other goroutines.
// Start one with the go keyword. This channel only carries the result back.
// Channels themselves are the next exercise.
// https://go.dev/tour/concurrency/1

// Sum returns the sum of nums, computed in another goroutine.
func Sum(nums []int) int {
	ch := make(chan int)
	go func() {
		ch <- 0
	}()
	return <-ch
}
