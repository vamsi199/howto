package main

import "fmt"

func commonOmxn(aa []int, bb []int) []int {
	bigOCnt := 0
	o := make([]int, 0, len(aa)) //
	for _, a := range aa {
		for _, b := range bb {
			bigOCnt++
			if a == b {
				o = append(o, a)
			}
		}
	}
	fmt.Printf("bigO count for commonOmaxmn = %v\n", bigOCnt)
	return o
}

//3*10*10+4=3*m*n+4
3*1000*1000+4
3000004
3000000
3*m*n


3*m+4
3m


func commonOmaxmn(aa []int, bb []int) []int {
	bigOCnt := 0
	out := make([]int, 0, len(aa)) //
	ai, bi := 0, 0
	for ai < len(aa) && bi < len(bb) {
		bigOCnt++
		if aa[ai] == bb[bi] {
			out = append(out, aa[ai])
			ai++
			bi++
		} else if aa[ai] < bb[bi] {
			ai++
		} else if bb[bi] < aa[ai] {
			bi++
		}
	}

	fmt.Printf("bigO count for commonOmaxmn = %v\n", bigOCnt)
	return out
}
