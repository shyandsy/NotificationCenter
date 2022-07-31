package helper

func Contains(array []string, element string) bool {
	found := false
	for _, item := range array {
		if item == element {
			found = true
			break
		}
	}
	return found
}
