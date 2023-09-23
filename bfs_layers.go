package main

import "fmt"

type MyGraph struct {
	vertex int
	graph  [][]int
}

func (g *MyGraph) graph_init(vertex int) {
	g.vertex = vertex
	g.graph = make([][]int, vertex)
}

func (g *MyGraph) add_edges(source int, sink int) {
	if len(g.graph[source]) == 0 {
		g.graph[source] = make([]int, 0)
		g.graph[source] = append(g.graph[source], sink)
	} else {
		g.graph[source] = append(g.graph[source], sink)
	}
}

func (g *MyGraph) bfs(vertex int) {
	visited := make([]bool, 0)
	for i := 0; i < g.vertex; i++ {
		visited = append(visited, false)
	}

	queue := make([]int, 0)

	is_empty := func() bool {
		if len(queue) == 0 {
			return true
		} else {
			return false
		}
	}

	enqueue := func(item int) {
		queue = append(queue, item)
	}

	dequeue := func() int {
		item := queue[0]
		queue = queue[1:]
		return item
	}

	fmt.Println("BFS traversal of the graph is:")

	visited[vertex] = true
	queue = append(queue, vertex)
	for is_empty() == false {
		q_len := len(queue)
		if q_len != 0 {
			fmt.Println(queue)
		} else {
			break
		}
		for q_len != 0 {
			item := dequeue()
			for _, val := range g.graph[item] {
				if visited[val] == false {
					enqueue(val)
					visited[val] = true
				}
			}
			q_len -= 1
		}
	}

}

func main() {
	graph := MyGraph{}

	graph.graph_init(4)
	graph.add_edges(2, 0)
	graph.add_edges(0, 2)
	graph.add_edges(2, 3)
	graph.add_edges(3, 3)
	graph.add_edges(0, 1)
	graph.add_edges(1, 2)

	fmt.Println("Performing level order traversal of the graph:")

	graph.bfs(2)
}
