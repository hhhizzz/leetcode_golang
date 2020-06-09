package _204

//正常作法
func countPrimes1(n int) int {
	if n <= 2 {
		return 0
	}
	primes := []int{2}
	for i := 3; i < n; i++ {
		flag := true
		for j := 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, i)
		}
	}
	return len(primes)
}

//筛选法
func countPrimes(n int) int {
	count := 0
	number := make([]bool, n)
	for i := 2; i < n; i++ {
		if number[i] == false {
			count++
		}
		for j := i * i; j < n; j += i {
			number[j] = true
		}

	}
	return count
}
