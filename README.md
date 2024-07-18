## Descrição

Repositório da API feita em Golang (Venda de ingressos)

## Rodar a aplicação

Dentro da pasta `golang` execute o comando abaixo para rodar os containers `Docker`:

```
docker compose up
```

Quando os containers estiverem prontos, precisamos acessar o container do `golang` e executar a aplicação:

```
// entrar no container:
docker compose exec golang sh

// instalar as dependências:
go mod tidy

// executar a aplicação:
go run cmd/events/main.go
```
