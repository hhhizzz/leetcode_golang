package _229

func majorityElement(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    //缓存两个值`a,b`和他们的计数，初始可以任意，反正count都是0，不算
    a := 0
    b := 0
    countA := 0
    countB := 0

    for _, num := range nums {
        //如果新的的这个num是a，那么计数加一
        if num == a {
            countA++
            continue
        }
        if num == b {
            countB++
            continue
        }

        //如果缓存计数为0，表示这个数失效，将它换成新的数
        if countA == 0 {
            a = num
            countA = 1
            continue
        }
        if countB == 0 {
            b = num
            countB = 1
            continue
        }
        //新的这个数与缓存的数都不相同，同时缓存数都没失效，也就是还剩下计数，那么一起提交（也就是缓存的数计数减一，新的数跳过）
        countA--
        countB--
    }
    //进行二次检查
    countA = 0
    countB = 0
    for _, num := range nums {
        if num == a {
            countA++
        } else if num == b {
            //这里巧妙地绕过了两个缓存的数相同的情况
            countB++
        }
    }
    var result []int
    if countA > len(nums)/3 {
        result = append(result, a)
    }
    if countB > len(nums)/3 {
        result = append(result, b)
    }
    return result
}
