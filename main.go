package main

import (
	"fmt"
	"lemin/function"
	"strconv"
)

//Graph stuct is the wraper of the whole graph.
// The graph are constists of map of *Room the key of each *Room is the room name
type Graph struct {
	Vertices map[string]*Room
}

// Room struct will be the container of the Room
//  The Room is consist of Name of the room, the coordinates(X and Y), the Status(unvisited or visited),
// and the Link(the links in the graph)
type Room struct {
	Name string
	X    int
	Y    int
	Link map[string]*Room
}

// Contains checks if their is no duplicate room name
func Contains(s map[string]*Room, room string) bool {
	for _, v := range s {
		if room == v.Name {
			return true
		}
	}
	return false
}

// AddRoom add the room "name", "x", "y", and initialize the the Link
func (g *Graph) AddRoom(name string, x, y int) {
	if !(Contains(g.Vertices, name)) {
		g.Vertices[name] = &Room{
			Name: name,
			X:    x,
			Y:    y,
			Link: make(map[string]*Room),
		}
	}
}

//AddEdge adds the links to  the vertices
func (g *Graph) AddEdge(from, to string) {
	fromRoom := g.GetRoom(from)
	toRoom := g.GetRoom(to)
	fromRoom.Link[toRoom.Name] = toRoom
}

//GetRoom returns the room vertices if the room exists and return nil if not
func (g *Graph) GetRoom(room string) *Room {
	for i, v := range g.Vertices {
		if v.Name == room {
			return g.Vertices[i]
		}
	}
	return nil
}

// NewGraph initialize a new graph
func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Room{},
	}
}

// Populate function populates the graph using the data from the file.
func (g *Graph) Populate(coordinates, links [][]string) {

	//add every Room
	for i := 0; i < len(coordinates); i++ {
		room := coordinates[i][0]
		x, _ := strconv.Atoi(coordinates[i][1])
		y, _ := strconv.Atoi(coordinates[i][2])
		g.AddRoom(room, x, y)
	}

	// add every connection
	for i := range links {
		from := links[i][0]
		to := links[i][1]
		g.AddEdge(from, to)
		g.AddEdge(to, from)
	}

	// print every room
	// for _, i := range coordinates {
	// 	fmt.Println(g.Vertices[i[0]])
	// }
}
func Dfs(g *Graph, start string) {
	fmt.Println(g.Vertices[start])
	for _, n := range g.Vertices[start].Link {
		Dfs(g, n.Name)
		return
	}
}

func main() {
	path := "example/"
	s := []string{"example00.txt", "example01.txt", "example02.txt", "example03.txt", "example04.txt", "example05.txt", "example06.txt", "example07.txt", "example08.txt", "example09.txt", "example10.txt", "badexample00.txt", "badexample01.txt", "badexample03.txt", "test0.txt"}
	g := NewGraph() // init the graph

	// validate the file
	a, f := function.ValidateFile(path + s[5])
	if !(a) {
		fmt.Println(f[0]) // this will print an error message
		return
	}
	info, coordinates, links := function.Clean(f)
	g.Populate(coordinates, links)
	g.Vertices[info.End].Link = map[string]*Room{}
	fmt.Print(g.Vertices[info.End])
}
