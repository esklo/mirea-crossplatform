all: web/main.wasm linux windows.exe
web/main.wasm:
	GOOS=js GOARCH=wasm go build -o web/main.wasm
linux:
	go build -o build/linux
windows.exe:
	GOOS=windows GOARCH=amd64 go build -o build/windows.exe