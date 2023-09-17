package main

import "fmt"

type Graph struct {
	graph  map[int][]int
	vertex int
}

func (graph *Graph) init_graph(vertex int) {
	graph.vertex = vertex
	graph.graph = make(map[int][]int)
}

func (graph *Graph) add_edges(source int, sink int) {
	_, ok := graph.graph[source]
	if ok == true {
		graph.graph[source] = append(graph.graph[source], sink)
	} else {
		graph.graph[source] = make([]int, 0)
		graph.graph[source] = append(graph.graph[source], sink)
	}
}

func (graph *Graph) bfs_traversal(start_vertex int) {
	queue := make([]int, 0)
	visited := make([]bool, 0)

	for i := 0; i < graph.vertex; i++ {
		visited = append(visited, false)
	}

	enqueue := func(element int) {
		if queue == nil {
			queue = make([]int, 0)
		} else {
			queue = append(queue, element)
		}
	}

	queue_empty := func() bool {
		if queue == nil {
			return false
		}
		if len(queue) > 0 {
			return true
		} else {
			return false
		}
	}

	dequeue := func() int {
		if len(queue) > 0 {
			element := queue[0]
			queue = queue[1:]
			return element
		} else {
			return -1
		}
	}

	fmt.Println("BFS_traversal")
	queue = append(queue, start_vertex)
	visited[start_vertex] = true
	for queue_empty() {
		myelement := dequeue()
		fmt.Println(myelement)
		for _, value := range graph.graph[myelement] {
			if visited[value] == false {
				enqueue(value)
				visited[value] = true
			}
		}
		if len(queue) == 0 {
			break
		}
	}
}

func main() {
	graph := Graph{}
	graph.init_graph(4)
	graph.add_edges(0, 2)
	graph.add_edges(2, 0)
	graph.add_edges(2, 3)
	graph.add_edges(3, 3)
	graph.add_edges(0, 1)
	graph.add_edges(1, 2)
	fmt.Println(graph.graph)
	fmt.Println("The bfs traversal of the graph:")
	graph.bfs_traversal(2)
}
