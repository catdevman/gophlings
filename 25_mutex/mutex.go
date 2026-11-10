package mutexes

// A sync.Mutex lets one goroutine at a time touch a piece of data.
// Lock before you use the data, and Unlock when you are done.
// defer Unlock so a later return still releases the lock.
// https://go.dev/tour/concurrency/1

type Counter struct {
	n int
}

func (c *Counter) Add(d int) {
	c.n += d
}

func (c *Counter) Value() int {
	return c.n
}
