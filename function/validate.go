package function

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ValidateFile(file string) (bool, []string) {
	s, _ := os.Open(file)
	f, _ := ioutil.ReadAll(s)
	splitF := strings.Split(string(f), "\n")
	if len(splitF) == 1 {
		return false, []string{"ERROR: empty file"}
	}
	num, err := strconv.Atoi(splitF[0])
	if num == 0 || splitF[0] == "" || err != nil {
		return false, []string{"ERROR: invalid data format, invalid number of Ants"}
	}
	if !(strings.Contains(string(f), "##start")) {
		return false, []string{"ERROR: invalid data format, no start room found"}
	}
	if !(strings.Contains(string(f), "##end")) {
		return false, []string{"ERROR: invalid data format, no end room found"}
	}
	return true, splitF
}
func Clean(args []string) (int, []string, []string, [][]string, [][]string) {
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
			// data[i+1] = "" // remove the start value
		}
		if w == "##end" { // look for the end coordinates
			e = data[i+1] // end the start value
			data[i] = ""  // remove the ##end
			// data[i+1] = "" // remove the end value
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
				c := strings.Split(w, " ")
				coordinates = append(coordinates, c) //store the value
			} else {
				l := strings.Split(w, "-")
				links = append(links, l)
			}
		}
	}
	antNum, _ := strconv.Atoi(ants)
	start := strings.Split(s, " ")
	end := strings.Split(e, " ")
	return antNum, start, end, coordinates, links
}
