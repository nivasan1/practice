package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gr "leetcode.com/leetcode/graph"
)

func TestGraph(t *testing.T) {
	// create a new graph
	graph := gr.NewGraph(2, true)
	// add edges to 1
	graph.AddEdge(0, 1)
	assert.Equal(t, graph.OutDegree(0), 1)
	// transpose
	transpose := graph.Transpose()
	assert.Equal(t, transpose.OutDegree(1), 1)
}

func TestBFS(t *testing.T) {
	// create test graph
	graph := gr.NewGraph(3, true)
	// add edges from 0 -> 1 -> 2
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	// get the bfs tree
	tree := graph.BFSTree(0)
	root := tree.Root.Vertex
	assert.Equal(t, tree.Root.Vertex, root)
	assert.Equal(t, len(tree.Root.Children), 1)
	next := tree.Root.Children[0]
	assert.Equal(t, next.Vertex, 1)
	next = next.Children[0]
	assert.Equal(t, next.Vertex, 2)
}

func TestQueue(t *testing.T) {
	q := gr.NewQueue[int]()
	q.Push(1)
	q.Push(2)
	val, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, val, 1)
	val, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, val, 2)
	_, ok = q.Pop()
	assert.False(t, ok)
}
