package _885

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	var result [][]int
	number := 1
	target := R * C
	//metric := make([][]int,R)
	// for i:=0;i<R;i++{
	//     metric[i] = make([]int, C)
	// }
	visited := map[[2]int]bool{[2]int{r0, c0}: true}
	checkValid := func(i, j int) bool {
		_, ok := visited[[2]int{i, j}]
		return !ok
	}
	// metric[r0][c0] = number
	number++
	result = append(result, []int{r0, c0})

	for number <= target {
		//right
		c0++
		for {
			visited[[2]int{r0, c0}] = true
			if r0 < R && c0 < C && r0 >= 0 && c0 >= 0 {
				// metric[r0][c0] = number
				result = append(result, []int{r0, c0})
				number++
			}
			if checkValid(r0+1, c0) {
				break
			}
			c0++
		}
		// fmt.Println(metric)
		//down
		r0++
		for {
			visited[[2]int{r0, c0}] = true
			if r0 < R && c0 < C && r0 >= 0 && c0 >= 0 {
				// metric[r0][c0] = number
				result = append(result, []int{r0, c0})
				number++
			}
			if checkValid(r0, c0-1) {
				break
			}
			r0++
		}
		// fmt.Println(metric)
		//left
		c0--
		for {
			visited[[2]int{r0, c0}] = true
			if r0 < R && c0 < C && r0 >= 0 && c0 >= 0 {
				// metric[r0][c0] = number
				result = append(result, []int{r0, c0})
				number++
			}
			if checkValid(r0-1, c0) {
				break
			}
			c0--
		}
		// fmt.Println(metric)
		//up
		r0--
		for {
			visited[[2]int{r0, c0}] = true
			if r0 < R && c0 < C && r0 >= 0 && c0 >= 0 {
				// metric[r0][c0] = number
				result = append(result, []int{r0, c0})
				number++
			}
			if checkValid(r0, c0+1) {
				break
			}
			r0--
		}
		// fmt.Println(metric)
	}
	return result
}
