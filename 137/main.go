package _137

func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = (num ^ one) & ^two
		two = (num ^ two) & ^one
	}
	return one
}
