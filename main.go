package main

import "fmt"

type Graph struct {
	nodes []*GraphNode
}

type GraphNode struct {
	id    int
	x     int
	y     int
	edges map[int]int
}

func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}
func (g *Graph) AddNode() (id int) {
	id = len(g.nodes)
	g.nodes = append(g.nodes, &GraphNode{
		id:    id,
		edges: make(map[int]int),
	})
	return
}

func (g *Graph) AddEdge(n1, n2 int, w int) {
	g.nodes[n1].edges[n2] = w
}

func (g *Graph) Neighbors(id int) []int {
	neighbors := []int{}
	for _, node := range g.nodes {
		for edge := range node.edges {
			if node.id == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}
func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.nodes))
	for i := range g.nodes {
		nodes[i] = i
	}
	return nodes
}
func (g *Graph) Edges() [][3]int {
	edges := make([][3]int, 0, len(g.nodes))
	for i := 0; i < len(g.nodes); i++ {
		for k, v := range g.nodes[i].edges {
			edges = append(edges, [3]int{i, k, int(v)})
		}
	}
	return edges
}
func main() {
	a := New()
	a.AddNode()
	a.AddNode()
	a.AddNode()
	a.AddNode()
	a.AddNode()
	a.AddEdge(0, 1, 3)
	a.AddEdge(0, 2, 7)
	a.AddEdge(2, 4, 7)
	a.AddEdge(4, 1, 7)
	a.AddEdge(2, 4, 7)
	fmt.Println(a.Nodes())
	fmt.Println(a.Neighbors(0))
	fmt.Print(a.Neighbors(1))
	fmt.Println(a.Neighbors(2))
	fmt.Println(a.Neighbors(3))
	fmt.Println(a.Edges())
}
