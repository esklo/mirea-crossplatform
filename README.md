## TV PLot (Crossplatofrm)

## Build

`make`

#### linux

`./app`

#### windows

`./windows.exe`

#### web

`go run web/server.go`
> Open [127.0.0.1:8080](http://127.0.0.1:8080) in browser

## Bootloader

### Install [u-root](github.com/u-root/u-root)

`go install github.com/u-root/u-root@latest`

`chmod +x bootloader.sh`

`./bootloader.sh`
> ctrl + alt + 3 (serial port)

## Testing

`cd game && go test`
> *Since the program uses random, we need to get rid of it*
>
> change `X=INT(10*RND(1)+1)` to `X=0` in `tvplot.bas` file
>
> change `return rand.Intn(10) + 1` to `return 1` in `game/main.go` file
>
> Download [vintbas](http://www.vintage-basic.net/download.html) to /game dir
>
> Download [vmlinuz-linux](http://ftp.swin.edu.au/archlinux/iso/2023.09.01/arch/boot/x86_64/) to project root