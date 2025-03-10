/*
Variáveis são como gaveas que armazenam valores.
Elas são declaradas com a palavra-chave var.
A sintaxe é var nomeDaVariavel tipoDaVariavel = valorDaVariavel
Ou a declaração por inferência de tipo: var nomeDaVariavel = valorDaVariavel
Ou também declaração por escopo: nomeDaVariavel := valorDaVariavel
O tipo da variável é opcional, mas é uma boa prática declarar o tipo da variável.

Tipos de variáveis:
int - armazena inteiros (números inteiros), como 123 ou -123
float32 - armazena números de ponto flutuante, com decimais, como 19.99 ou -19.99
string - armazena texto, como "Hello World". Valores de string são cercados por aspas duplas
bool - armazena valores com dois estados: true ou false

Regras de nomenclatura de variáveis:
- O nome da variável deve começar com uma letra ou um sublinhado
- O nome da variável não pode começar com um número
- O nome da variável não pode conter espaços
- O nome da variável não pode conter caracteres especiais, exceto sublinhados
- O nome da variável é sensível a maiúsculas e minúsculas
- O nome da variável deve ser único no escopo em que é declarado
- O nome da variável deve ser descritivo, para que você e outros possam entender o que a variável contém
- O nome da variável não pode ser uma palavra-chave reservada, como var, func, etc.

Nomes de variáveis com mais de uma palavra podem ser difíceis de ler.
Existem várias técnicas que você pode usar para torná-los mais legíveis:
CamelCase: A primeira letra da primeira palavra é minúscula e a primeira letra de cada palavra subsequente é maiúscula. Exemplo: minhaVariavelExemplo
PascalCase: A primeira letra de cada palavra é maiúscula. Exemplo: MinhaVariavelExemplo
snake_case: Todas as letras são minúsculas e as palavras são separadas por sublinhados. Exemplo: minha_variavel_exemplo

No Go, a convenção mais usada é o CamelCase para nomes de variáveis e PascalCase para nomes de tipos exportados (como structs e interfaces).
*/

package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" // tipo é string
	var student2 = "Jane"        // tipo é inferido
	x := 2                       // tipo é inferido

	// E possível declarar sem usar o valor
	var y int
	var z string
	var falso bool

	// É possível atribuir valores a variáveis já declaradas
	y = 3

	// A diferença entre var e := é que:
	// var pode ser usado dentro e fora de funções
	// := só pode ser usado dentro de funções
	// A declaração de variáveis e a atribuição de valores podem ser feitas separadamente com var
	// A declaração de variáveis e a atribuição de valores não podem ser feitas separadamente com := (devem ser feitas na mesma linha)

	// Declaração de múltiplas variáveis
	var a, b, c, d int = 1, 3, 5, 7

	// Declarações de múltiplas variáveis ​​também podem ser agrupadas em um bloco para maior legibilidade:
	var (
		x1 = 1
		y1 = 2
		z1 = 3
	)

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(y)     // 3
	fmt.Println(z)     // ""
	fmt.Println(falso) // false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)
}
