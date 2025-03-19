/*

Os operadores condicionais são usados para comparar valores e tomar decisões com base no resultado dessas comparações.

Os operadores condicionais em Go são:
Menor que (<)
Maior que (>)
Menor ou igual a (<=)
Maior ou igual a (>=)
Igual a (==)
Different de (!=)

Além disso, existem operadores lógicos que podem ser usados para combinar várias expressões condicionais. Os operadores lógicos em Go são:
E lógico (&&)
Ou lógico (||)
Não lógico (!)

Já as estruturas condicionais em Go são:
if = se
else = senão
else if = senão se
switch = selecione

*/

package main

import "fmt"

func main() {
	// Exemplo de uso de operadores condicionais e estruturas condicionais

	idade := 20

	// Verifica se a idade é maior ou igual a 18
	if idade >= 18 {
		// Se a condição for verdadeira, executa este bloco
		fmt.Println("Você é maior de idade.")
	} else {
		// Caso contrário, executa este bloco
		fmt.Println("Você é menor de idade.")
	}

	// Exemplo de uso de else if
	nota := 85

	if nota >= 90 {
		// Se a nota for maior ou igual a 90
		fmt.Println("Parabéns! Você tirou um A.")
	} else if nota >= 70 {
		// Se a nota for maior ou igual a 70, mas menor que 90
		fmt.Println("Bom trabalho! Você tirou um B.")
	} else {
		// Se a nota for menor que 70
		fmt.Println("Você precisa estudar mais.")
	}

	// Exemplo de uso de operadores lógicos
	temperatura := 30
	umidade := 70

	// Verifica se a temperatura é maior que 25 e a umidade é maior que 60
	if temperatura > 25 && umidade > 60 {
		fmt.Println("Está quente e úmido.")
	}

	// Verifica se a temperatura é menor que 15 ou a umidade é menor que 40
	if temperatura < 15 || umidade < 40 {
		fmt.Println("Está frio ou seco.")
	}

	// Verifica se não está quente
	if !(temperatura > 25) {
		fmt.Println("Não está quente.")
	}
}
