# delivery- api

## O quê é?

Aplicativo que gerencia atividades de um serviço de pedidos em um restaurante. Desde a base de clientes, catálogo de produtos, pedidos e fila de preparo.


## Como executar


###  Docker 

Ao iniciar o projeto, basta executar o comando:

```
docker compose up
```

Ele irá inicializar a aplicação c


### Migrations




A aplicação conta com versionamento de migrações feita externamente. Para isso, utilizamos o pacote [golang-migrate](https://github.com/golang-migrate/migrate). Caso ainda não tenha instalado, você pode seguir a documentação ou executar o comando abaixo. Lembre-se que é necessário ter o `go` instalado em sua máquina.

```
make local-install
```

Para executar as migrations no banco, bastar estar com a instância do banco local rodando, e executar o comando abaixo.

```
make setup-dev
```

### Tests 

```
go test -v ./... --coverprofile=c.out && go tool cover -html=c.out
```

### Update Swagger

Instale essas ferramentas: https://github.com/swaggo/gin-swagger

````
swag init -g ./cmd/api/main.go -o ./cmd/api/docs --parseDependency  
```