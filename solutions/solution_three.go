package solutions

import (
	"strings"
	"sync"

	"github.com/iamgrewal7/masv/util"
)

type safeData struct {
	result []string
	m      sync.Mutex
}

func (d *safeData) update(i int, ans string) {
	d.m.Lock()
	d.result[i] = ans
	d.m.Unlock()
}

// SolutionThree uses goroutines to process each argument on different thread.
// It uses mutex to protect the data to which processes result is written
func SolutionThree(args []string) []string {
	d := safeData{result: make([]string, len(args))}
	var wg sync.WaitGroup
	for i, arg := range args {
		wg.Add(1)
		go solutionThreeHelper(i, arg, &d, &wg)
	}
	wg.Wait()
	return d.result
}

// solutionThreeHelper seperates digit and string.
// It uses rotate function to increase position of each
// character and writing the resulting string to the safeData struct
func solutionThreeHelper(index int, param string, d *safeData, wg *sync.WaitGroup) {
	digit := rune(param[0] - '0')
	var res strings.Builder
	for _, v := range param[1:] {
		res.WriteRune(util.Rotate(v, digit))
	}
	d.update(index, res.String())
	wg.Done()
}
