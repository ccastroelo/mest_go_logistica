package normaliza

import (
	"errors"
	"math"
)

func Normaliza(varIndep [][]float64) ([][]float64, []float64, []float64, error) {
	linhas := len(varIndep)
	col := len(varIndep[0])

	minMax := make([]float64, col)
	means := make([]float64, col)
	norms := make([][]float64, linhas)
	for i := 0; i < linhas; i++ {
		norms[i] = make([]float64, col)
	}
	for j := 0; j < col; j++ {
		var max float64
		min := math.MaxFloat64
		soma := 0.0
		for i := 0; i < linhas; i++ {
			if max < varIndep[i][j] {
				max = varIndep[i][j]
			}
			if min > varIndep[i][j] {
				min = varIndep[i][j]
			}
			soma += varIndep[i][j]
		}
		means[j] = soma / float64(linhas)
		minMax[j] = max - min
		if minMax[j] == 0 {
			return nil, nil, nil, errors.New("Min max igual a 0")
		}
		for i := 0; i < linhas; i++ {
			norms[i][j] = (varIndep[i][j] - means[j]) / minMax[j]
		}
	}

	return norms, means, minMax, nil
}
