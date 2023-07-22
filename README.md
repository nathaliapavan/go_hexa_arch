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
