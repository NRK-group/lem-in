package main

import (
	"fmt"
	"lemin/function"
	"strconv"
)

type Graph struct {
	Vertices []*Vertex
}
type Vertex struct {
	// Key      int
	Name     string
	X        int
	Y        int
	Status   string
	Adjacent []*Vertex
}

func (g *Graph) AddVertex(name string, x, y int) {
	g.Vertices = append(g.Vertices,
		&Vertex{
			// Key:  k,
			Name: name,
			X:    x,
			Y:    y,
		})
}
func NewGraph() *Graph {
	return &Graph{
		Vertices: []*Vertex{},
	}
}
func Contains(s []*Vertex, room string) bool {
	for _, v := range s {
		if room == v.Name {
			return true
		}
	}
	return false
}
func (g *Graph) Populate(file string) {
	a, f := function.ValidateFile(file)
	if !(a) {
		fmt.Println(f[0]) // this will print an error message
		return
	}
	_, _, _, coordinates, _ := function.Clean(f)
	// grph := NewGraph()
	for i := 0; i < len(coordinates); i++ {
		room := coordinates[i][0]
		x, _ := strconv.Atoi(coordinates[i][1])
		y, _ := strconv.Atoi(coordinates[i][2])
		g.AddVertex(room, x, y)
	}
	// fmt.Println(ants, start, end, coordinates, links)
	for i := range coordinates {
		fmt.Println(g.Vertices[i])
		// fmt.Print(grph.Vertices[1])
	}
}
func main() {
	path := "example/"
	s := []string{"example00.txt", "example01.txt", "example02.txt", "example03.txt", "example04.txt", "example05.txt", "example06.txt", "example07.txt", "example08.txt", "example09.txt", "example10.txt", "badexample00.txt", "badexample01.txt"}
	g := NewGraph()
	g.Populate(path + s[10])
}
