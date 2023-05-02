package gradientedescendente

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func CalcCoef2(data [][]float64, y []float64, dims []int, num_iters int, alpha float64) []float64 {

	theta := make([]float64, len(dims)+1)

	for i := 0; i < num_iters; i++ {
		//Number of training examples
		m := len(y)
		//Slice helper to calculate our new versions of theta
		thetaTemp := make([]float64, len(theta))

		//Sum (hi-yx)xi
		for rowI := 0; rowI < m; rowI++ {
			// calcula o yEst
			hi := logistic(ComputeHypothesis(data[rowI], theta))

			sumRowI := computeSumRowI(data[rowI], hi, y[rowI])
			for t := 0; t < len(theta); t++ {
				thetaTemp[t] += sumRowI[t]
			}
		}
		//Update theta
		for t := 0; t < len(theta); t++ {
			theta[t] = theta[t] - (alpha/float64(m))*thetaTemp[t]
		}

	}
	return theta
}

func ComputeHypothesis(x []float64, theta []float64) float64 {
	result := theta[0]
	for i := 1; i < len(theta); i++ {
		result += theta[i] * x[i-1]
	}
	return result
}

func ComputeCost(data [][]float64, y []float64, theta []float64) (float64, error) {
	m := len(y)
	sum := 0.0
	for rowI := 0; rowI < m; rowI++ {

		//Sum
		hi := ComputeHypothesis(data[rowI], theta)
		sum += (hi - y[rowI]) * (hi - y[rowI])

	}
	return (1 / float64(2*m)) * sum, nil
}

func computeSumRowI(x []float64, hi float64, yi float64) []float64 {
	theta := make([]float64, len(x)+1)
	theta[0] = hi - yi
	for i := 1; i < len(theta); i++ {
		theta[i] = (hi - yi) * x[i-1]
	}
	return theta
}

func CalcCoef(varIndep [][]float64, varDep []float64, dims []int, passos int, taxa float64) []float64 {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	coefs := make([]float64, len(dims)+1)

	// inicializa os coeficientes com valores aleatórios
	for i := 0; i < len(dims)+1; i++ {
		coefs[i] = r.Float64()
	}

	// a matriz varIndep não possui o xi para o B0, então é acrescentado 1
	// para a formula b0*1 + b1x1 + b2x2 ... bnxn
	mx := make([][]float64, len(varIndep))
	for i := 0; i < len(varIndep); i++ {
		mx[i] = make([]float64, len(dims)+1)
		for j := 0; j < len(dims); j++ {
			if j == 0 {
				mx[i][j] = 1.0
			}
			mx[i][j+1] = varIndep[i][dims[j]]
		}
	}

	// inicia o passo para otimizacao
	for i := 0; i < passos; i++ {
		var somaErroQua float64

		for j := 0; j < len(varDep); j++ {
			// linha recebe os valores de x para essa iteracao
			linha := mx[j]
			// soma b0*1 + b1x1 + b2x2 ....
			somaCoef := somaCoef(coefs, linha)
			// calcula o y estimado com a funcao sigmoide
			yEst := logistic(somaCoef)
			// obtem o erro entre o y encontrado e o observado
			erro := varDep[j] - yEst

			somaErroQua += math.Pow(erro, 2)
			// atualiza os coeficientes com base na taxa de aprendizado, o erro e o y estimado
			for k := 0; k < len(linha); k++ {
				coefs[k] += taxa * erro * yEst * (1 - yEst) * linha[k]
			}
			//			println(bce(varDep, yEst))
		}

	}
	return coefs
}

func somaCoef(coefs []float64, linha []float64) float64 {
	soma := 0.0
	for i := 0; i < len(linha); i++ {
		soma += coefs[i] * linha[i]
	}
	return soma
}

func logistic(x float64) float64 {
	p := 1.0 / (1.0 + math.Exp(-x))
	// retirar
	if p > 1.0 {
		fmt.Println(p)
	}
	return p
}
