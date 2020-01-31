package _294

func canWin(s string) bool {
    m := map[string]bool{}
    return helper(s, m)
}

//依据策梅洛定理进行记忆化搜索
func helper(s string, m map[string]bool) bool {
    if result, ok := m[s]; ok {
        return result
    }
    //遍历我的每一步可能的行为
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '+' && s[i+1] == '+' {
            temp := []byte(s)
            temp[i] = '-'
            temp[i+1] = '-'

            //检查此次行为后对方是否必胜
            result := helper(string(temp), m)
            m[string(temp)] = result
            //如果对方没有必胜说明我此次行为是必胜的
            if !result {
                m[s] = true
                return !result
            }
        }
    }
    //如果每个可能的行为后对方都是必胜的，那么我是必败的
    return false
}
