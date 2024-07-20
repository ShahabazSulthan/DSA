package main

import "fmt"

// Graph structure using adjacency list
type Graph struct {
	vertices map[int][]int // A map to store adjacency list representation
}

// Initialize a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int), // Initialize the vertices map
	}
}

// Add an edge to the graph
func (g *Graph) AddEdge(v1, v2 int) {
	g.vertices[v1] = append(g.vertices[v1], v2) // Add v2 to the adjacency list of v1
	g.vertices[v2] = append(g.vertices[v2], v1) // Add v1 to the adjacency list of v2 (undirected graph) // Remove this line for a directed graph
}

func remove(slice []int, elem int) []int {
	for i, v := range slice {
		if v == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Remove an edge from the graph
func (g *Graph) RemoveEdge(v1, v2 int) {
	g.vertices[v1] = remove(g.vertices[v1], v2) // Remove v2 from the adjacency list of v1
	g.vertices[v2] = remove(g.vertices[v2], v1) // Remove v1 from the adjacency list of v2
}

// Print the adjacency list representation of the graph
func (g *Graph) Print() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%d -> %v \n", vertex, edges) // Print each vertex and its list of connected vertices
	}
}

// Depth-First Search (DFS) starting from a given vertex
func (g *Graph) DFS(start int, visited map[int]bool) {
	if visited[start] {
		return // If the vertex is already visited, return to avoid cycles
	}
	visited[start] = true // Mark the vertex as visited
	fmt.Println(start)    // Print the current vertex

	// Recursively visit all the connected vertices
	for _, v := range g.vertices[start] {
		g.DFS(v, visited)
	}
}

// Breadth-First Search (BFS) starting from a given vertex
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool) // Initialize the visited map
	queue := []int{start}         // Initialize the queue with the start vertex
	visited[start] = true         // Mark the start vertex as visited

	// Process the queue until it is empty
	for len(queue) > 0 {
		vertex := queue[0]  // Get the first vertex in the queue
		queue = queue[1:]   // Remove the first vertex from the queue
		fmt.Println(vertex) // Print the current vertex

		// Add all unvisited connected vertices to the queue
		for _, v := range g.vertices[vertex] {
			if !visited[v] {
				visited[v] = true        // Mark the vertex as visited
				queue = append(queue, v) // Add the vertex to the queue
			}
		}
	}
}

func main() {
	graph := NewGraph() // Create a new graph

	graph.AddEdge(1, 2) // Add edge between vertex 1 and vertex 2
	graph.AddEdge(1, 3) // Add edge between vertex 1 and vertex 3
	graph.AddEdge(2, 4) // Add edge between vertex 2 and vertex 4
	graph.AddEdge(4, 3) // Add edge between vertex 4 and vertex 3
	graph.Print()       // Print the adjacency list representation of the graph

	fmt.Println("DFS Traversal:")
	graph.DFS(1, make(map[int]bool)) // Perform DFS starting from vertex 1

	fmt.Println("BFS Traversal:")
	graph.BFS(1) // Perform BFS starting from vertex 1

	// Remove an edge and print the graph again
	graph.RemoveEdge(1, 3) // Remove the edge between vertex 1 and vertex 3
	fmt.Println("Graph after removing the edge between 1 and 3:")
	graph.Print()
}
