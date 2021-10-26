package solutions

import (
	"strings"

	"github.com/iamgrewal7/masv/util"
)

func SolutionOne(args []string) []string {
	result := make([]string, len(args))
	for i, arg := range args {
		result[i] = solutionOneHelper(arg)
	}
	return result
}

// solutionOneHelper seperates digit and string.
// It uses rotate function to increase position of each
// character and return the resulting string
func solutionOneHelper(param string) string {
	digit := rune(param[0] - '0')
	var res strings.Builder
	for _, v := range param[1:] {
		res.WriteRune(util.Rotate(v, digit))
	}
	return res.String()
}
