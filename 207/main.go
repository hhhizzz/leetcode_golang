package _207

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]bool, numCourses)
	input := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		g[i] = make([]bool, numCourses)
	}
	for i := 0; i < len(prerequisites); i++ {
		g[prerequisites[i][0]][prerequisites[i][1]] = true
		input[prerequisites[i][1]] += 1
	}

	stack := []int{}
	finished := 0
	for i := 0; i < numCourses; i++ {
		if input[i] == 0 {
			stack = append(stack, i)
			finished += 1
		}
	}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := 0; i < len(g[current]); i++ {
			if g[current][i] {
				g[current][i] = false
				input[i] -= 1
				if input[i] == 0 {
					stack = append(stack, i)
					finished += 1
				}
			}
		}
	}
	return finished == numCourses
}
