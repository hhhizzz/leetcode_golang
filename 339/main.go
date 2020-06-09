package _339

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

func dfs(integer *NestedInteger, level int) int {
	if integer.IsInteger() {
		return level * integer.GetInteger()
	} else {
		result := 0
		for _, next := range integer.GetList() {
			result += dfs(next, level+1)
		}
		return result
	}
}

func depthSum(nestedList []*NestedInteger) int {
	result := 0
	for _, nestedInteger := range nestedList {
		result += dfs(nestedInteger, 1)
	}
	return result
}
