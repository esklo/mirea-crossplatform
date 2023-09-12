all: web/main.wasm app windows.exe
web/main.wasm:
	GOOS=js GOARCH=wasm go build -o web/main.wasm
app:
	go build -o app
windows.exe:
	GOOS=windows GOARCH=amd64 go build -o windows.exe