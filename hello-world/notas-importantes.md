# Comandos importantes

## Comando para executar arquivos em Go

```shell
go run main.go
```

## Comando para gerar o executável Go

```shell
go build main.go
```

Caso seja necessário gerar um executável para outro tipo de sistema operacional, podemos fazer isso através desse comando:

```shell
GOOS=linux go build main.go
```

## Comando para instalar bibliotecas externas

```shell
go get github.com/google/uuid
```
