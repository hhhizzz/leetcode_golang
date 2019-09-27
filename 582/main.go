package _582

type node struct {
    id       int
    children []*node
}

func killProcess(pid []int, ppid []int, kill int) []int {
    root := node{id: 0, children: []*node{}}
    nodes := map[int]*node{0: &root}
    for _, p := range pid {
        newNode := node{id: p, children: []*node{}}
        nodes[p] = &newNode
    }
    for i, pp := range ppid {
        p := pid[i]
        parentNode := nodes[pp]
        childrenNode := nodes[p]
        parentNode.children = append(parentNode.children, childrenNode)
    }
    killNode := nodes[kill]
    stack := []*node{killNode}
    var result []int
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.id)
        for _, child := range current.children {
            stack = append(stack, child)
        }
    }
    return result
}
