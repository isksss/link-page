build:
	go build ./cmd/server/server.go
	GOOS=js GOARCH=wasm go build -o main.wasm ./cmd/wasm/wasm.go