/*

Arrays em GO são usados para armazenar vários valores do mesmo tipo, em uma única variável, em vez de declarar variáveis separadas para cada valor

Podemos declarar um array de duas maneiras:
1. var nome [tamanho]tipo
2. nome := [tamanho]tipo{valores}

Para acessar um elemento de um array, usamos o índice do elemento. O índice do primeiro elemento é 0, o segundo elemento é 1 e assim por diante.
prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])


Um array pode ser não inicializado, parcialmente inicializado ou totalmente inicializado.
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

Podemos também inicializar apenas alguns elementos de um array, e o restante dos elementos será inicializado com zero.
	  arr1 := [5]int{1:10,2:40}


*/

package main

import (
	"fmt"
)

func main() {
	// Este exemplo declara duas matrizes (arr1 e arr2) com comprimentos definidos:
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	// Este exemplo declara duas matrizes (arr3 e arr4) com comprimentos inferidos:
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	// Array de strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println(cars)
	// Acessando um elemento de um array
	fmt.Println(cars[1])

	// Alterando um elemento de um array
	cars[3] = "Audi"
	fmt.Println(cars[3])

	arr5 := [5]int{}              //não inicializado
	arr6 := [5]int{1, 2}          //parcialmente inicializado
	arr7 := [5]int{1, 2, 3, 4, 5} //totalmente inicializado

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	// Inicializando apenas alguns elementos de um array
	arr8 := [5]int{0: 10, 4: 40}
	fmt.Println(arr8)

	// Tamanho de um array
	fmt.Println(len(arr8))

}
