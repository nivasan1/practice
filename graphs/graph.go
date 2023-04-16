package graph

type Graph interface {
	InDegree(vertex int) int
	OutDegree(vertex int) int
	Transpose() Graph
	AddEdge(int, int)
}

type adjacencyList struct {
	edges [][]int
}

type adjacencyMatrix struct {
	matrix [][]int
}

func NewGraph(vertices int, list bool) Graph {
	if list {
		return newAdjacencyList(vertices)
	}
	return nil
}

func newAdjacencyList(vertices int) *adjacencyList {
	edges := make([][]int, vertices)
	for i := range edges {
		edges[i] = make([]int, vertices)
	}
	return &adjacencyList{
		edges: edges,
	}
}

func (a adjacencyList) OutDegree(vertex int) int {
	if vertex > len(a.edges)-1 {
		return 0
	}
	return len(a.edges[vertex])
}

func (a adjacencyList) InDegree(vertex int) int {
	degree := 0
	for i, edges := range a.edges {
		// skip vertex's edge-set
		if i == vertex {
			continue
		}
		// if vertex i has an edge to vertex add it
		degree += edges[vertex]
	}
	return degree
}

func (a adjacencyList) Transpose() Graph {
	return nil
}

func (a *adjacencyList) AddEdge(fromVertex, toVertex int) {
	if fromVertex > len(a.edges)-1 || toVertex > len(a.edges)-1 {
		return
	}
}
