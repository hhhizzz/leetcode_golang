package _490

func hasPath(maze [][]int, start []int, destination []int) bool {
	var queue [][]int
	if maze[start[0]][start[1]] != 1 {
		queue = append(queue, start)
	}
	visited := make([][]bool, len(maze))
	for i := range maze {
		visited[i] = make([]bool, len(maze[i]))
	}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current[0]][current[1]] {
			continue
		}
		visited[current[0]][current[1]] = true
		if current[0] == destination[0] && current[1] == destination[1] {
			return true
		}
		if current[0] > 0 && maze[current[0]-1][current[1]] != 1 {
			next := []int{current[0], current[1]}
			for next[0] > 0 && maze[next[0]-1][next[1]] != 1 {
				next[0]--
			}
			queue = append(queue, next)
		}
		if current[0] < len(maze)-1 && maze[current[0]+1][current[1]] != 1 {
			next := []int{current[0], current[1]}
			for next[0] < len(maze)-1 && maze[next[0]+1][next[1]] != 1 {
				next[0]++
			}
			queue = append(queue, next)
		}
		if current[1] > 0 && maze[current[0]][current[1]-1] != 1 {
			next := []int{current[0], current[1]}
			for next[1] > 0 && maze[next[0]][next[1]-1] != 1 {
				next[1]--
			}
			queue = append(queue, next)
		}
		if current[1] < len(maze[current[0]])-1 && maze[current[0]][current[1]+1] != 1 {
			next := []int{current[0], current[1]}
			for next[1] < len(maze[current[0]])-1 && maze[next[0]][next[1]+1] != 1 {
				next[1]++
			}
			queue = append(queue, next)
		}
	}
	return false
}
