package mqo

import "fmt"

func CalcCoef(varIndep [][]float64, varDep []float64, dim int) []float64 {

	x := make([]float64, len(varIndep))
	for i := 0; i < len(varIndep); i++ {
		x[i] = varIndep[i][dim]
	}
	mediaX := media(x)
	mediaY := media(varDep)
	sXY := 0.0
	sXX := 0.0

	var result []float64
	for i := 0; i < len(x); i++ {
		sXY = sXY + (x[i]-mediaX)*(varDep[i]-mediaY)
		sXX = sXX + (x[i]-mediaX)*(x[i]-mediaX)
	}
	b1 := sXY / sXX
	b0 := mediaY - b1*mediaX
	result = append(result, b0, b1)
	return result

}

func CalcCoefViaMatriz(varIndep [][]float64, varDep []float64, dims []int) []float64 {
	/*
		A := [][]float64{
			{3, 1},
			{5, 2},
		}
		printMatrix(A)
		fmt.Println("*****")
		printMatrix(A)
		fmt.Println("*****")
			b := matrizInversa(A)
				printMatrix(A)
				fmt.Println("*****")
				printMatrix(b) */

	my := make([][]float64, len(varDep))
	for i := 0; i < len(varDep); i++ {
		my[i] = make([]float64, 1)
		my[i][0] = varDep[i]
	}

	mx := make([][]float64, len(varIndep))
	for i := 0; i < len(varIndep); i++ {
		mx[i] = make([]float64, len(dims)+1)
		for j := 0; j < len(dims); j++ {
			if j == 0 {
				mx[i][j] = 1
			}
			mx[i][j+1] = varIndep[i][dims[j]]

		}
	}

	m1 := matrizTransposta(mx)
	m2 := multiplicaMatriz(m1, mx)
	m3 := matrizInversa(m2)
	m4 := multiplicaMatriz(m3, m1)
	m5 := multiplicaMatriz(m4, my)
	m6 := matrizTransposta(m5)

	return m6[0]
}

func media(elem []float64) float64 {
	var total float64 = 0
	for _, e := range elem {
		total = total + e
	}
	return total / float64(len(elem))
}

func matrizTransposta(varIndep [][]float64) [][]float64 {
	mT := make([][]float64, len(varIndep[0]))
	for i := 0; i < len(varIndep[0]); i++ {
		mT[i] = make([]float64, len(varIndep))
		for j := 0; j < len(varIndep); j++ {
			mT[i][j] = varIndep[j][i]
		}
	}
	return mT
}

func multiplicaMatriz(a [][]float64, b [][]float64) [][]float64 {
	linhasA := len(a)
	colsA := len(a[0])
	linhasB := len(b)
	colsB := len(b[0])

	if colsA != linhasB {
		panic("Dimensoes incompativeis")
	}

	mR := make([][]float64, linhasA)
	for i := range mR {
		mR[i] = make([]float64, colsB)
	}

	for i := 0; i < linhasA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				mR[i][j] = mR[i][j] + a[i][k]*b[k][j]
			}
		}
	}

	return mR
}

func matrizInversa(m [][]float64) [][]float64 {

	mAux := clone(m)
	c := len(mAux)

	// cria e preenche identidade
	mI := make([][]float64, c)
	for i := range mI {
		mI[i] = make([]float64, c)
		mI[i][i] = 1
	}

	for k := 0; k < c; k++ {
		for i := 0; i < c; i++ {
			if i != k {
				elem := mAux[i][k] / mAux[k][k]
				for j := 0; j < c; j++ {
					mAux[i][j] = mAux[i][j] - elem*mAux[k][j]
					mI[i][j] = mI[i][j] - elem*mI[k][j]
				}
			}
		}

		aux := 1 / mAux[k][k]
		for j := 0; j < c; j++ {
			mAux[k][j] = mAux[k][j] * aux
			mI[k][j] = mI[k][j] * aux
		}
	}

	return mI
}

func imprimeMatriz(A [][]float64) {
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%12.4f ", A[i][j])
		}
		fmt.Println()
	}
}

func clone(m [][]float64) [][]float64 {
	x := make([][]float64, len(m))
	for i := range m {
		x[i] = make([]float64, len(m[0]))
		for j := range m[i] {
			x[i][j] = m[i][j]
		}
	}
	return x
}
