package _364

type NestedInteger struct{}

func (NestedInteger) IsInteger() bool {
	panic("implement me")
}

func (NestedInteger) GetInteger() int {
	panic("implement me")
}

func (NestedInteger) SetInteger(value int) {
	panic("implement me")
}

func (NestedInteger) Add(elem NestedInteger) {
	panic("implement me")
}

func (NestedInteger) GetList() []*NestedInteger {
	panic("implement me")
}

func dfs(integer *NestedInteger, split *[]int, level int) {
	if level >= len(*split) {
		*split = append(*split, 0)
	}
	if integer.IsInteger() {
		(*split)[level] += integer.GetInteger()
	} else {
		for _, next := range integer.GetList() {
			dfs(next, split, level+1)
		}
	}
}

func depthSumInverse(nestedList []*NestedInteger) int {
	var split []int
	for _, integer := range nestedList {
		dfs(integer, &split, 0)
	}
	result := 0
	for i := 0; i < len(split); i++ {
		result += split[i] * (len(split) - i)
	}
	return result
}
