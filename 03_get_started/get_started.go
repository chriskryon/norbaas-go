/*

Go Get Started:
Para começar a utilizar o Go, precisamos de duas coisas:
- Um editor de texto ou IDE para escrever o código, como o Visual Studio Code, Atom, Sublime Text, etc.
- Compilador como o GCC, para traduzir o código Go em código de máquina.

Instalação:
Para instalar o Go, acesse o site oficial: https://golang.org/dl/
Siga as instruções de instalação para o seu sistema operacional.

Após a instalação, abra o terminal e digite o comando `go version` para verificar se a instalação foi bem-sucedida.

Instalação de uma IDE:
Sugiro a instalação do Visual Studio Code, pois é uma IDE leve e fácil de usar.
Para instalar o Visual Studio Code, acesse o site oficial: https://code.visualstudio.com/
Siga as instruções de instalação para o seu sistema operacional.

Abra o editor VS Code
Abrir o gerenciador de extensões ou, alternativamente, pressionar Ctrl + Shift + x
Na caixa de pesquisa, digitar "go" e pressionar Enter
Encontrar a extensão Go da equipe GO do Google e instale a extensão
Após a instalação, abra a paleta de comandos pressionando Ctrl + Shift + p
Executar o comando Go: Install/Update Tools
Selecionar todas as ferramentas fornecidas e clicar em OK

Agora, você está pronto para começar a escrever código Go.
Abra o terminal e digite go mod init example.com/hello para criar um módulo Go.

Um módulo Go é uma coleção de pacotes Go que são versionados juntos como uma unidade. Ele é definido por um arquivo go.mod, que especifica o nome do módulo e as dependências necessárias.

O que faz go mod init?
O comando go mod init cria um novo arquivo go.mod no diretório atual. Este arquivo contém informações sobre o módulo, incluindo seu nome e as dependências que ele usa.
Seria como se fosse um npm init.

*/

package main

import (
	"fmt"
)

func get_started() {
	fmt.Println("Para começar a utilizar o Go, precisamos de um editor de texto ou IDE e um compilador.")
	fmt.Println("Instale o Go do site oficial: https://golang.org/dl/")
	fmt.Println("Verifique a instalação com `go version`.")
	fmt.Println("Sugiro a instalação do Visual Studio Code como IDE.")
	fmt.Println("Instale a extensão Go no VS Code e atualize as ferramentas Go.")
	fmt.Println("Crie um módulo Go com `go mod init example.com/hello`.")
	fmt.Println("O comando `go mod init` cria um arquivo go.mod com informações sobre o módulo.")
}
