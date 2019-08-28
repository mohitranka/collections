package graph

import (
	"fmt"
)

type Graph struct {
	adj map[int][]int
	V   int
}

func NewGraph(V int) *Graph {
	return &Graph{adj: make(map[int][]int), V: V}
}

func (g *Graph) AddEdge(src int, dest int) {
	g.adj[src] = append(g.adj[src], dest)
}

func (g *Graph) Print() [][]int {
	s := make([][]int, 0)
	for _, v := range g.adj {
		s = append(s, v)
	}
	return s
}

func (g *Graph) DFS(start int) {
	visited := make([]bool, g.V)
	s := make([]int, 0)
	s = append(s, start)
	var item int
	for len(s) > 0 {
		item, s = s[len(s)-1], s[:len(s)-1]
		visited[item] = true
		fmt.Println(item)
		for _, i := range g.adj[item] {
			if !visited[i] {
				s = append(s, i)
			}
		}
	}
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.V)
	q := make([]int, 0)
	q = append(q, start)
	var item int
	for len(q) > 0 {
		item, q = q[0], q[1:]
		visited[item] = true
		fmt.Println(item)
		for _, i := range g.adj[item] {
			if !visited[i] {
				q = append(q, i)
			}
		}
	}
}

func (g *Graph) TopologicalSort() *[]int {
	stack := make([]int, 0)
	visited := make([]bool, g.V)
	for v := 0; v < g.V; v++ {
		// process children
		if !visited[v] {
			g.topologicalUtil(v, visited, &stack)
		}
	}
	return &stack
}

func (g *Graph) topologicalUtil(v int, visited []bool, stack *[]int) {
	visited[v] = true
	for c := range g.adj[v] {
		if !visited[c] {
			g.topologicalUtil(c, visited, stack)
		}
	}
	*stack = append(*stack, v)
}
