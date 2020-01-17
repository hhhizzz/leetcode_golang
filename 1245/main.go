package _1245

func bfs(start int, edgeMap map[int][]int) (int, int) {
    dis := map[int]int{start: 0}
    queue := []int{start}

    maxDis := 0
    maxNode := start

    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]

        for _, next := range edgeMap[current] {
            if _, ok := dis[next]; !ok {
                queue = append(queue, next)
                dis[next] = dis[current] + 1
                if dis[next] > maxDis {
                    maxDis = dis[next]
                    maxNode = next
                }
            }
        }
    }
    return maxDis, maxNode
}

func treeDiameter(edges [][]int) int {
    edgeMap := map[int][]int{}
    for _, edge := range edges {
        edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
        edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
    }
    _, node := bfs(0, edgeMap)

    dis, _ := bfs(node, edgeMap)

    return dis
}
