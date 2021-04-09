build: cmd/app/main.go
	pkger list
	pkger -o ./cmd/app
	GOOS=linux GOARCH=amd64 go build -v -o bin/main-linux-amd64 ./cmd/app

cross: cmd/app/main.go
	pkger list
	pkger -o ./cmd/app
	go build -v -o bin/main ./cmd/app
