package helper

func SliceContains(slice []string, elem string) bool {
	for _, t := range slice {
		if t == elem {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []string, elem string) []string {
	idx := -1
	for i, el := range slice {
		if el == elem {
			idx = i
			break
		}
	}
	if idx > -1 {
		// order matters, not efficient
		slice = append(slice[:idx], slice[idx+1:]...)
	}
	return slice
}

func RemoveByIdx(s []int, i int) []int {
	// order doesnt matter
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
