#!/bin/bash

cd /www/web

echo "Aplicando permissões de escrita na pasta STORAGE/"
chmod -R 777 storage

echo "Verificando dependências:"
go mod tidy  

echo "Executando testes:"
go test ./... 
 
echo "Iniciando a Aplicação..."
go run public/index.go