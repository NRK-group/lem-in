package main

import (
	"fmt"
	"io/ioutil"
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
func (g *Graph) MakeVisited(start, end string, path Array, make bool) {
	for _, name := range path {
		if start != name && end != name {
			g.Rooms[name].Visited = make
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

func (g *Graph) FindPath(start, end string, path Array, swtch bool) []string {
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
			newPath = g.FindPath(node.Name, end, path, swtch)
			// fmt.Println(newPath)
			if len(newPath) > 0 {
				if swtch {
					if newPath.HasPropertyOf(start) && newPath.HasPropertyOf(end) { //example04.txt && example05.txt is not working with this code
						if len(shortest) == 0 || (len(newPath) < len(shortest)) { //example04.txt && example05.txt is working with this code
							shortest = newPath
						}
					}
				}
				if !(swtch) {
					if newPath.HasPropertyOf(start) && newPath.HasPropertyOf(end) { //example04.txt && example05.txt is not working with this code
						return newPath
					}
				}

			}
		}
	}
	return shortest
}

//RoomNameList is a method of the graph that return an array
//of room name from the graph
func (g *Graph) RoomNameList() []string {
	temp := []string{}
	for i := range g.Rooms {
		temp = append(temp, i)
	}
	return temp
}

func (g *Graph) PathList(start, end string, swtch bool) [][]string {
	AntsPaths := [][]string{} //container for the paths
	var p Array               //init p for the parameter of the ShortestPath method
	var path Array            // container for the shortest path
	cnt := 0
	c := 0
	for cnt != len(g.Rooms[start].Links) {
		path = g.FindPath(start, end, p, swtch) //look for the path
		if len(path) != 0 {
			if len(AntsPaths) == 0 {
				AntsPaths = append(AntsPaths, path)
			} else if len(AntsPaths[cnt-1]) == len(path) {
				for i := 0; i < len(path); i++ {
					if AntsPaths[cnt-1][i] == path[i] {
						c++
					}
				}
				if c != len(path) {
					AntsPaths = append(AntsPaths, path)
				}
			} else {
				AntsPaths = append(AntsPaths, path)
			}
		}
		g.MakeVisited(start, end, path, true)
		cnt++
	}

	return SortPaths(AntsPaths)
}

func SortPaths(path [][]string) [][]string {
	for i := range path {
		for j := range path {
			if len(path[i]) < len(path[j]) {
				path[i], path[j] = path[j], path[i]
			}
		}
	}
	return path
}

func ComparePaths(AntsPaths, AntsPaths2 [][]string) [][]string {
	if len(AntsPaths) > len(AntsPaths2) {
		return AntsPaths
	} else if len(AntsPaths) < len(AntsPaths2) {
		return AntsPaths2
	} else {
		antp1 := 0
		antp2 := 0
		for _, paths := range AntsPaths {
			antp1 = antp1 + len(paths)
		}
		for _, paths := range AntsPaths2 {
			antp2 = antp2 + len(paths)
		}

		if antp1 < antp2 {
			return AntsPaths
		} else {
			return AntsPaths2
		}
	}
}

func main() {
	FilePath := os.Args[1]
	s, _ := os.Open(FilePath)            // open the file
	f, _ := ioutil.ReadAll(s)            // read the file
	ok, file := function.ValidateFile(f) // validation
	if !(ok) {                           //checks if the file is valid
		fmt.Println(file[0]) //error message
		return
	}
	FarmInfo := function.CleanData(file) //assign the
	lemin := NewGraph()                  //init lemin as a empty graph
	lemin.Populate(FarmInfo)             // populate the lemin using the Populate method
	// lemin.PrintGraph()
	AntsPaths := lemin.PathList(FarmInfo.Start, FarmInfo.End, true)
	lemin.MakeVisited(FarmInfo.Start, FarmInfo.End, lemin.RoomNameList(), false)
	AntsPaths2 := lemin.PathList(FarmInfo.Start, FarmInfo.End, false)
	a := ComparePaths(AntsPaths, AntsPaths2)
	if len(a) == 0 {
		fmt.Println("ERROR: invalid data format, no path found")
		return
	}
	fmt.Println(string(f))
	fmt.Println()
	fmt.Println(a)
}
