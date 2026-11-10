package loops

// Go has one loop keyword. The full form is for init; condition; post.
// Leave out the init and post and it works as a while loop.
// for { } runs until you break or return.
// range walks a slice, one element at a time.
// https://go.dev/tour/flowcontrol/1

// SumTo returns the sum of the integers from 1 through n.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// SumWhile returns the same sum using a for loop that has only a condition.
func SumWhile(n int) int {
	total := 0
	i := 1
	for i <= n {
		total += i
		i++
	}
	return total
}

// Find returns the index of target, or -1 when it is not there.
// Use a forever loop and leave it with return.
func Find(nums []int, target int) int {
	i := 0
	for {
		if i == len(nums) {
			return -1
		}
		if nums[i] == target {
			return i
		}
		i++
	}
}

// SumRange returns the sum of nums. Use range.
func SumRange(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
