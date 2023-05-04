package main

import (
	"bufio"
	"fmt"
	"linear/avaliacao_validacao"
	"linear/carrega_csv"
	"linear/gradiente"

	"os"
	"strconv"
)

func main() {
	var hVarDep string
	var hVarIndep []string
	var varDepTest, varDep, coefs []float64
	var varIndep, varIndepTest [][]float64
	var dims []int
	sair := false

	//Debug
	/*	_hVarDep, _varDep, _varDepTest, _hVarIndep, _varIndep, _varIndepTest, err := carrega_csv.CarregaCSV("./assets/training.csv", 0)
		if err != nil {
			fmt.Println(err)
		}
		hVarDep = _hVarDep
		varDep = _varDep
		varDepTest = _varDepTest
		hVarIndep = _hVarIndep
		varIndep = _varIndep
		varIndepTest = _varIndepTest

		coefs = []float64{}
		dims = []int{0}
		fmt.Println("")
		fmt.Println("Arquivos particionados")
		fmt.Println("Arquivo de treino: ", len(varIndep), " linhas")
		fmt.Println("Arquivo de teste: ", len(varIndepTest), "linhas")
		fmt.Println("Qtd Variáveis independentes: ", len(varIndep[0]), " colunas")
		_coefs := gradientedescendente.CalcCoef(varIndep, varDep, dims, 1000, 0.3)
		coefs = _coefs
		fmt.Println("")
		imprimeFormula(hVarDep, hVarIndep, dims, coefs)
		_, _, acuracia, BCE := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndep, varDep, dims)
		fmt.Println("")
		//			fmt.Println("MAE = ", MAE)
		//			fmt.Println("MSE = ", MSE)
		fmt.Println("Binary Cross Entropy = ", BCE)
		fmt.Println("Acuracia = ", acuracia)
		fmt.Println("")

		_coefs2 := gradiente.CalcGradient(varIndep, varDep, dims, 1000, 0.3)
		coefs = _coefs2
		fmt.Println("")
		imprimeFormula(hVarDep, hVarIndep, dims, coefs)
		_, _, acuracia2, BCE2 := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndep, varDep, dims)
		fmt.Println("")
		//			fmt.Println("MAE = ", MAE)
		//			fmt.Println("MSE = ", MSE)
		fmt.Println("Binary Cross Entropy = ", BCE2)
		fmt.Println("Acuracia = ", acuracia2)
		fmt.Println("")
		sair = true */
	// fim debug

	for !sair {
		fmt.Println("*******************************************************************************************")
		fmt.Println("Opções: ")
		fmt.Println("0 > Sair ")
		fmt.Println("1 > Carregar arquivo e particionar treino e teste ")
		fmt.Println("2 > Calcula o modelo")
		fmt.Println("3 > Avaliação do modelo")
		fmt.Println("4 > Validação do modelo (acuracia em comparacao a base de teste)")

		opcao := bufio.NewScanner(os.Stdin)
		opcao.Scan()
		fmt.Println("")
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println("")

		switch {
		case opcao.Text() == "0":
			sair = true
		case opcao.Text() == "1":
			fmt.Println("")
			fmt.Println(">> Carregar Arquivo")
			fmt.Println("")
			fmt.Println("Qual o nome do arquivo .csv?")
			arq := bufio.NewScanner(os.Stdin)
			arq.Scan()
			fmt.Println("")
			fmt.Println("Arquivo de teste separado? (s/n)")
			opcaoTeste := bufio.NewScanner(os.Stdin)
			opcaoTeste.Scan()
			if opcaoTeste.Text() == "S" || opcaoTeste.Text() == "s" {
				fmt.Println("Qual o nome do arquivo de teste .csv?")
				arqTeste := bufio.NewScanner(os.Stdin)
				arqTeste.Scan()
				_hVarDep, _varDep, _, _hVarIndep, _varIndep, _, err := carrega_csv.CarregaCSV("./assets/"+arq.Text()+".csv", 0)
				if err != nil {
					fmt.Println(err)
				}
				_, _varDepTest, _, _, _varIndepTest, _, err := carrega_csv.CarregaCSV("./assets/"+arqTeste.Text()+".csv", 0)
				if err != nil {
					fmt.Println(err)
				}
				hVarDep = _hVarDep
				varDep = _varDep
				varDepTest = _varDepTest
				hVarIndep = _hVarIndep
				varIndep = _varIndep
				varIndepTest = _varIndepTest
			} else {
				fmt.Println("")
				fmt.Println("Qual o percentual utilizado para teste? Formato: 99")
				opcao := bufio.NewScanner(os.Stdin)
				opcao.Scan()
				testPerc, err := strconv.ParseFloat(opcao.Text(), 64)

				_hVarDep, _varDep, _varDepTest, _hVarIndep, _varIndep, _varIndepTest, err := carrega_csv.CarregaCSV("./assets/"+arq.Text()+".csv", testPerc/100)
				if err != nil {
					fmt.Println(err)
				}
				hVarDep = _hVarDep
				varDep = _varDep
				varDepTest = _varDepTest
				hVarIndep = _hVarIndep
				varIndep = _varIndep
				varIndepTest = _varIndepTest

			}
			coefs = []float64{}
			fmt.Println("")
			fmt.Println("Arquivos particionados")
			fmt.Println("Arquivo de treino: ", len(varIndep), " linhas")
			fmt.Println("Arquivo de teste: ", len(varIndepTest), "linhas")
			fmt.Println("Qtd Variáveis independentes: ", len(varIndep[0]), " colunas")
			listaDimensoes(hVarIndep)
		case opcao.Text() == "2":
			fmt.Println("")
			fmt.Println(">> Calcula regressão logistica")

			fmt.Println("")
			fmt.Println("Informe as dimensões (informe o número <enter> para finalizar):")
			listaDimensoes(hVarIndep)
			fmt.Println(">")
			_dims := make([]int, 0, len(hVarIndep))
			sair := false
			for !sair {
				opcao := bufio.NewScanner(os.Stdin)
				opcao.Scan()
				if opcao.Text() == "" {
					sair = true
					continue
				}
				dim, err := strconv.Atoi(opcao.Text())
				if err != nil {
					fmt.Println(err)
				}
				_dims = append(_dims, dim)
			}
			dims = _dims
			fmt.Println("")
			fmt.Println("Qual a quantidade de iteracoes?")
			_passos := bufio.NewScanner(os.Stdin)
			_passos.Scan()
			passos, err := strconv.Atoi(_passos.Text())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("")
			fmt.Println("Qual a taxa de aprendizagem?")
			_taxa := bufio.NewScanner(os.Stdin)
			_taxa.Scan()
			taxa, err := strconv.ParseFloat(_taxa.Text(), 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_coefs := gradiente.CalcGradient(varIndep, varDep, dims, passos, taxa)
			coefs = _coefs
			fmt.Println("")
			imprimeFormula(hVarDep, hVarIndep, dims, coefs)
			fmt.Println("")
		case opcao.Text() == "3":
			fmt.Println("")
			fmt.Println(">> Avaliação do modelo")
			if len(coefs) == 0 {
				fmt.Println("")
				fmt.Println(">> Modelo não calculado")
				break
			}
			_, _, acuracia, BCE := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndep, varDep, dims)
			fmt.Println("")
			fmt.Println("Binary Cross Entropy = ", BCE)
			fmt.Println("Acurácia = ", acuracia)
			fmt.Println("")
		case opcao.Text() == "4":
			fmt.Println("")
			fmt.Println(">> Validação do modelo")
			if len(coefs) == 0 {
				fmt.Println("")
				fmt.Println(">> Modelo não calculado")
				break
			}
			yEstimado, yClass, acuracia, BCE := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndepTest, varDepTest, dims)
			fmt.Println("")
			fmt.Println("Observado => Estimado => classificado")
			for i, y := range yClass {
				fmt.Println(varDepTest[i], " => ", strconv.FormatFloat(yEstimado[i], 'f', -1, 32), " => ", y)
			}
			fmt.Println("")
			fmt.Println("Binary Cross Entropy = ", BCE)
			fmt.Println("Acuracia = ", acuracia)
			fmt.Println("")
			imprimeFormula(hVarDep, hVarIndep, dims, coefs)

		}
	}

}

func listaDimensoes(hVarIndep []string) {
	for i, h := range hVarIndep {
		fmt.Println(i, ")", h)
	}
}

func imprimeFormula(hVarDep string, hVarIndep []string, dims []int, coefs []float64) {
	formula := "probabilidade de " + hVarDep + " = 1 / ( 1 + exp(- b0 "
	for i := 1; i < len(coefs); i++ {
		formula = formula + " - b" + strconv.Itoa(i) + " * " + hVarIndep[dims[i-1]]
	}
	formula = formula + " )"
	fmt.Println("")
	fmt.Println("Formula:")
	fmt.Println(formula)
	fmt.Println("")
	for i := 0; i < len(coefs); i++ {
		fmt.Println("b"+strconv.Itoa(i), " = ", strconv.FormatFloat(coefs[i], 'f', -1, 32))
	}
	fmt.Println("")

}
