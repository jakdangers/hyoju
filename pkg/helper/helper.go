package helper

func EnumToPointer[T any](enum T) *T {
	return &enum
}
