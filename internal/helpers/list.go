package helpers

func Contains[T comparable](list []T, val T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}

	return false
}
