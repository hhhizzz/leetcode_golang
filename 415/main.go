package _415

func addStrings(num1 string, num2 string) string {
    if len(num1) < len(num2) {
        //make num1 always not shorter than num2
        num2, num1 = num1, num2
    }
    nums := make([]byte, len(num1)+1)
    current1 := len(num1) - 1
    current2 := len(num2) - 1

    var carry byte = 0
    for current2 >= 0 {
        n1 := num1[current1] - '0'
        n2 := num2[current2] - '0'
        n := n1 + n2 + carry
        carry = n / 10
        nums[current1+1] = n%10 + '0'
        current1--
        current2--
    }
    for current1 >= 0 {
        n1 := num1[current1] - '0'
        n := n1 + carry
        carry = n / 10
        nums[current1+1] = n%10 + '0'
        current1--
    }
    if carry != 0 {
        nums[0] = '1'
    } else {
        nums = nums[1:]
    }

    return string(nums)
}
