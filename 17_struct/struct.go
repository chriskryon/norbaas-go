/*
Um struct é usado para agrupar membros de diferentes tipos de dados em uma única variável.

Enquanto arrays armazenam múltiplos valores do mesmo tipo, structs armazenam múltiplos valores de tipos diferentes.

Structs são úteis para agrupar dados e criar registros.

Em TypeScript, um struct pode ser representado por uma interface ou uma classe.
interface Pessoa {
	nome: string;
	idade: number;
	ativo: boolean;
}
const pessoa: Pessoa = {
	nome: "João",
	idade: 30,
	ativo: true,
};

*/

package main

import (
	"fmt"
)

type Carro struct {
	modelo string
	ano    int
	marca  string
	preco  int
}

func main() {
	var carro1 Carro
	var carro2 Carro

	// Especificação do Carro1
	carro1.modelo = "Golf GTI MK7"
	carro1.ano = 2017
	carro1.marca = "Volkswagen"
	carro1.preco = 30000

	// Especificação do Carro2
	carro2.modelo = "Audi A3 2.0"
	carro2.ano = 2018
	carro2.marca = "Audi"
	carro2.preco = 35000

	// Acessar e imprimir informações do Carro1
	fmt.Println("Modelo: ", carro1.modelo)
	fmt.Println("Ano: ", carro1.ano)
	fmt.Println("Marca: ", carro1.marca)
	fmt.Println("Preço: ", carro1.preco)

	// Acessar e imprimir informações do Carro2
	fmt.Println("Modelo: ", carro2.modelo)
	fmt.Println("Ano: ", carro2.ano)
	fmt.Println("Marca: ", carro2.marca)
	fmt.Println("Preço: ", carro2.preco)
}
