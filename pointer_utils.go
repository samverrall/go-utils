package goutils

// ToPointer takes any value and returns a pointer of said value.
func ToPointer[T any](value T) *T {
	return &value
}
