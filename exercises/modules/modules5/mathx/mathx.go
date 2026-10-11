// Package mathx is a small module of number helpers.
package mathx

// Mean returns the arithmetic mean of xs, or 0 for no values.
func Mean(xs ...float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}
