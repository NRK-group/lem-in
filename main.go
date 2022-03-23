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
	// Key      int
	Name     string
	X        int
	Y        int
	Status   string
	Adjacent map[string]*Vertex
	// Adjacent []*Vertex
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
		// g.Vertices = append(g.Vertices,
		// 	&Vertex{
		// 		// Key:  k,
		// 		Name:     name,
		// 		X:        x,
		// 		Y:        y,
		// 		Adjacent: make(map[string]*Vertex),
		// 	})
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
	if fromVertex == nil && toVertex == nil {
		// return
		fmt.Println("HELLO")
	} else {
		fromVertex.Adjacent[toVertex.Name] = toVertex
	}

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
	// grph := NewGraph()
	for i := 0; i < len(coordinates); i++ {
		room := coordinates[i][0]
		x, _ := strconv.Atoi(coordinates[i][1])
		y, _ := strconv.Atoi(coordinates[i][2])
		g.AddVertex(room, x, y)
	}
	// fmt.Println(ants, start, end, coordinates, links)
	for i := range links {
		from := links[i][0]
		to := links[i][1]
		g.AddEdge(from, to)
		g.AddEdge(to, from)
	}
	// for _, i := range coordinates {
	// 	fmt.Println(g.Vertices[i[0]])
	// 	// fmt.Print(grph.Vertices[1])
	// }
	// fmt.Println(g.Vertices[0].Adjacent[0].Name)
}
func main() {
	path := "example/"
	s := []string{"example00.txt", "example01.txt", "example02.txt", "example03.txt", "example04.txt", "example05.txt", "example06.txt", "example07.txt", "example08.txt", "example09.txt", "example10.txt", "badexample00.txt", "badexample01.txt"}
	g := NewGraph()
	g.Populate(path + s[5])
	fmt.Print(g.Vertices["start"].Adjacent["A0"].Adjacent["A1"].Adjacent["A2"].Adjacent["end"])
}
