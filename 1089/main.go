package _1089

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 && arr[i] == 0 {
			temp := arr[i+1]
			arr[i+1] = 0
			for j := i + 2; j < len(arr); j++ {
				arr[j], temp = temp, arr[j]
			}
			i++
		}
	}
}
