/*
Existem três funções para output em Go: Print, Println e Printf.

A Print() exibe o texto na tela sem pular linha.
Println() exibe o texto na tela e pula linha.
Printf() exibe o texto formatado na tela.

A função Printf() primeiro formata seu argumento com base no verbo de formatação fornecido e, em seguida, os imprime.

Tipos de formatação:
	%v  - valor dos argumentos
	%#v - valor no formato Go-syntax
	%T  - tipo dos argumentos
	%%  - símbolo de porcentagem

*/

package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	fmt.Print(i, "\n", j)
	fmt.Print(i, " ", j)
	fmt.Print(i, j)

	// Println exibe o texto na tela e pula linha após cada chamada
	fmt.Println(i, j)

	// Printf exibe o texto formatado na tela
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)

	var ii = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", ii)
	fmt.Printf("%#v\n", ii)
	fmt.Printf("%v%%\n", ii)
	fmt.Printf("%T\n", ii)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

}
