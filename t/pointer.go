package t

// Pointer .. 指针取值
func Pointer[T any](input *T, defaultVal T) T {
	if input != nil {
		return *input
	}

	return defaultVal
}
