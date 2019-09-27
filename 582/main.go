package _582

type node struct {
    id       int
    children []*node
}

func killNodes(current *node, result *[]int) {
    *result = append(*result, current.id)
    for _, child := range current.children {
        killNodes(child, result)
    }
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
    var result []int
    killNodes(killNode, &result)
    return result
}
