build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/app/main.go

run:
	go run cmd/server/main.go