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

//RoomNameList is a method of the graph that return an array
//of room name from the graph
func (g *Graph) RoomNameList() []string {
	temp := []string{}
	for i := range g.Rooms {
		temp = append(temp, i)
	}
	return temp
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

//Array is an slices of string.
//Initialize a global slice to create a method into slice.
type Array []string

type DArray [][]string

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

//FindPath is a method of the *Graph struct that find paths from the start to end.
func (g *Graph) FindPath(start, end string, path Array, swtch bool) []string {
	var newPath Array
	shortest := make(Array, 0)

	if _, exist := g.Rooms[start]; !exist {
		return path
	}

	path = append(path, start)
	if start == end {
		return path
	}

	for _, node := range g.Rooms[start].Links {
		//the if statement below checks if the current node is not visited \
		// and if the current path dont have the same room
		if !(g.IsVisited(node.Name)) && !path.HasPropertyOf(node.Name) {
			newPath = g.FindPath(node.Name, end, path, swtch) //recursion
			if len(newPath) > 0 {
				//if the swtch is true it will find the shortest path in the graph
				if swtch {
					if newPath.HasPropertyOf(start) && newPath.HasPropertyOf(end) {
						if len(shortest) == 0 {
							shortest = newPath
						}
						if len(newPath) < len(shortest) {
							shortest = newPath
							//fmt.Println(shortest)
						}

					}
				}

				//if the switch is false it will return the first path it's finds
				if !(swtch) {
					if newPath.HasPropertyOf(start) && newPath.HasPropertyOf(end) {
						return newPath
					}
				}

			}
		}
	}

	return shortest //this will be returned if the swtch is true
}

//PathList is a method of the *Graph struct that return a slice of slices of paths
func (g *Graph) PathList(start, end string, swtch bool) [][]string {
	AntsPaths := [][]string{} //container for the paths
	var p Array               //init p for the parameter of the ShortestPath method
	var path Array            // container for the shortest path
	cnt := 0
	c := 0

	//the for loop below will loop until cnt is not equal to the length of the adjacent list of the start room
	for cnt != len(g.Rooms[start].Links) {
		path = g.FindPath(start, end, p, swtch) //look for the path
		if len(path) != 0 {
			if len(AntsPaths) == 0 {
				AntsPaths = append(AntsPaths, path)
			} else if len(AntsPaths[cnt-1]) == len(path) {
				//if the current path and the previous path have the same length
				//checks if the path is not similar to the previous paths
				//if it is not similar append the path into the AntsPaths
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
		g.MakeVisited(start, end, path, true) // make the paths visiteds
		cnt++
	}
	return AntsPaths
}

//ComparePaths is a function that receives two slices of slice of paths and compare which one is the
//best path to use. This function will return the best paths to use in asceding order based on path length
func ComparePaths(AntsPaths, AntsPaths2 [][]string) [][]string {
	//this function will check which one had more paths and return it
	if len(AntsPaths) > len(AntsPaths2) {
		return AntsPaths
	} else if len(AntsPaths) < len(AntsPaths2) {
		return function.SortPaths(AntsPaths2)
	} else {
		//if length of both slices of slice of path is equal to each order it will check which one had the
		//less room inside the slices of slice of path and return it
		antp1 := 0
		antp2 := 0
		for _, paths := range AntsPaths {
			antp1 = antp1 + len(paths)
			antp2 = antp2 + len(paths)
		}
		if antp1 < antp2 {
			return function.SortPaths(AntsPaths)
		} else {
			return function.SortPaths(AntsPaths2)
		}
	}
}

//MoveOfAnts is function that will receive the name of the ant and the path that their taking and return a slice
//that contains the ants movement. For example, the ant name is 1 and the path is [A0 A1 A2 end]
//the result will be:
//  [L1-A0 L1-A1 L1-A2 L1-end]
func MoveOfAnts(nameOfAnt int, paths []string) []string {
	result := []string{}
	str := ""
	antName := strconv.Itoa(nameOfAnt)
	for _, room := range paths {
		str = "L" + antName + "-" + room
		result = append(result, str)
		str = ""
	}
	return result
}

//LenofMoves is the function that cacluates the exact
//number of line it will take to print the movements of the ants downward.
//For example, the container [][][]string contains the slices below.
//  [[[L1-2 L1-3 L1-1] [L2-2 L2-3 L2-1] [L3-2 L3-3 L3-1] [L4-2 L4-3 L4-1]]]
//The container contains one path so the maxlen will be len of the container[0] to get the number of ants inside that path
//and the lenOfMove will be the len of the container[0][0] to get the number of moves of each ants in that path.
func LenOfMoves(container [][][]string) int {
	result := 0
	maxLen := 0
	pos := 0
	if len(container) > 1 {
		for i := range container {
			for j := range container {
				if len(container[i]) < len(container[j]) {
					maxLen = len(container[j])
					pos = j
				}
			}
		}
	} else {
		maxLen = len(container[0]) //assign to the first path if the number of path in the container is one
	}
	lenOfMove := len(container[pos][0])
	//the number of ants in the path - 1 + the len of ants movemnts in that path
	//will give you the exact amount of line to print it downwards
	result = (maxLen - 1) + (lenOfMove)
	return result
}

//PrintAntsMoves is the function that print the ants movements
//This will receives [][][]string that contains the movement of that ants and print it downwards
//for example, the container [][][]string contains
//  [[[L1-2 L1-3 L1-1] [L2-2 L2-3 L2-1] [L3-2 L3-3 L3-1] [L4-2 L4-3 L4-1]]]
//it will be converted to the result below.
//  L1-3 L2-2
//  L1-1 L2-3 L3-2
//  L2-1 L3-3 L4-2
//  L3-1 L4-3
//  L4-1
func PrintAntsMoves(container [][][]string, lenofMoves int) string {
	result := ""
	antsMoves := make([][]string, lenofMoves)
	for _, c := range container {
		for j, paths := range c {
			for k, p := range paths {
				antsMoves[j+k] = append(antsMoves[j+k], p)
			}
		}
	}
	for _, a := range antsMoves {
		for _, v := range a {
			result += v + " "
		}
		result += "\n"
	}
	return result
}

//PathsSeletion pick the correct amount of paths and the correct amount of ants for each paths
func PathsSeletion(nAnts int, pahts [][]string) [][][]string {
	container := make([][][]string, len(pahts))

	if len(pahts) > 1 {
		cnt := 0
		i := 1
		for i != nAnts+1 {
			if cnt == len(pahts)-1 {
				cnt = 0
			}
			x := len(pahts[cnt]) + len(container[cnt])
			y := len(pahts[cnt+1]) + len(container[cnt+1])
			if !(x > y) {
				container[cnt] = append(container[cnt], MoveOfAnts(i, pahts[cnt][1:]))
			} else {
				if cnt == len(pahts)-1 {
					cnt = 0
					container[0] = append(container[0], MoveOfAnts(i, pahts[0][1:]))
				} else {
					cnt++
					container[cnt] = append(container[cnt], MoveOfAnts(i, pahts[cnt][1:]))
				}
			}
			i++
		}
	} else {
		i := 1
		for i != nAnts+1 {
			container[0] = append(container[0], MoveOfAnts(i, pahts[0][1:]))
			i++
		}
	}
	return container
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
	FarmInfo := function.CleanData(file)
	lemin := NewGraph()      //init lemin as a empty graph
	lemin.Populate(FarmInfo) // populate the lemin using the Populate method
	// lemin.PrintGraph()
	AntsPaths1 := lemin.PathList(FarmInfo.Start, FarmInfo.End, true)
	lemin.MakeVisited(FarmInfo.Start, FarmInfo.End, lemin.RoomNameList(), false)
	AntsPaths2 := lemin.PathList(FarmInfo.Start, FarmInfo.End, false)
	a := ComparePaths(AntsPaths1, AntsPaths2)
	if len(a) == 0 {
		fmt.Println("ERROR: invalid data format, no path found")
		return
	}
	fmt.Println(string(f))
	fmt.Println()
	container := PathsSeletion(FarmInfo.NumAnts, a)
	fmt.Print(PrintAntsMoves(container, LenOfMoves(container)))
}
