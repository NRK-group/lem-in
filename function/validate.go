package function

import (
	"strconv"
	"strings"
)

//This fucntion validate the file.
//It checks if the file is not empty, the number of ants is valid, and
//the file has start and end points.
//It will return true and the content of the file in an array if the file is formatted correctly and
// return false and an error message for invalid file.
func ValidateFile(fileByte []byte) (bool, []string) {
	splitF := strings.Split(string(fileByte), "\n") // split the data with newline and assign it to splitF
	if len(splitF) == 1 {                           //check if the length of the file is 1 and return error
		return false, []string{"ERROR: invalid data format, empty file"}
	}
	num, err := strconv.Atoi(splitF[0])
	if num == 0 || splitF[0] == "" || err != nil { // checks if their is a valid number of ants
		return false, []string{"ERROR: invalid data format, invalid number of Ants"}
	}
	if !(strings.Contains(string(fileByte), "##start")) { // checks if their is a start room
		return false, []string{"ERROR: invalid data format, no start room found"}
	}
	if !(strings.Contains(string(fileByte), "##end")) { // checks if their is a end room
		return false, []string{"ERROR: invalid data format, no end room found"}
	}
	return true, splitF // return true and the whole array of file
}
