package gradiente

import (
	"math"
)

func CalcGradient(varIndep [][]float64, varDep []float64, dims []int, passos int, taxa float64) []float64 {

	// pode-se gerar os coeficientes iniciais randomicamente
	//	s := rand.NewSource(time.Now().UnixNano())
	//	r := rand.New(s)

	coefs := make([]float64, len(dims)+1)

	// inicializa os coeficientes com valores aleatórios
	//	for i := 0; i < len(dims)+1; i++ {
	//		coefs[i] = r.Float64()
	//	}

	// o B0 (intercepto) não possui X, então utiliza-se um X0 = 1, então é acrescentado 1
	// para a formula ficar b0*1 + b1x1 + b2x2 ... bnxn
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

	// inicia o laço para otimizacao
	for i := 0; i < passos; i++ {

		coefsAux := make([]float64, len(coefs))
		// varre a base de treino atualizando os coeficientes auxiliares
		for j := 0; j < len(varDep); j++ {
			// xn = os valores de x para essa iteracao
			xn := mx[j]

			// calcula o y estimado com base nos coeficientes atuais
			yEst := hipotese(somaCoefX(coefs, xn))

			// obtem o erro entre o y encontrado e o observado
			erro := yEst - varDep[j]
			// atualiza os coeficientes auxiliares com base nos valores de xn * erro
			// a ideia é q o erro vá diminuindo, dminuindo assim o valor (tamanho dos passos)
			// de atualizacao do coefAux e de coefs
			for k := 0; k < len(coefsAux); k++ {
				coefsAux[k] = coefsAux[k] + erro*xn[k]
			}
		}
		// atualiza os coeficientes com base na taxa de aprendizado e nos valores dos auxiliares encontrados na
		// varrida anterior
		for k := 0; k < len(coefs); k++ {
			coefs[k] = coefs[k] - taxa*coefsAux[k]/float64(len(varDep))
		}

	}
	return coefs
}

func somaCoefX(coefs []float64, linha []float64) float64 {
	soma := 0.0
	for i := 0; i < len(linha); i++ {
		soma += coefs[i] * linha[i]
	}
	return soma
}

func hipotese(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
