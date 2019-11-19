package _119

func getRow(rowIndex int) []int {
	array := make([][]int,2)
	array[0] = make([]int,rowIndex+1)
	array[1] = make([]int,rowIndex+1)
	for i:=0;i<=rowIndex;i++{
			current := i % 2
			last := (i-1) % 2
			for j:=0;j<i+1;j++{
					if j==0 || j==i{
							array[current][j] = 1
					}else{
							array[current][j] = array[last][j-1]+array[last][j]
					}
			}
	}
	return array[rowIndex % 2]
}