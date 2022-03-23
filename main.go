package main

import (
	"fmt"
	"lemin/function"
	"strconv"
)

type Graph struct {
	Vertices map[string]*Vertex
}
type Vertex struct {
	Name     string
	X        int
	Y        int
	Status   string
	Adjacent map[string]*Vertex
}

func Contains(s map[string]*Vertex, room string) bool { // check that their is no duplicate room name
	for _, v := range s {
		if room == v.Name {
			return true
		}
	}
	return false
}
func (g *Graph) AddVertex(name string, x, y int) { // Add vertex
	if !(Contains(g.Vertices, name)) {
		g.Vertices[name] = &Vertex{
			Name:     name,
			X:        x,
			Y:        y,
			Adjacent: make(map[string]*Vertex),
		}

	}
}
func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	fromVertex.Adjacent[toVertex.Name] = toVertex
}
func (g *Graph) GetVertex(room string) *Vertex {
	for i, v := range g.Vertices {
		if v.Name == room {
			return g.Vertices[i]
		}
	}
	return nil
}
func NewGraph() *Graph { //initialize a graph
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) Populate(file string) {
	a, f := function.ValidateFile(file)
	if !(a) {
		fmt.Println(f[0]) // this will print an error message
		return
	}
	_, _, _, coordinates, links := function.Clean(f)
	for i := 0; i < len(coordinates); i++ { //add every vertex
		room := coordinates[i][0]
		x, _ := strconv.Atoi(coordinates[i][1])
		y, _ := strconv.Atoi(coordinates[i][2])
		g.AddVertex(room, x, y)
	}
	// fmt.Println(ants, start, end, coordinates, links)
	for i := range links { // add every connection
		from := links[i][0]
		to := links[i][1]
		g.AddEdge(from, to)
		g.AddEdge(to, from)
	}
	// if len(g.Vertices[start[0]].Adjacent) == 0 { // check if their is a path from start room
	// 	fmt.Print("ERROR: invalid data format, graph incomplete")
	// 	return
	// }
	// if len(g.Vertices[end[0]].Adjacent) == 0 { // check if their is a path to the end room
	// 	fmt.Print("ERROR: invalid data format, no path found to the end room")
	// 	return
	// }
	// for _, i := range coordinates { // print every room
	// 	fmt.Println(g.Vertices[i[0]])
	// }
}
func Dfs(g *Graph, start string) {
	fmt.Println(g.Vertices[start])
	for _, n := range g.Vertices[start].Adjacent {
		g.Vertices[start].Status = "visited"
		if n.Status != "visited" {
			Dfs(g, n.Name)
			return
		}
	}
}
func main() {
	path := "example/"
	s := []string{"example00.txt", "example01.txt", "example02.txt", "example03.txt", "example04.txt", "example05.txt", "example06.txt", "example07.txt", "example08.txt", "example09.txt", "example10.txt", "badexample00.txt", "badexample01.txt", "badexample03.txt"}
	g := NewGraph()
	// for i := range s {
	g.Populate(path + s[5])
	// }
	g.Vertices["end"].Adjacent = map[string]*Vertex{}
	g.Vertices["start"].Status = "visited"
	Dfs(g, "G0")
}
