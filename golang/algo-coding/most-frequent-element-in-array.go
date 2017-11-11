package main

func mostFrequent(arr []int) int { // assuming no tie
	m := map[int]int{}
	var maxCnt int
	var freq int
	for _, a := range arr {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			freq = a
		}
	}

	return freq
}
