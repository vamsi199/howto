package main

import (
	"testing"
)

func TestSortFloat32(t *testing.T) {

	type testcase struct {
		name  string
		input []float32
		want  []float32
	}

	testcases := []testcase{}

	testcases = append(testcases, testcase{
		"testcase1",
		[]float32{1, 2, 2, 3, 1.2, 2.5, 2.9, 4.2, 0.8, 0.9, 0.9, 1, 1.01, 1.02},
		[]float32{0.8, 0.9, 0.9, 1, 1, 1.01, 1.02, 1.2, 2, 2, 2.5, 2.9, 3, 4.2},
	})

	for _, tc := range testcases {
		sortFloat32(tc.input)
		if !compareFloat32s(tc.want, tc.input, true) {
			t.Errorf("%v failed: wanted %v, but got %v", tc.name, tc.want, tc.input)
		}

	}

}

func TestDedupFloat32s(t *testing.T) {

	type testcase struct {
		name  string
		input []float32
		want  []float32
	}

	testcases := []testcase{}

	testcases = append(testcases, testcase{
		"testcase1",
		[]float32{0.8, 0.9, 0.9, 1, 1, 1.01, 1.02, 1.2, 2, 2, 2.5, 2.9, 3, 4.2},
		[]float32{0.8, 0.9, 1, 1.01, 1.02, 1.2, 2, 2.5, 2.9, 3, 4.2},
	})

	for _, tc := range testcases {
		got := dedupFloat32s(tc.input)
		sortFloat32(got)
		if !compareFloat32s(tc.want, got, true) {
			t.Errorf("%v failed: wanted %v, but got %v", tc.name, tc.want, got)
		}

	}

}
