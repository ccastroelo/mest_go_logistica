package carrega_csv

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func CarregaCSV(fileName string, testSlicePerc float64) (string, []float64, []float64, []string, [][]float64, [][]float64, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return "", nil, nil, nil, nil, nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return "", nil, nil, nil, nil, nil, err
	}

	// sorteia as linhas do arquivo csv que serão utilizados como dados de teste
	// e armazena os indices no array abaixo
	// quando, do parse do arquivo csv, se a linha i estiver no array abaixo, a linha será carregada
	// para o array de dados de teste
	testSlice, err := getRandomSliceIdx(len(lines)-1, testSlicePerc)
	if err != nil {
		fmt.Println(err)
	}

	// array linha que receberá o array colunas de variavéis independente dos dados de treinamento
	// comprimento = número de registros menos o header e o número  registros selecionados para teste
	varIndep := make([][]float64, len(lines)-len(testSlice)-1)

	// array linha que receberá o array colunas de variavéis independente dos dados de teste
	// comprimento = número de registros selecionados para teste
	varIndepTest := make([][]float64, len(testSlice))

	// array que receberá a variavel independente dos dados de treinamento
	// comprimento = número de registros menos o header e o número  registros selecionados para teste
	varDep := make([]float64, len(lines)-len(testSlice)-1)

	// array que receberá a variavel independente dos dados de teste
	// comprimento = número de registros selecionados para teste
	varDepTest := make([]float64, len(testSlice))

	// array que armazenará os headers das colunas das variaveis independente
	// comprimento = headers sem a primeira coluna (pais) e sem a segunda coluna (variável dependente)
	hVarIndep := make([]string, len(lines[0])-1)

	// variavel que armazenará o header da coluna da variavel dependente
	// header da variável dependente
	hVarDep := ""

	// controle do indice do array que irá receber os dados de teste
	testIdx := 0
	// controle do indice do array que irá receber os dados de treinamento
	trainIdx := 0 // training array control index

	// variavel para facilitar o direcionamento do registro, se de teste ou de treinamento
	selectedForTest := false

	proxIdxTest := -1

	// Loop de leitura das linhas do arquivo csv
	for i, line := range lines {
		// carregamento dos headers para o array de header e para a variavel
		if i == 0 {
			for j := 0; j < len(line)-1; j++ {
				hVarIndep[j] = line[j]
			}
			// carregado o header da ultima coluna, variavel dependente, para a variavel
			hVarDep = line[len(line)-1]
			// pula a linha de hearder
			continue
		}

		if testIdx < len(testSlice) {
			proxIdxTest = testSlice[testIdx]
		}
		// Inicializa a segunda dimensão do arrays de teste e de treinamento das variaveis independente
		if i == proxIdxTest { // verifica se a linha i foi sorteada para ser um registro de teste
			selectedForTest = true
			varIndepTest[testIdx] = make([]float64, len(line)-1) // -1 porque descarta a ultima (var dependente) coluna
			// um registro sorteado foi encontrado, se ainda houverem mais registros a serem separados para teste,
			// incrementa o idx para pegar o número da proxima linha de teste quando o loop voltar
			testIdx++
		} else {
			selectedForTest = false
			varIndep[trainIdx] = make([]float64, len(line)-1) // -1 porque descarta a ultima (var dependente) coluna
			trainIdx++
		}
		// laco das colunas da linha i
		for j, col := range line {
			// as colunas de 1 a 4 vieram como string e com virgula no separador decimal
			// então são tratados de forma diferente
			/*			if j < 5 {
						s := strings.Replace(col, ",", ".", 1)
						f, err := strconv.ParseFloat(s, 64)
						if err != nil {
							return "", nil, nil, nil, nil, nil, err
						}
						if selectedForTest {
							varIndepTest[testIdx-1][j-1] = f
						} else {
							varIndep[trainIdx-1][j-1] = f
						}
					} else {*/
			f, err := strconv.ParseFloat(col, 64)
			if err != nil {
				return "", nil, nil, nil, nil, nil, err
			}
			// se for a antes da ultima coluna, é variavel independente
			// se for a ultima coluna, é variavel dependente
			// se for um, registro selecionado para ser dado de teste,
			// é carregado para o seu devido array
			if j < len(line)-1 {
				if selectedForTest {
					varIndepTest[testIdx-1][j] = f
				} else {
					varIndep[trainIdx-1][j] = f
				}
			} else {
				if selectedForTest {
					varDepTest[testIdx-1] = f
				} else {
					varDep[trainIdx-1] = f
				}
			}
			//			}
		}
	}
	return hVarDep, varDep, varDepTest, hVarIndep, varIndep, varIndepTest, nil
}

// passado um percentual, a funcao abaixo calcula a quantidade de linhas dos dados
//
//	que serão sorteadas para serem dados de testes
//
// são então sorteadas quais linhas e seus indices são retornados em um array
func getRandomSliceIdx(max int, slicePerc float64) ([]int, error) {
	// calcula a quantidade de linha
	qtd := int(float64(max) * slicePerc)
	// cria o array que irá armazenar o número da linha sorteada
	var alreadyUse []int

	// laço enquanto a quantidade de linha não é sorteado
	for len(alreadyUse) < qtd {
		// sorteio
		rand.Seed(time.Now().UnixNano())
		newRand := rand.Intn(max) + 1
		// verifica se aquela linha já foi sorteada, se não, armazena
		if !contains(alreadyUse, newRand) {
			alreadyUse = append(alreadyUse, newRand)
		}
	}
	// ordena
	sort.Ints(alreadyUse)
	return alreadyUse, nil
}

func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}
