package t

func Pointer[T any](input *T, defaultVal T) T {
	if input != nil {
		return *input
	}

	return defaultVal
}
