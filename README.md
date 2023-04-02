# Clean architecture com golang
Implementando Clean Architecture com Golang

# Database
Instale a CLI migrate para gerar os arquivos de migrations necessários para o projeto.

# Gerar o arquivo de migrations
migrate create -ext sql -dir database/migrations -seq create_product_table

Edite o arquivo gerado em database/migrations/000001.create_product_table.up.sql com o SQL da criação da tabela product.

CREATE TABLE product (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  price FLOAT NOT NULL,
  description VARCHAR(500) NOT NULL
);

E também altere o arquivo database/migrations/000001.create_product_table.down.sql.

DROP TABLE IF EXISTS product;

# Go modules
Vamos inicializar os módulos do go com o comando:
    go mod init github.com/booscaaa/clean-go

go run adapter/http/main.go