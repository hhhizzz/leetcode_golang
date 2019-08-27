package _679

import "math"

//暴力搜索所有数字和符号组合即可，方法是选两个数字进行加减乘除变成3个，然后继续变成2个，变成1个看这个数是不是24，注意float的比较不能直接==
func judgePoint24(nums []int) bool {
    var floatNums []float64
    for i := 0; i < len(nums); i++ {
        floatNums = append(floatNums, float64(nums[i]))
    }
    return solve(&floatNums)
}

func solve(nums *[]float64) bool {
    if len(*nums) <= 0 {
        return false
    } else if len(*nums) == 1 {
        if math.Abs((*nums)[0]-24) <= 1e-6 {
            return true
        }
    } else {
        for i := 0; i < len(*nums); i++ {
            for j := 0; j < len(*nums); j++ {
                if i != j {
                    var floatNums []float64
                    for k := 0; k < len(*nums); k++ {
                        if k != i && k != j {
                            floatNums = append(floatNums, (*nums)[k])
                        }
                    }
                    for operator := 0; operator < 4; operator++ {
                        if operator == 0 {
                            floatNums = append(floatNums, (*nums)[i]+(*nums)[j])
                        } else if operator == 1 {
                            floatNums = append(floatNums, (*nums)[i]-(*nums)[j])
                        } else if operator == 2 {
                            floatNums = append(floatNums, (*nums)[i]*(*nums)[j])
                        } else if operator == 3 {
                            floatNums = append(floatNums, (*nums)[i]/(*nums)[j])
                        }
                        if !solve(&floatNums) {
                            floatNums = floatNums[:len(floatNums)-1]
                        } else {
                            return true
                        }
                    }
                }
            }
        }
    }
    return false
}
