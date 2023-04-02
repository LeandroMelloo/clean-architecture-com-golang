# Clean architecture com golang
Implementando Clean Architecture com Golang

# Database
Instale a CLI migrate para gerar os arquivos de migrations necessários para o projeto.

# Gerar o arquivo de migrations
migrate create -ext sql -dir database/migrations -seq create_product_table

Edite o arquivo gerado em database/migrations/000001.create_product_table.up.sql com o SQL da criação da tabela product.
