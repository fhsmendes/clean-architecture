# clean-architecture
## Desafio fullcycle - Clean Architecture

Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

---

## Pré-requisitos

- [Go 1.18+](https://go.dev/dl/)
- [Git](https://git-scm.com/)
- [Docker Desktop (com WSL2 se estiver no Windows)](https://www.docker.com/products/docker-desktop)
- [VSCode](https://code.visualstudio.com/) com extensão `REST Client` (opcional para testar `.http`)

---

## Como executar

### 1 - Realizar o clone do projeto
---

### 2 - Pré requisitos
---

#### gRPC 

```bash
# instale o protoc e o plugin do Go (se ainda não tiver)

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
---

#### Wire
```bash
# Instale o `google wire`:

go install github.com/google/wire/cmd/wire@latest
```
#### Migrate 

```bash
# Instale o `golang-migrate`:

brew install golang-migrate  # macOS
# ou
sudo apt install -y migrate  # Ubuntu/Debian (pode requerer repo externo)
# ou binário:
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin
``` 

### 3- Necessário atualizar as dependências:
```bash
go mod tidy
```
---

### 4 - Subir o Docker

> Certifique-se de estar na raiz do projeto.

```bash
docker-compose up -d
```

Este comando irá iniciar:

- MySQL (porta `3306`)
---

### 5 - Executar as migrations (após o MySQL subir)

```bash
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" up
```

> Substitua o nome do banco se necessário.

---

### 6 - Rodar a aplicação

```bash
go run cmd/ordersystem/main.go wire_gen.go
```

---

##  Testar GraphQL (via arquivo `.http`)

Crie um arquivo `graphql.http`:

```http
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "{ healthCheck }"
}
```

Clique em **"Send Request"** (com a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)).

---

## Comandos úteis

```bash
docker-compose down             # Para parar os serviços
migrate -version                # Ver versão da migration atual
migrate -rollback               # Rollback da última versão
sudo docker exec -it mysql bash # Acessar o container do docker
mysql -u root -p                # Acessar o my sql dentro do docker
```
---

evans -r repl

https://github.com/ktr0731/evans

protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

OrdersList