package _402

func removeKdigits1(num string, k int) string {
    if len(num) == k {
        return "0"
    }
    nums := []byte(num)
    removed := true
    for removed && k > 0 {
        removed = false
        for i := 0; i < len(nums)-1; i++ {
            if nums[i] > nums[i+1] {
                removed = true
                nums = append(nums[:i], nums[i+1:]...)
                k--
                break
            }
        }
    }

    if k > 0 {
        nums = nums[:len(nums)-k]
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] != '0' {
            return string(nums[i:])
        }
    }

    return "0"
}

func removeKdigits(num string, k int) string {
    var s []byte
    for i := 0; i < len(num); i++ {
        if len(s) == 0 || num[i] >= s[len(s)-1] {
            if len(s) != 0 || num[i] != '0' {
                s = append(s, num[i])
            }
        } else {
            for j := len(s) - 1; j >= 0 && num[i] < s[j]; j-- {
                s = s[:j]
                k--
                if k == 0 {
                    break
                }
            }
            s = append(s, num[i])
        }
        if k == 0 {
            s = append(s, num[i+1:]...)
            break
        }
    }
    if k > 0 {
        s = s[:len(s)-k]
    }
    for i := 0; i < len(s); i++ {
        if s[i] != '0' {
            return string(s[i:])
        }
    }
    return "0"
}
