## TV PLot (Crossplatofrm)

## Build

`make`

#### linux

`./build/linux`

#### windows

`./build/windows.exe`

#### web

`go run web/server.go`
> Open [127.0.0.1:8080](http://127.0.0.1:8080) in browser

## Bootloader

> For bootloader I use qemu with linux kernel. program compiles into initramfs with u-root

### Install [u-root](https://github.com/u-root/u-root)

### Install [QEMU](https://www.qemu.org/download/)

`go install github.com/u-root/u-root@latest`

`chmod +x bootloader.sh`

`./bootloader.sh`
> ctrl + alt + 3 (serial port)

## Testing

### Download [vintbas](http://www.vintage-basic.net/download.html) to /game dir

### Download [vmlinuz-linux](http://ftp.swin.edu.au/archlinux/iso/2023.09.01/arch/boot/x86_64/) to project root

`cd game && go test`
> *Since the program uses random, we need to get rid of it*
>
> change `X=INT(10*RND(1)+1)` to `X=0` in `tvplot.bas` file
>
> change `return rand.Intn(10) + 1` to `return 1` in `game/main.go` file

## Structure

```
.
├── Makefile
├── README.md
├── bootloader.sh
├── build
│   ├── linux
│   └── windows.exe
├── exec -> start game proccess
│   ├── cli.go -> compiles if not web
│   ├── exec.go -> entrypoint
│   └── web.go -> compiles only for web
├── game -> gameport
│   ├── main.go
│   ├── main_test.go
│   ├── tvplot.bas
│   └── vintbas
├── go.mod
├── go.sum
├── init -> entrypoint for bootloader (init – required name for u-root)
│   └── init.go
├── main.go
├── vmlinuz-linux
└── web
    ├── index.html
    ├── main.wasm
    ├── server.go
    └── wasm_exec.js
```

