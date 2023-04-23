package graph

type Graph interface {
	InDegree(vertex int) int
	OutDegree(vertex int) int
	Transpose() Graph
	Distance(int) int
	BFSTree(int) *BfsTree
	BFS(int)
	AddEdge(int, int)
}

type color = int

const (
	white color = iota
	gray
	black
)

type node struct {
	edges []int
	color color
	distance int
	pred *node
}

type adjacencyList struct {
	nodes []*node
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

func newNode(maxEdges int) *node {
	return &node{
		edges: make([]int, maxEdges),
	}
}

func newAdjacencyList(vertices int) *adjacencyList {
	nodes := make([]*node, vertices)
	for i := range nodes {
		nodes[i] = newNode(vertices)
	}
	return &adjacencyList{
		nodes: nodes,
	}
}

func (a adjacencyList) OutDegree(vertex int) int {
	if vertex > len(a.nodes)-1 {
		return 0
	}
	out := 0
	for _, edge := range a.nodes[vertex].edges {
		out += edge
	}
	return out
}

func (a adjacencyList) InDegree(vertex int) int {
	degree := 0
	for i, node := range a.nodes {
		// skip vertex's edge-set
		if i == vertex {
			continue
		}
		// if vertex i has an edge to vertex add it
		degree += node.edges[vertex]
	}
	return degree
}

func (a adjacencyList) Transpose() Graph {
	newAdjList := newAdjacencyList(len(a.nodes))

	for i, node := range a.nodes {
		for j, edge := range node.edges {
			if edge == 1 {
				newAdjList.AddEdge(j, i)
			}
		}
	}
	return newAdjList
}

func (a *adjacencyList) AddEdge(fromVertex, toVertex int) {
	if fromVertex > len(a.nodes)-1 || toVertex > len(a.nodes)-1 {
		return
	}
	a.nodes[fromVertex].edges[toVertex] = 1
}

// set color of all nodes to white, called before BFS is executed on the adjacency list
func (a *adjacencyList) reset() {
	for _, node := range a.nodes {
		node.distance = 0
		node.color = white
		node.pred = nil
	}
}

type BfsNode struct {
	Vertex int
	Children []*BfsNode
}

func newBFSNode(vertex int) *BfsNode {
	return &BfsNode{
		Vertex: vertex,
		Children: make([]*BfsNode, 0),
	}
}

type BfsTree struct {
	Root *BfsNode
}

// given a node and a vertex, perform a BFS over the graph, and return the BFS-tree
func (g *adjacencyList) BFSTree(s int) *BfsTree {
	if s < 0 || s >= len(g.nodes) {
		return nil
	}
	// set clean slate
	g.reset()
	// create the tree + root node
	root := newBFSNode(s)
	// pass the rootNode + s to bfsWithNode
	g.bfsWithNode(g.nodes[s], root)
	// return tree
	return &BfsTree{
		root,
	}
}

func (g *adjacencyList) bfsWithNode(n *node, root *BfsNode) {
	type data struct {
		n *node
		bfs *BfsNode
	}
	// set distance on n
	if n.pred == nil {
		n.distance = 0
	} else {
		n.distance = n.pred.distance + 1
	}
	datas := make([]data, 0)
	// for each child of node
	for vertexIdx, isAdjacent := range n.edges {
		if isAdjacent == 1 {
			vertex := g.nodes[vertexIdx]
			// set color + pred
			vertex.color = gray
			vertex.pred = n
			// create bfs node, and add to data
			bfsNode := newBFSNode(vertexIdx)
			root.Children = append(root.Children, bfsNode)
			datas = append(datas, data{
				n: vertex,
				bfs: bfsNode,
			})
		}
	}
	// set n black 
	n.color = black
	// call bfsWithNode for all children
	for _, data := range datas {
		// don't traverse again if alr traversed
		if data.n.color != black {
			g.bfsWithNode(data.n, data.bfs)
		}
	}
}
// fifo queue implementation
type Queue[C any] interface {
	Push(C)
	Pop() (C, bool)
}

func NewQueue[C any]() Queue[C] {
	return &queueImpl[C]{}
}

type queueNode[C any] struct {
	val C
	next *queueNode[C]
}

type queueImpl[C any] struct {
	head *queueNode[C]
}

func newQueueNode[C any](val C) *queueNode[C] {
	return &queueNode[C]{
		val: val,
	}
}

func (q *queueImpl[C]) Push(val C)  {
	if q.head == nil {
		q.head = newQueueNode(val)
		return
	}
	// go to end of queue
	var node *queueNode[C]
	for node = q.head; node.next != nil; node = node.next {}
	node.next = newQueueNode(val)
}

func (q *queueImpl[C]) Pop() (val C, ok bool) {
	if q.head == nil {
		return
	}
	val = q.head.val
	ok = true
	q.head = q.head.next
	return
}

func (g *adjacencyList) BFS(s int) {
	if s < 0 || s >= len(g.nodes) {
		return
	}
	// reset the adjacency list
	g.reset()
	q := NewQueue[*node]()
	q.Push(g.nodes[s])
	for node, ok := q.Pop(); ok; {
		// for each node adjacent to node
		for vertexIdx, isAdjacent := range node.edges {
			if isAdjacent == 1 {
				vertex := g.nodes[vertexIdx]
				// if the node has alr been visited skip
				if vertex.color != white {
					continue
				}
				// set the color to gray
				vertex.color = gray
				vertex.distance = node.distance + 1
				vertex.pred = node
				// queue this node to be traversed as well
				q.Push(vertex)
			}
		}
		// set the color of the node to black
		node.color = black
	}
}

func (g *adjacencyList) Distance(i int) int {
	if i < 0 || i > len(g.nodes) {
		return -1
	}
	return g.nodes[i].distance
}
