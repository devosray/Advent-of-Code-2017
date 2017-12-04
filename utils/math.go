package utils

// Why no built-in int abs function >:(
func IntAbs(a int) int{
	if a < 0 {
		return -1 * a
	}
	return a
}