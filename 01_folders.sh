#!/bin/bash

# Lista de pastas numeradas
pastas=(
  "01_home"
  "02_tutorial"
  "03_introduction"
  "04_get_started"
  "05_syntax"
  "06_comments"
  "07_variables"
  "08_constants"
  "09_output"
  "10_data_types"
  "11_arrays"
  "12_slices"
  "13_operators"
  "14_conditions"
  "15_switch"
  "16_loops"
  "17_functions"
  "18_struct"
  "19_maps"
)

# Loop para criar pastas e arquivos .go
for pasta in "${pastas[@]}"; do
  # Remove a numeração e o underscore do nome do arquivo .go
  arquivo_go="${pasta}/${pasta#*_}.go"
  
  # Cria a pasta (se não existir)
  mkdir -p "$pasta"
  
  # Cria o arquivo .go vazio
  touch "$arquivo_go"
  
  # Define permissões para a pasta (rwx para o usuário)
  chmod a+rw  "$pasta"
  
  # Define permissões para o arquivo .go (rw para o usuário)
  chmod a+rw  "$arquivo_go"
  
  echo "Arquivo criado: ${arquivo_go}"
done

echo "Todas as pastas e arquivos .go foram criados com sucesso!"