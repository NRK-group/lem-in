package function

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func ValidateFile(file string) (bool, error) {
	s, _ := os.Open(file)
	f, _ := ioutil.ReadAll(s)
	splitF := strings.Split(string(f), "\n")
	if len(splitF) == 1 {
		return false, errors.New("ERROR: empty file")
	}
	if splitF[0] == "0" || splitF[0] == "" {
		return false, errors.New("ERROR: invalid data format, invalid number of Antss")
	}
	if !(strings.Contains(string(f), "##start")) {
		return false, errors.New("ERROR: invalid data format, no start room found")
	}
	if !(strings.Contains(string(f), "##end")) {
		return false, errors.New("ERROR: invalid data format, no end room found")
	}
	return true, nil
}
