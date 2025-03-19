/*
O forloop percorre um bloco de código um número especificado de vezes.

Existe só uma forma de loop em Go, o forloop.
O forloop pode receber até três instruções:
- Inicialização: Executada antes do primeiro loop
- Condição: Avaliada antes de cada loop
- Pós: Executada no final de cada loop


O break é usado para sair de um loop.
O continue é usado para pular a iteração atual e ir para a próxima.

A rangepalavra-chave é usada para iterar mais facilmente pelos elementos de um array, slice ou map. Ela retorna tanto o índice quanto o valor.

A rangepalavra-chave é usada assim:
for index, value := range array|slice|map {
   // code to be executed for each iteration
}

*/

package main

import "fmt"

func main() {
	// Loop básico
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Loop estilo `while`
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Loop infinito
	for {
		fmt.Println("Loop infinito")
		break // Use `break` para sair do loop
	}

	// Iterando sobre coleções
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Ignorando valores com `_`
	for _, value := range nums {
		fmt.Println(value)
	}

	// Saindo de loops
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Pula números pares
		}
		if i > 7 {
			break // Sai do loop se i > 7
		}
		fmt.Println(i)
	}

	// Exemplo adicional de uso do `continue`
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue // Pula números divisíveis por 3
		}
		fmt.Println(i)
	}
}
