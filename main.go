package main

// SumInts adds together the value of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloat64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
