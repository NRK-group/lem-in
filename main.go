package main

import (
	"fmt"
	"lemin/function"
	"os"
	"strconv"
)

// https://git.learn.01founders.co/Adriell/lem-in
// https://git.learn.01founders.co/root/public/src/branch/master/subjects/lem-in/audit
type Graph struct {
	Rooms map[string]*Room
}
type Room struct {
	Name    string
	X       int
	Y       int
	Visited bool
	Links   []*Room
}

//AddRoom is a method of the Graph struct that will
//receives name of the room, x coordinates and y coordinates
//and append it to the &Room
func (g *Graph) AddRoom(name string, x, y int) {
	g.Rooms[name] = &Room{
		Name:    name,
		X:       x,
		Y:       y,
		Visited: false,
	}
}

//AddLinks is a method of the Graph struct that will
//receive two room name and connect both rooms to each other
func (g *Graph) AddLinks(from, to string) {
	fromRoom := g.Rooms[from]
	toRoom := g.Rooms[to]
	fromRoom.Links = append(fromRoom.Links, toRoom)
}

// PrintGraph is a graph method that
//print the Room and their links for visualization
func (g *Graph) PrintGraph() {
	for _, v := range g.Rooms {
		fmt.Printf("\nRoom %v : ", v.Name)
		for _, v := range v.Links {
			fmt.Printf(" %v ", v.Name)
		}
	}
	fmt.Println()
}

//NewGraph is a function that initialise a new
//graph by creating an empty room.
func NewGraph() *Graph {
	return &Graph{
		Rooms: map[string]*Room{},
	}
}

//Populate is graph method that populates the graph by adding the Room Using the AddRoom method
//and Adding the Links by the AddLinks metghod.
func (g *Graph) Populate(FarmInfo *function.FarmInfo) *Graph {

	//The code below loops through the FarmInfo.Coordinates
	//and get the room name, x, and y coordinates and
	//add it to the graph using the AddRoom method.
	for i := range FarmInfo.Coordinates {
		room := FarmInfo.Coordinates[i][0]               // room name == index 0
		x, _ := strconv.Atoi(FarmInfo.Coordinates[i][1]) // x coordinates == index 1
		y, _ := strconv.Atoi(FarmInfo.Coordinates[i][2]) // y coordinates == index 2
		g.AddRoom(room, x, y)
	}

	//The code below loops through the FarmInfo.Links
	//and get the 0 index and 1 index of each links
	//and add the the links to the graph using the AddLinks method
	for i := range FarmInfo.Links {
		from := FarmInfo.Links[i][0]
		to := FarmInfo.Links[i][1]

		//The code below makes the graph undirected
		g.AddLinks(from, to)
		// if to != FarmInfo.End && from != FarmInfo.Start {
		g.AddLinks(to, from)
		// }
	}
	return g
}

//IsVisited is a method of *Graph that checks if the Room is visited.
//It will return true if it is visited and false if not
func (g *Graph) IsVisited(name string) bool {
	return g.Rooms[name].Visited
}

//MakeVisited is a *Graph struct method that make the path visited in the graph
func (g *Graph) MakeVisited(start, end string, path Array) {
	for _, name := range path {
		if start != name && end != name {
			g.Rooms[name].Visited = true
		}
	}
}

//Array is an slices of slices.
//Initialize a global slice to create a method into it.
type Array []string

//HasPropertyOf is a method of Array that loops through the slice and
//check if the slice contains the specific string.
//It will receice a string parameter and retrn true if the slice contains the string and
// false if not.
func (arr Array) HasPropertyOf(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func (g *Graph) ShortestPath(start, end string, path Array) []string {
	if _, exist := g.Rooms[start]; !exist {
		return path
	}
	path = append(path, start)
	if start == end {
		return path
	}
	shortest := make(Array, 0)
	var newPath Array
	for _, node := range g.Rooms[start].Links {
		if !(g.IsVisited(node.Name)) && !path.HasPropertyOf(node.Name) {
			newPath = g.ShortestPath(node.Name, end, path)
			// fmt.Println(newPath)
			if len(newPath) > 0 {
				if newPath.HasPropertyOf(start) && newPath.HasPropertyOf(end) { //example04.txt && example05.txt is not working with this code
					if len(shortest) == 0 || (len(newPath) < len(shortest)) { //example04.txt && example05.txt is working with this code
						shortest = newPath
						// fmt.Println(newPath, shortest)
					}
					// return newPath
				}

			}
		}
	}
	return shortest
}

func main() {
	FilePath := os.Args[1]
	ok, file := function.ValidateFile(FilePath) // validation
	if !(ok) {                                  //checks if the file is valid
		fmt.Println(file[0])
		return
	}
	FarmInfo := function.CleanData(file) //assign the
	// fmt.Println(FarmInfo)
	lemin := NewGraph()      //init lemin as a empty graph
	lemin.Populate(FarmInfo) // populate the lemin using the Populate method
	// fmt.Println(lemin)
	lemin.PrintGraph()
	AntsPaths := [][]string{} //container for the paths
	var p Array               //init p for the parameter of the ShortestPath method
	var path Array            // container for the shortest path
	for len(AntsPaths) != len(lemin.Rooms[FarmInfo.Start].Links) {
		path = lemin.ShortestPath(FarmInfo.Start, FarmInfo.End, p) //look for the path
		AntsPaths = append(AntsPaths, path)
		lemin.MakeVisited(FarmInfo.Start, FarmInfo.End, path)
	}
	fmt.Print(AntsPaths)
}
