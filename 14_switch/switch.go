/*

O switch é uma maneira de avaliar uma expressão e executar diferentes blocos de código com base em valores correspondentes a essa expressão.

É como uma série de declarações if-else simplificadas.

A expressão é avaliada uma vez
O valor da switchexpressão é comparado com os valores de cadacase
Se houver uma correspondência, o bloco de código associado será executado
A default palavra-chave é opcional. Ela especifica algum código a ser executado se não houver case correspondência


Existe também o multi-case switch, onde você pode agrupar vários valores de case em um único case.


*/

package main

import (
	"fmt"
)

func main() {
	dia := 4

	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	}

	// Multi-case switch
	switch dia {
	case 1, 3, 5:
		fmt.Println("Dia útil ímpar")
	case 2, 4:
		fmt.Println("Dia útil par")
	case 6, 7:
		fmt.Println("Fim de semana")
	default:
		fmt.Println("Número de dia inválido")
	}
}
