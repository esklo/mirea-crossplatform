all: web/main.wasm build/linux build/windows.exe
web/main.wasm:
	GOOS=js GOARCH=wasm go build -o web/main.wasm
build/linux:
	go build -o build/linux
build/windows.exe:
	GOOS=windows GOARCH=amd64 go build -o build/windows.exe