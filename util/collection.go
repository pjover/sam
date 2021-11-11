package util

func StringInList(str string, list []string) bool {
	for _, b := range list {
		if b == str {
			return true
		}
	}
	return false
}
