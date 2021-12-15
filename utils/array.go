package utils

func ValueInArray(v string, l []string) bool {
	for _, _v := range l {
		if _v == v {
			return true
		}
	}
	return false
}
