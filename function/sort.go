package function

//SortPaths is function that sort the slice inside the slice base on their length.
//This function will return the ordered paths. The algoriothms that used in this function
//is bubble sort.
func SortPaths(path [][]string) [][]string {
	for i := range path {
		for j := range path {
			//if path[i] is less than path[j] swap them
			if len(path[i]) < len(path[j]) {
				path[i], path[j] = path[j], path[i]
			}
		}
	}
	return path
}
