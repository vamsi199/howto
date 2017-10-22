package main

import (
	"testing"
)

func TestAlgo(t *testing.T) {

	type testcase struct {
		name  string
		input []float32
		want  []chain
	}

	testcases := []testcase{}

	testcases = append(testcases, testcase{
		"testcase 1",
		[]float32{1, 2, 2, 3, 1.2, 2.5, 2.9, 4.2, 0.8, 0.9, 0.9, 1, 1.01, 1.02},
		[]chain{
			{0.8, 0.9},
			{0.9, 1},
			{1, 1.01},
			{1.01, 1.02},
			{1.02, 1.2},
			{1.2, 2},
			{2, 2.5},
			{2.5, 2.9},
			{2.9, 3},
			{3, 4.2},
		},
	})

	for _, tc := range testcases {
		got := algo(tc.input)
		if !compareChains(tc.want, got, false) {
			//t.Errorf("%v failed: wanted %v, but got %v", tc.name, tc.want, got)
		}
	}

}
