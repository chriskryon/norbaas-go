# Estudo de Golang

Este projeto tem como objetivo estudar a linguagem de programação Go (Golang) utilizando como base o conteúdo disponível no [W3Schools Go Tutorial](https://www.w3schools.com/go/).

## Estrutura do Projeto

O projeto está organizado em diretórios numerados, cada um representando um tópico específico do estudo de Go.
Cada diretório contém um arquivo `.go` com exemplos e explicações sobre o tópico correspondente.

## Script `01_folders.sh`

O script `01_folders.sh` foi criado para automatizar a criação dos diretórios e arquivos `.go` necessários para o estudo. Ele realiza as seguintes ações:

1. Cria os diretórios numerados (ex.: `01_home`, `02_introduction`, etc.).
2. Cria os arquivos `.go` dentro de cada diretório, com nomes baseados no diretório correspondente.
3. Define permissões de leitura e escrita para os diretórios e arquivos criados.

### Como executar o script

1. Certifique-se de que você tem o Bash instalado no seu sistema operacional.
2. Navegue até o diretório do projeto no terminal.
3. Execute o comando:

   ```bash
   ./01_folders.sh
   ```

O script criará todos os diretórios e arquivos necessários para o estudo.

## Instalação do Go

Para instalar o Go, acesse o site oficial: [https://golang.org/dl/](https://golang.org/dl/).  
Baixe e siga as instruções de instalação para o seu sistema operacional.

Após a instalação, verifique se o Go foi instalado corretamente executando o comando:

```bash
go version
```

## Instalação do Visual Studio Code

Para instalar o Visual Studio Code, acesse o site oficial: [https://code.visualstudio.com/](https://code.visualstudio.com/).  
Baixe e siga as instruções de instalação para o seu sistema operacional.

Após instalar o VS Code, instale a extensão Go para facilitar o desenvolvimento:

1. Abra o VS Code.
2. Acesse o gerenciador de extensões (ou pressione `Ctrl + Shift + X`).
3. Pesquise por "Go" e instale a extensão oficial da equipe Go do Google.
4. Após a instalação, abra a paleta de comandos (`Ctrl + Shift + P`) e execute o comando `Go: Install/Update Tools`.
5. Selecione todas as ferramentas e clique em OK.

Agora você está pronto para começar a estudar Go!

## Como começar

1. Certifique-se de que o Go está instalado e configurado corretamente.
2. Execute o script `01_folders.sh` para criar a estrutura do projeto.
3. Navegue pelos diretórios e explore os exemplos de código em cada arquivo `.go`.
4. Compile e execute os exemplos utilizando o comando:

   ```bash
   go run <diretorio>/<nome_do_arquivo>.go
   ```

Bom estudo!