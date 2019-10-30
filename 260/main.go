package _260

func singleNumber(nums []int) []int {
    sign := 0
    for _, num := range nums {
        sign ^= num
    }

    //最终这个sign表示目标a，b这两个数不同的比特位
    //我们可以取其中的一位出来 用以区分这两个数，（与上自己的负数表示取自己最低的那位1）
    sign &= -sign
    a, b := 0, 0
    for _, num := range nums {
        //这里假如a是sign那一位是1的，那么这个与操作可以把nums里面所有那位是1的筛选出来，进一步异或找出a来，b也同理
        if num&sign == sign {
            a ^= num
        } else {
            b ^= num
        }
    }
    return []int{a, b}
}
