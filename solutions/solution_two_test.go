package solutions

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestSolutionTwo(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{
			[]string{"0apple", "1lzru", "6gump", "8hello", "5zyo", "9a", "1abcdefgh"},
			[]string{"apple", "masv", "masv", "pmttw", "edt", "j", "bcdefghi"},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Input #: %d", i+1)
		t.Run(testname, func(t *testing.T) {
			result := SolutionTwo(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Got = %s, Expected = %s", result, tt.expected)
			}
		})
	}
}

func BenchmarkSolutionTwo(b *testing.B) {
	for k := 0; k <= 10; k++ {
		n := int(math.Pow(2, float64(k)))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				args := make([]string, n)
				for j := range args {
					args[j] = "1abcdefgh"
				}
				SolutionTwo(args)
			}
		})
	}
}
