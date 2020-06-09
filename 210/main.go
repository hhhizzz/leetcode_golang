package _210

func findOrder(numCourses int, prerequisites [][]int) []int {
	input := make([]int, numCourses)
	output := make([]map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		output[i] = map[int]bool{}
	}
	for _, prerequisite := range prerequisites {
		input[prerequisite[0]] += 1
		output[prerequisite[1]][prerequisite[0]] = true
	}
	var stack []int
	for i := 0; i < numCourses; i++ {
		if input[i] == 0 {
			stack = append(stack, i)
		}
	}
	var result []int
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current)
		for k := range output[current] {
			input[k] -= 1
			if input[k] == 0 {
				stack = append(stack, k)
			}
		}
	}
	if len(result) < numCourses {
		result = []int{}
	}
	return result

}
