build:
	go build -ldflags="-s -w" -trimpath ./cmd/server/server.go 
	set GOARCH=wasm
	set GOOS=js
	go build -ldflags="-s -w" -trimpath -o ./static/wasm/main.wasm ./cmd/wasm/wasm.go 

debug: build
	.\server.exe