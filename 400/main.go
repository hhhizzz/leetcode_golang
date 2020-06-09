package _400

func findNthDigit(n int) int {
	var number []int
	sum := 0
	current := 9
	currentNumber := 1
	for sum <= 1<<31 {
		number = append(number, current*currentNumber)
		sum += current
		currentNumber += 1
		current *= 10
	}
	//fmt.Println(number)
	i := 0
	for ; n > number[i]; i++ {
		n -= number[i]
	}
	//fmt.Println(n)

	digits := 1
	for j := 0; j < i; j++ {
		digits *= 10
	}
	//fmt.Println(digits)

	pos := (n-1)/(i+1) + digits

	//fmt.Println(pos)

	pos2 := (n - 1) % (i + 1)

	//fmt.Println(pos2)

	stringPos := strconv.Itoa(pos)

	return int(stringPos[pos2] - '0')
}
