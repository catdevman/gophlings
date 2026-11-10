package channels

// A send on an unbuffered channel waits until another goroutine receives.
// A buffered channel holds a fixed number of values and the send waits
// only when the buffer is full.
// https://go.dev/tour/concurrency/1

// Ping sends msg on a channel and returns what was received.
func Ping(msg string) string {
	ch := make(chan string)
	go func() {
		ch <- msg
	}()
	return <-ch
}

// BufferedSum sends 1 and then 2 on one channel, then returns their sum.
// Both sends happen before either receive, so the channel needs room for 2.
func BufferedSum() int {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	return <-ch + <-ch
}
