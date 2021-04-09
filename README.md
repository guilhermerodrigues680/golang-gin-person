
## Instalar CLI PKGER
```sh
$ go get github.com/markbates/pkger/cmd/pkger@v0.17.1
```

## Examina o código e lista os arquivos que devem ser empacotados

```sh
$ pkger list
```

## empacota os arquivos
```sh
$ pkger -o ./cmd/app
```

## Deploy Binário
```sh
$ PORT=8080 GIN_MODE=release ./bin/main-linux-amd64
```