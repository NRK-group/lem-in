package function

//SortPaths is function that sort the slice inside the slice base on their length.
//This function will return the ordered paths. The algoriothms that used in this function is bubble sort.
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

//SortStringPaths is function that sort the slice inside the slice.
//This function will return the ordered paths. The algoriothms that used in this function is bubble sort.
func SortStringPaths(path []string) []string {
	for i :=0; i < len(path)-1; i++ {
			//if path[i] is less than path[j] swap them
			if len(path[i]) < len(path[i+1]) {
				path[i], path[i+1] = path[i+1], path[i]
			}
	}
	return path
}