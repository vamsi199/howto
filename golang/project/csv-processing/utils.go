package main

import (
	"sort"
)

type ByVal []float32

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i] < a[j] }
func sortFloat32(s []float32)      { sort.Sort(ByVal(s)) }

func compareFloat32s(s1, s2 []float32, sorted bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	if sorted {
		for i, _ := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}

	}

	//TODO: need to handle the case of unsorted sets

	return true
}

func dedupFloat32s(input []float32) []float32 { //TODO: need to improve the dedup algorithm performance
	temp := map[float32]bool{}
	for _, v := range input {
		temp[v] = true
	}

	output := make([]float32, 0, len(temp))
	for v, _ := range temp {
		output = append(output, v)
	}

	return output
}
