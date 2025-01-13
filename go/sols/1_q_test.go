package sols

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {

	type testcase struct {
		n       int
		prereqs [][]int
		output  []int
	}

	testcases := []testcase{
		{
			n: 2,
			prereqs: [][]int{
				[]int{1, 0},
			},
			output: []int{0, 1},
		},
		{
			n: 4,
			prereqs: [][]int{
				[]int{1, 0},
				[]int{2, 0},
				[]int{3, 1},
				[]int{3, 2},
			},
			output: []int{0, 1, 2, 3},
		},
		{
			n:       1,
			prereqs: [][]int{},
			output:  []int{0},
		},
	}
	for _, testcase := range testcases {
		output, err := Q1(testcase.n, testcase.prereqs)
		assert.Equal(t, err, nil, "error should be nil")
		//todo: implement any valid case (not needed for practise)
		// assert.Equal(t, testcase.output, output, "should be equal")
		fmt.Println(output)
	}

}
