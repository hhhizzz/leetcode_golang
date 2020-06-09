package _869

func toByte(num int) []byte {
	var result []byte
	for num >= 10 {
		mod := byte(num % 10)
		result = append(result, mod)
		num /= 10
	}
	result = append(result, byte(num))
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	qSort(result)
	return result
}

func qSort(nums []byte) {
	if len(nums) <= 1 {
		return
	}
	pivot := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			pivot++
			nums[i], nums[pivot] = nums[pivot], nums[i]
		}
	}
	nums[pivot], nums[0] = nums[0], nums[pivot]
	qSort(nums[:pivot])
	qSort(nums[pivot+1:])
}

//其实不用那么复杂 就是突然想不用标准库纯手写一下
func reorderedPowerOf2(N int) bool {
	power := 1
	var powerList [][]byte
	for power <= 1e9 {
		powerList = append(powerList, toByte(power))
		power *= 2
	}
	bytes := toByte(N)
	for _, powerBytes := range powerList {
		if len(powerBytes) != len(bytes) {
			continue
		} else {
			same := true
			for i := 0; i < len(bytes); i++ {
				if powerBytes[i] != bytes[i] {
					same = false
					break
				}
			}
			if same {
				return true
			}
		}
	}
	return false
}
