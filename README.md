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

- [Go 1.23+](https://go.dev/dl/)
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

```bash
# instale o evans para testar o gRPC (se ainda não tiver)
go install github.com/ktr0731/evans@latest

# com o evans instalado execute (depois de subir a aplicação)
evans -r repl
```

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
``` 
---
### 3 - Necessário atualizar as dependências:
```bash
go mod tidy
```
---

### 4 - Subir o Docker

> Certifique-se de estar na raiz do projeto.

```bash
sudo docker-compose up -d
```

### 5 Serviços disponiveis

| Serviço   | Porta  | Protocolo | Descrição                           |
|-----------|--------|-----------|-------------------------------------|
| REST      | 8000   | HTTP      | Endpoints REST para criação e listagem de orders |
| GraphQL   | 8080   | HTTP      | Playground e API GraphQL            |
| gRPC      | 50051  | gRPC      | Interface gRPC da aplicação         |
| MySQL     | 3306   | TCP       | Banco de dados relacional usado pela aplicação |


## Comandos úteis

```bash
docker-compose down             # Para parar os serviços
make migrate                    # para executar a migration
make migratedown                # para fazer rollback da migration
sudo docker exec -it mysql bash # Acessar o container do docker
mysql -u root -p                # Acessar o my sql dentro do docker
evans -r repl                   # Executar o Evans para tests
```
---