package main

import "fmt"

type Graph struct {
	vertex int
	graph  [][]int
}

func (g *Graph) graph_init(vertex int) {
	g.vertex = vertex
	g.graph = make([][]int, vertex)
}

func (g *Graph) add_edges(source int, sink int) {
	if len(g.graph[source]) == 0 {
		g.graph[source] = make([]int, 0)
		g.graph[source] = append(g.graph[source], sink)
	} else {
		g.graph[source] = append(g.graph[source], sink)
	}
}

func (g *Graph) dfs_util(visited []bool, index int) {
	visited[index] = true
	fmt.Println(index)
	for _, val := range g.graph[index] {
		if visited[val] == false {
			g.dfs_util(visited, val)
		}
	}

}

func dfs(g Graph) {
	visited := make([]bool, 0)
	for i := 0; i < g.vertex; i++ {
		visited = append(visited, false)
	}
	for i := 0; i < g.vertex; i++ {
		if visited[i] == false {
			g.dfs_util(visited, i)
		}
	}
}

func main() {
	fmt.Println("Constructing the graph.")
	graph := Graph{}

	fmt.Println("Constructing the graph:")
	graph.graph_init(4)
	graph.add_edges(0, 2)
	graph.add_edges(2, 0)
	graph.add_edges(2, 3)
	graph.add_edges(3, 3)
	graph.add_edges(0, 1)
	graph.add_edges(1, 2)

	dfs(graph)
}
