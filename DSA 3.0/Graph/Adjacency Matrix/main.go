package main

import "fmt"

// Graph structure using adjacency matrix
type Graph struct {
	vertices int
	matrix   [][]int
}

// Initialize a new graph with a given number of vertices
func NewGraph(vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &Graph{
		vertices: vertices,
		matrix:   matrix,
	}
}

// Add an edge to the graph with a given weight
func (g *Graph) AddEdges(v1, v2, weight int) {
	g.matrix[v1][v2] = weight
	g.matrix[v2][v1] = weight
}

// Print the adjacency matrix of the graph
func (g *Graph) Print() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%d: %v\n", i, g.matrix[i])
	}
}

func main() {
	graph := NewGraph(5)

	graph.AddEdges(0, 1, 1)
	graph.AddEdges(0, 2, 1)
	graph.AddEdges(1, 2, 1)
	graph.AddEdges(2, 3, 1)
	graph.AddEdges(3, 4, 1)
	graph.AddEdges(4, 1, 1)
	graph.Print()
}
