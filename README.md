## Linux, Windows, Web

`make`

```
./app
./windows.exe
open web/index.html
```

> for web use webserver

## Bootloader

[linux kernel](http://ftp.swin.edu.au/archlinux/iso/2023.09.01/arch/boot/x86_64/)

`chmod +x bootloader.sh`

`./bootloader.sh`
> ctrl + alt + 3 (serial port)

## Testing

`cd game && go test`
> change `X=INT(10*RND(1)+1)` to `X=0` in `.bas` file
>
> change `return rand.Intn(10) + 1` to `return 1` in `game/main.go` file