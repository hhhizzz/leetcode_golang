package _60

// 对于 n位长的数来说，有n!种排列，其中每个第一位有(n-1)!中排列，例如4位数，总共有24中排列，其中 1xxx有6种 2xxx有6种等等
// 利用这个特性可以逐步求出第k个排列的每一位
func getPermutation(n int, k int) string {
	var res []byte
	var seq []byte
	factorial := 1
	for i := 1; i <= n; i++ {
		seq = append(seq, byte(i)+'0')
		factorial *= i
	}
	k -= 1
	for i := 0; i < n; i++ {
		factorial /= n - i
		times := k / factorial
		res = append(res, seq[times])
		seq = append(seq[:times], seq[times+1:]...)
		k -= factorial * times
	}
	return string(res)
}
