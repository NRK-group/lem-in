package function

import (
	"strconv"
	"strings"
)

//FarmInfo is a struct that will contain all the data of the farm
type FarmInfo struct {
	NumAnts     int
	Start       string
	End         string
	Coordinates [][]string
	Links       [][]string
}

// This function receive a txt file and return a stuct.
// The stuct "*FarmInfo" consist of number of ants, starting points, ending points, coordinates ["room name", "x", "y"],
//and links ["room name", "room name"].
func CleanData(args []string) *FarmInfo {
	ants := args[0]             // get the number of ants
	s := ""                     // get the start values
	e := ""                     // get the end values
	coordinates := [][]string{} //store all the coordinates
	links := [][]string{}       // store all the links
	data := args[1:]            // exclude the number of ants

	for i, w := range data {
		if w == "##start" { // look for the starting coordinates
			s = data[i+1] // store the start value
			data[i] = ""  // remove the ##start
		}
		if w == "##end" { // look for the end coordinates
			e = data[i+1] // store the end value
			data[i] = ""  // remove the ##end
		}
		if len(w) != 0 {
			if w[0] == '#' { // remove all the value that start with #
				data[i] = ""
			}
		}
	}
	for _, w := range data {
		if w != "" { // exclude the empty
			if strings.Contains(w, " ") { // look for the coordinates
				c := strings.Split(w, " ")           // split the string with the space to get the room name, x and y
				coordinates = append(coordinates, c) //store the value
			} else {
				l := strings.Split(w, "-") // split the string with "-" to get the links
				links = append(links, l)   // store the value
			}
		}
	}

	antNum, _ := strconv.Atoi(ants)
	start := strings.Split(s, " ")
	end := strings.Split(e, " ")

	// populate the FarmInfo struct
	FarmInfo := &FarmInfo{
		NumAnts:     antNum,
		Start:       start[0],
		End:         end[0],
		Coordinates: coordinates,
		Links:       links,
	}
	return FarmInfo
}
