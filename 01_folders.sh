#!/bin/bash

# Lista de pastas numeradas
pastas=(
  "01_home"
  "02_introduction"
  "03_get_started"
  "04_syntax"
  "05_comments"
  "06_variables"
  "07_constants"
  "08_output"
  "09_data_types"
  "10_arrays"
  "11_slices"
  "12_operators"
  "13_conditions"
  "14_switch"
  "15_loops"
  "16_functions"
  "17_struct"
  "18_maps"
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