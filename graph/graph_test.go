package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	g "leetcode.com/leetcode/graph"
)

func TestGraph(t *testing.T) {
	// create a new graph
	graph := g.NewGraph(2, true)
	// add edges to 1
	graph.AddEdge(0, 1)
	assert.Equal(t, graph.OutDegree(0), 1)
	// transpose
	transpose := graph.Transpose()
	assert.Equal(t, transpose.OutDegree(1), 1)
}
