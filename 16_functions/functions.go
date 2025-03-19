/*

Uma função é um bloco de código que executa uma tarefa específica.
São usadas para dividir o código em partes menores e modulares.
Também para reutilizar o código e facilitar a leitura e manutenção do código.

Para criar (geralmente chamado de declarar) uma função, faça o seguinte:

Use a funcpalavra-chave.
Especifique um nome para a função, seguido de parênteses ().
Por fim, adicione o código que define o que a função deve fazer, entre chaves {}.

Regras de nomenclatura para funções Go
Um nome de função deve começar com uma letra
Um nome de função pode conter apenas caracteres alfanuméricos e sublinhados ( A-z, 0-9, e _ )
Os nomes das funções diferenciam maiúsculas de minúsculas
Um nome de função não pode conter espaços
Se o nome da função consistir em várias palavras, técnicas introduzidas para nomenclatura de variáveis ​​com várias palavras podem ser usadas

Informações podem ser passadas para funções como um parâmetro. Parâmetros agem como variáveis ​​dentro da função.
Sintaxe de uma função com parâmetros:
func nomeDaFuncao(parametro1 tipo, parametro2 tipo) {...}

O return é usado para retornar um valor de uma função.

*/

package main

import (
	"fmt"
)

func criarCarro(marca string) string {
	return "Carro criado: " + marca
}

func minhaFuncao(x int, y string) (resultado int, texto string) {
	resultado = x * 2
	texto = y + " Mundo!"
	return
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	fmt.Println(criarCarro("Golf"))
	fmt.Println(criarCarro("A3"))
	fmt.Println(criarCarro("320i"))

	dez, ola := minhaFuncao(10, "Olá")
	fmt.Println(dez, ola)

	fmt.Println("Fatorial de 5:", fatorial(5))

}
