package main

import "fmt"

func main() {

	/* ARRAYS:
	Array é uma estrutura de dados que armazena uma coleção de variaveis do mesmo tipo
	em uma sequência de posições de memória. */

	// Como declarar um array e utilizar um array:

	var exemplo1 [5]string
	exemplo1[0] = "valor do indice 0"
	exemplo1[1] = "valor do indice 1"
	exemplo1[2] = "valor do indice 2"
	exemplo1[3] = "valor do indice 3"
	exemplo1[4] = "valor do indice 4"

	fmt.Println(exemplo1)    // irá imprimir todos os valores do array
	fmt.Println(exemplo1[1]) // irá imprimir apenas o valor da posição 1 `valor do indice 1`

	/*======================================================================================*/

	exemplo2 := [5]string{
		"valor do indice 0",
		"valor do indice 1",
		"valor do indice 2",
		"valor do indice 3",
		"valor do indice 4",
	}

	fmt.Println(exemplo2)    // irá imprimir todos os valores do array
	fmt.Println(exemplo2[1]) // irá imprimir apenas o valor da posição 1 `valor do indice 1`

	/*======================================================================================*/

	/*SLICE:
	É semelhante a um array, mas sua capacidade pode ser modificada durante a execução do programa,
	permitindo que mais elementos sejam adicionados ou removidos conforme necessário */

	var exemplo3 []string
	exemplo3 = append(exemplo3, "valor do indice 0") // é possivel adicionar 1 ou N valores
	exemplo3 = append(exemplo3, "valor do indice 1", "valor do indice 2", "valor do indice 3")

	fmt.Println(exemplo3)    // irá imprimir todos os valores do slice
	fmt.Println(exemplo3[1]) // irá imprimir apenas o valor da posição 1 `valor do indice 1`

	exemplo4 := []string{
		"valor do indice 0",
		"valor do indice 1",
	}
	exemplo4 = append(exemplo4, "valor do indice 2") // é possivel adicionar 1 ou N valores
	exemplo4 = append(exemplo4, "valor do indice 3", "valor do indice 4", "valor do indice 5")

	fmt.Println(exemplo3)    // irá imprimir todos os valores do slice
	fmt.Println(exemplo3[1]) // irá imprimir apenas o valor da posição 1 `valor do indice 1`

	exemplo5 := make([]string, 0)
	exemplo5[0] = "valor do indice 0"
	exemplo5[1] = "valor do indice 1"
	exemplo5[2] = "valor do indice 2"
	exemplo5[3] = "valor do indice 3"
	exemplo5[4] = "valor do indice 4"

	fmt.Println(exemplo5)    // irá imprimir todos os valores do slice
	fmt.Println(exemplo5[1]) // irá imprimir apenas o valor da posição 1 `valor do indice 1`

	/*======================================================================================*/

	/* MAP:
	o map é uma estrutura de dados que permite armazenar pares de chave-valor de forma dinâmica.
	O mapa é definido usando a sintaxe map[TipoChave]TipoValor, onde TipoChave é o tipo da chave e TipoValor é o tipo do valor associado a essa chave.
	Por exemplo, para criar um mapa que associa strings a inteiros, usamos a sintaxe map[string]int:
	*/

	mapExemplo1 := make(map[string]int)
	mapExemplo1["chave1"] = 1
	mapExemplo1["chave2"] = 2

	fmt.Println(mapExemplo1)           // irá imprimir todos os valores do Map
	fmt.Println(mapExemplo1["chave1"]) // irá imprimir apenas o valor da chave

	mapExemplo2 := map[string]int{
		"chave1": 1,
		"chave2": 2,
	}

	fmt.Println(mapExemplo2)           // irá imprimir todos os valores do Map
	fmt.Println(mapExemplo2["chave1"]) // irá imprimir apenas o valor da chave

	// Podemos acessar os valores do mapa fornecendo a chave correspondente entre colchetes:
	valorDaChave1 := mapExemplo2["chave1"]
	fmt.Println(valorDaChave1)

	/* Se a chave não existir no mapa, o valor retornado será o valor default do tipo de valor do mapa( No caso de inteiros, isso é 0)
	Podemos verificar se uma chave existe no mapa usando a sintaxe:
	 `valor, ok := m[chave]`
	onde `ok`` é uma variável booleana que indica se a chave existe no mapa:*/

	valor, ok := mapExemplo2["chave3"] // o nome da varíavel ok pode ser qualquer nome
	if ok {
		fmt.Println("O valor da chave3 é", valor)
	} else {
		fmt.Println("A chave3 não existe no mapa")
	}

}

/*
Funções são blocos de código que executam uma ação.
Elas são declaradas usando a palavra-chave "func" seguida do nome da função, uma lista de argumentos entre parênteses (se houver), e um tipo de retorno (se houver).
Veja um exemplo de uma função simples em Go que recebe 2 argumentos e retorna uma valor do tipo inteiro:
*/
func soma(a int, b int) int {
	return a + b
}

// Veja um exemplo de uma função que não retorna nenhum tipo de valor:
func imprimir(mensagem string) {
	fmt.Println(mensagem)
}

// Exemplo de uma função que possui um retorno nomeado.
func subtrair(a int, b int) (resultado int) {
	resultado = a - b
	return
}

// É possível ter várias declarações de retorno em uma única função.
func divide(a float64, b float64) (float64, string) {
	if b == 0 {
		return 0, "Não é possível dividir por zero"
	}
	return a / b, ""
}
