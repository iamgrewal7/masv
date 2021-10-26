package solutions

import (
	"strings"
	"sync"

	"github.com/iamgrewal7/masv/util"
)

type data struct {
	arg   string
	index int
}

// SolutionTwo uses goroutines to process each argument on different thread.
// It makes use of buffered channel for sending and receiving processed result.
func SolutionTwo(args []string) []string {
	var wg sync.WaitGroup
	c := make(chan data, 20)

	for i, arg := range args {
		wg.Add(1)
		go solutionTwoHelper(arg, i, c, &wg)
	}

	// Wait for all goroutines to finish and close channel
	go func() {
		wg.Wait()
		close(c)
	}()

	result := make([]string, len(args))
	for v := range c {
		result[v.index] = v.arg
	}

	return result
}

// solutionTwoHelper seperates digit and string.
// It uses rotate function to increase position of each
// character and return the resulting struct to a channel
func solutionTwoHelper(param string, index int, c chan data, wg *sync.WaitGroup) {
	digit := rune(param[0] - '0')
	var res strings.Builder
	for _, v := range param[1:] {
		res.WriteRune(util.Rotate(v, digit))
	}
	c <- data{res.String(), index}
	wg.Done()
}
