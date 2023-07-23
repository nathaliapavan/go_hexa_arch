# go_hexa_arch

### para subir o container

```bash
docker-compose up -d
```

### para acessar o container e executar bash

```bash
docker exec -it CONTAINER_ID bash
```

### para fazer com que o go baixe as dependencias automaticamente

```bash
go mod init github.com/nathaliapavan/go_hexa_arch
```

### para baixar pacote em uma versão especifica

```bash
go get github.com/stretchr/testify@v1.7.0
```

### para executar os testes (dentro do container)

```bash
go test ./...
```

### para gerar mocks de product.go

```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application/
```

### para instalar o sqlite3

```bash
docker exec -it -u 0 7dc4206ab790 bash # consegui sudo assim
apt-get update
apt-get install sqlite3
```

### para criar o banco sqlite3

```bash
touch db.sqlite
```

### para acessar o banco sqlite3

```bash
sqlite3 db.sqlite
```

### para criar uma tabela no banco sqlite3

```sql
create table products(id string, name string, price float, status string);
```

### para listar tabelas no banco sqlite3

```bash
.tables
```
