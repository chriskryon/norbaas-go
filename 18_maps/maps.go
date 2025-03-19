/*
Mmaps e structs são duas estruturas de dados diferentes, cada uma com suas próprias características e casos de uso. Vamos explorar as diferenças entre elas:

Maps
Definição: Um map é uma coleção de pares chave-valor, onde cada chave é única e mapeia para um valor. É semelhante a um dicionário em outras linguagens.

Tipos: Um map é homogêneo em termos de tipos, ou seja, todas as chaves devem ser do mesmo tipo, e todos os valores devem ser do mesmo tipo. Por exemplo, um map[string]int tem chaves do tipo string e valores do tipo int.

Mutabilidade: Maps são mutáveis, ou seja, você pode adicionar, remover e modificar pares chave-valor após a criação.

Uso comum: Maps são usados quando você precisa de uma coleção dinâmica de dados onde a chave é usada para acessar o valor correspondente.

A diferença de map para struct é que map é uma coleção de pares chave:valor, enquanto struct é uma coleção de membros de diferentes tipos de dados.

carro := map[string]string{
	"marca": "Toyota",
	"modelo": "Corolla",
	"ano":    "2020",
}
fmt.Println(carro)

Tipos de Chave Permitidos
A chave de um map pode ser de qualquer tipo de dado para o qual o operador de igualdade (==) esteja definido. Isso inclui:

Booleans (booleanos)
Numbers (números)
Strings (strings)
Arrays (arrays)
Pointers (ponteiros)
Structs (estruturas)
Interfaces (desde que o tipo dinâmico suporte igualdade)
Tipos de chave inválidos são:

Slices (fatias)
Maps (maps)
Functions (funções)

*/

package main

import (
	"fmt"
)

func main() {
	// Criar um map com chave do tipo string e valor do tipo string
	carro := map[string]string{
		"marca":  "Toyota",
		"modelo": "Corolla",
		"ano":    "2020",
	}

	m := make(map[string]int)
	m["chave1"] = 10
	m["chave2"] = 20
	fmt.Println(m["chave1"]) // Saída: 10

	// Imprimir o map
	fmt.Println(carro)
	fmt.Println(m)

	// Criando um map onde as chaves são strings e os valores são inteiros
	idades := make(map[string]int)

	// Adicionando valores ao map
	idades["João"] = 30
	idades["Maria"] = 25
	idades["Carlos"] = 40

	// Acessando valores do map
	fmt.Println("Idade de João:", idades["João"])
	fmt.Println("Idade de Maria:", idades["Maria"])

	// Verificando se uma chave existe no map
	idade, existe := idades["Ana"]

	if existe {
		fmt.Println("Idade de Ana:", idade)
	} else {
		fmt.Println("Ana não está no map.")
	}

	// Modificando um valor no map
	idades["Carlos"] = 41
	fmt.Println("Nova idade de Carlos:", idades["Carlos"])

	// Removendo um valor do map
	delete(idades, "Maria")
	fmt.Println("Maria foi removida do map.")

	// Iterando sobre o map
	fmt.Println("\nLista de idades:")
	for nome, idade := range idades {
		fmt.Printf("%s: %d anos\n", nome, idade)
	}

	// Iterar sobre mapas
	i := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range i {
		fmt.Printf("%v : %v, ", k, v)
	}
	fmt.Println()

	// Iterar sobre mapas em uma ordem específica
	var j []string
	j = append(j, "one", "two", "three", "four")

	for _, element := range j {
		fmt.Printf("%v : %v, ", element, i[element])
	}
}
