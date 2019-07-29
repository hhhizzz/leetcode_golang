package _22

func gen(n int, left int, right int, s string, list *[]string) {
    if left == n && right == n {
        *list = append(*list, s)
        return
    }
    if left < n {
        gen(n, left+1, right, s+"(", list)
    }
    if right < n && right < left {
        gen(n, left, right+1, s+")", list)
    }
}

func generateParenthesis(n int) []string {
    list := make([]string, 0)
    gen(n, 0, 0, "", &list)
    return list
}
