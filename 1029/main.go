package _1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	A := []int{}
	B := []int{}
	for i, cost := range costs {
		if cost[0] < cost[1] {
			A = append(A, i)
		} else {
			B = append(B, i)
		}
	}
	if len(A) < len(B) {
		sort.Slice(B, func(i, j int) bool {
			return costs[B[i]][0]-costs[B[i]][1] > costs[B[j]][0]-costs[B[j]][1]
		})
		diff := (len(B) - len(A)) >> 1
		A = append(A, B[len(B)-diff:]...)
		B = B[:len(B)-diff]
	} else if len(B) < len(A) {
		sort.Slice(A, func(i, j int) bool {
			return costs[A[i]][0]-costs[A[i]][1] < costs[A[j]][0]-costs[A[j]][1]
		})
		diff := (len(A) - len(B)) >> 1
		B = append(B, A[len(A)-diff:]...)
		A = A[:len(A)-diff]
	}
	result := 0
	for _, a := range A {
		result += costs[a][0]
	}
	for _, b := range B {
		result += costs[b][1]
	}
	return result
}
