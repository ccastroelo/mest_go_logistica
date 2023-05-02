package matrizes

func Media(elem []float64) float64 {
	var total float64 = 0
	for _, e := range elem {
		total = total + e
	}
	return total / float64(len(elem))
}
