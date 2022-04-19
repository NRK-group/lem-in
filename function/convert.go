package function

import "strings"

//ConvertToArray convert a string of the to an array
func ConvertToArray(s string) []string {
	arr := strings.Fields(s[1 : len(s)-1])
	return arr
}
