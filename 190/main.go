package _190

func reverseBits(num uint32) uint32 {
	bytes := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = num % 2
		num /= 2
	}

	var result uint32 = 0
	var bit uint32 = 1
	for i := 31; i >= 0; i-- {
		result += bit * bytes[i]
		bit *= 2
	}
	return result
}
