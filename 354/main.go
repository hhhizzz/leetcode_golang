package _354

type envelope struct {
    width    int
    height   int
    inDegree int
    out      []*envelope
    length   int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func (en *envelope) Bigger(another *envelope) bool {
    return en.width > another.width && en.height > another.height
}

//DAG中找最长路径算法
func maxEnvelopes1(envelopes [][]int) int {
    var array []envelope
    for i := 0; i < len(envelopes); i++ {
        array = append(array, envelope{envelopes[i][0], envelopes[i][1], 0, []*envelope{}, 0})
    }
    //构建DAG
    for i := 0; i < len(array); i++ {
        for j := i + 1; j < len(array); j++ {
            if array[i].Bigger(&array[j]) {
                array[i].out = append(array[i].out, &array[j])
                array[j].inDegree += 1
            } else if array[j].Bigger(&array[i]) {
                array[j].out = append(array[j].out, &array[i])
                array[i].inDegree += 1
            }
        }
    }
    var stack []*envelope
    for i := 0; i < len(array); i++ {
        if array[i].inDegree == 0 {
            array[i].length = 1
            stack = append(stack, &array[i])
        }
    }
    //找到最长路径
    result := 0
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = max(result, current.length)
        for _, out := range current.out {
            out.length = max(out.length, current.length+1)
            out.inDegree--
            if out.inDegree == 0 {
                stack = append(stack, out)
            }
        }
    }
    return result
}

func less(a, b []int) bool {
    if a[0] < b[0] {
        return true
    } else if a[0] == b[0] {
        if a[1] > b[1] {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

func qsort(array [][]int) {
    if len(array) <= 1 {
        return
    }
    pivot := 0
    for i := 0; i < len(array); i++ {
        if less(array[i], array[0]) {
            pivot++
            array[i], array[pivot] = array[pivot], array[i]
        }
    }
    array[0], array[pivot] = array[pivot], array[0]
    qsort(array[:pivot])
    qsort(array[pivot+1:])
}

//方法2，转化为LIS问题
func maxEnvelopes(envelopes [][]int) int {
    qsort(envelopes)
    var LIS []int
    for i := 0; i < len(envelopes); i++ {
        number := envelopes[i][1]
        if i == 0 || number > LIS[len(LIS)-1] {
            LIS = append(LIS, number)
        } else {
            j := len(LIS) - 1
            for ; j >= 0; j-- {
                if LIS[j] < number {
                    break
                }
            }
            LIS[j+1] = number
        }
    }
    return len(LIS)
}
