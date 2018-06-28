package main

import (
	"fmt"
)

func merge(l, m int, in [][]int) (int, int, [][]int) {
	var out [][]int
	var limit = l
	var max = m
	for _, i := range in {
		if i[0] > limit && max-limit == 0 {
			limit = i[0]
		} else if i[0] < limit {
			limit = i[0]
		}
		if i[0] >= limit && i[0] <= max && i[0]+i[1] > max {
			max = i[0] + i[1]
		} else if i[0] > max {
			if max == 0 {
				max = i[0] + i[1]
			} else {
				out = append(out, i)
			}

		}
	}
	if len(out) > 0 && len(in) != len(out) {
		return merge(limit, max, out)
	}
	return limit, max, out
}

func main() {

	var in = [][]int{[]int{1, 3}, []int{2, 6}, []int{4, 1}, []int{10, 1}}
	var in2 = [][]int{[]int{1, 3}, []int{2, 6}, []int{4, 1}, []int{10, 1}, []int{0, 15}}
	{
		limit, max, result1 := merge(0, 0, in)
		result1 = append(result1, []int{limit, max - limit})

		fmt.Printf("result = %v\n", result1)
	}
	{

		limit, max, result2 := merge(0, 0, in2)
		result2 = append(result2, []int{limit, max - limit})

		fmt.Printf("result2 = %v\n", result2)
	}
}
