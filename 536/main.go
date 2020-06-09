package _536

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func array2tree(s []byte, pos *int) *TreeNode {
	num := 0
	minus := false
	for ; *pos < len(s) && s[*pos] != '(' && s[*pos] != ')'; (*pos)++ {
		if s[*pos] == '-' {
			minus = true
		} else {
			num *= 10
			num += int(s[*pos] - '0')
		}
	}
	if minus {
		num = -num
	}
	result := &TreeNode{Val: num}
	if s[*pos] == ')' {
		(*pos)++
		return result
	}
	if s[*pos] == '(' {
		(*pos)++
		result.Left = array2tree(s, pos)
	}
	if s[*pos] == ')' {
		(*pos)++
		return result
	}
	if s[*pos] == '(' {
		(*pos)++
		result.Right = array2tree(s, pos)
	}
	if s[*pos] == ')' {
		(*pos)++
		return result
	}
	return result
}

func str2tree(s string) *TreeNode {
	if s == "" {
		return nil
	}
	array := []byte(s + ")")
	pos := 0
	return array2tree(array, &pos)
}
