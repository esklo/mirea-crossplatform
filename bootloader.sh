GOOS=linux GOARCH=amd64 u-root  -defaultsh="" ./init
qemu-system-x86_64 -kernel vmlinuz-linux -initrd /tmp/initramfs.linux_amd64.cpio -append "console=ttyS0"