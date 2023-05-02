package avaliacao_validacao

import (
	"fmt"
	"math"
)

func AvaliaValidaModelo(coefs []float64, varIndep [][]float64, varDep []float64, dims []int) ([]float64, []float64, float64, float64) {
	//	aMAE := 0.0
	//	aMSE := 0.0
	//	mediaY := matrizes.Media(varDep)
	//	SomaDifMedia := 0.0
	yClass := make([]float64, len(varDep))
	yEstimado := make([]float64, len(varDep))
	truePosNeg := 0
	for i, linha := range varIndep {
		exp := 0.0
		for j, dim := range dims {
			if j == 0 {
				exp = exp + coefs[j]
			}
			exp = exp + coefs[j+1]*linha[dim]
		}
		yClass[i], yEstimado[i] = logistic(exp)
		if yClass[i] == varDep[i] {
			truePosNeg++
		}
		/*		aMAE = aMAE + math.Abs(varDep[i]-yEst)
				aMSE = aMSE + math.Pow((varDep[i]-yEst), 2)
				SomaDifMedia = SomaDifMedia + math.Pow((varDep[i]-mediaY), 2)*/
	}
	acuracia := float64(truePosNeg) / float64(len(varDep))
	bce, ind := bce(varDep, yEstimado)
	if ind > 0 {
		fmt.Println(ind, varDep[ind], yEstimado[ind])

	}
	/*	R2 := 1 - (aMSE / SomaDifMedia)
		MSE := aMSE / float64(len(varDep))
		RMSE := math.Sqrt(MSE)
		MAE := aMAE / float64(len(varDep)) */

	return yEstimado, yClass, acuracia, bce
}

func bce(varDep []float64, yEstimado []float64) (float64, int) {
	r := 0.0
	ind := 0
	for i, y := range varDep {
		if y == 1 && yEstimado[i] == 0 {
			ind = i
		}
		r += y*math.Log(yEstimado[i]) + (1-y)*math.Log(1-yEstimado[i])
	}
	return -r, ind
}

func logistic(x float64) (float64, float64) {
	p := 1.0 / (1.0 + math.Exp(-x))
	if p >= 0.5 {
		return 1.0, p
	}
	return 0.0, p
}
