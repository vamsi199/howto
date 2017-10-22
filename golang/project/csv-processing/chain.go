package main

import "sort"

type chain struct {
	Start_Chain float32
	End_Chain   float32
}

type byChainStart []chain

func (a byChainStart) Len() int           { return len(a) }
func (a byChainStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byChainStart) Less(i, j int) bool { return a[i].Start_Chain < a[j].Start_Chain }
func sortChains(c []chain)                { sort.Sort(byChainStart(c)) }

func compareChains(s1, s2 []chain, sorted bool) bool {

	if !sorted {
		sortChains(s1)
		sortChains(s2)
	}

	if len(s1) != len(s2) {
		return false
	}

	for i, _ := range s1 {
		if s1[i].Start_Chain != s2[i].Start_Chain ||
			s1[i].End_Chain != s2[i].End_Chain {
			return false
		}
	}

	return true
}

func prepareChains(input []float32) []chain {

	if len(input) < 2 {
		return []chain{}
	}

	output := make([]chain, 0, len(input))
	for i, _ := range input {
		if i == 0 {
			continue
		}
		output = append(output, chain{input[i-1], input[i]})
	}
	return output
}
