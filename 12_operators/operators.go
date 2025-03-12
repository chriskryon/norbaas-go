/*
Os operadores em Go são utilizados para realizar operações em operandos e valores. Go suporta os seguintes tipos de operadores:

Operadores Aritméticos:
+  serve para adição
-  serve para subtração
*  serve para multiplicação
/  serve para divisão
%  serve para módulo
++ serve para incremento
-- serve para decremento

Operadores de Atribuição:
=  serve para atribuição

Operadores de Comparação:
== serve para igualdade
!= serve para desigualdade
>  serve para maior que
<  serve para menor que
>= serve para maior ou igual que
<= serve para menor ou igual que

Operadores Lógicos:
&& serve para operador lógico E
|| serve para operador lógico OU
!  serve para operador lógico NÃO

Operadores Bitwise:
&  serve para AND bit a bit
|  serve para OR bit a bit
^  serve para XOR bit a bit
<< serve para deslocamento à esquerda
>> serve para deslocamento à direita
*/

package main

import (
	"fmt"
)

func main() {
	// Operadores Aritméticos
	a, b := 10, 5
	fmt.Println("Adição:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Módulo:", a%b)
	a++
	fmt.Println("Incremento:", a)
	a--
	fmt.Println("Decremento:", a)

	// Operadores de Atribuição
	c := 10
	fmt.Println("Atribuição:", c)

	// Operadores de Comparação
	fmt.Println("Igualdade:", a == b)
	fmt.Println("Desigualdade:", a != b)
	fmt.Println("Maior que:", a > b)
	fmt.Println("Menor que:", a < b)
	fmt.Println("Maior ou igual que:", a >= b)
	fmt.Println("Menor ou igual que:", a <= b)

	// Operadores Lógicos
	d, e := true, false
	fmt.Println("E lógico:", d && e)
	fmt.Println("OU lógico:", d || e)
	fmt.Println("NÃO lógico:", !d)

	// Operadores Bitwise
	f, g := 10, 2
	fmt.Println("AND bit a bit:", f&g)
	fmt.Println("OR bit a bit:", f|g)
	fmt.Println("XOR bit a bit:", f^g)
	fmt.Println("Deslocamento à esquerda:", f<<g)
	fmt.Println("Deslocamento à direita:", f>>g)
}
