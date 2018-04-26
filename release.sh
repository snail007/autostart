#!/bin/bash
rm -rf autostart goautostart autostart.exe goautostart.exe *.tar.gz

#linux
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o autostart -ldflags "-s -w" && tar zcfv "autostart-linux-386.tar.gz" autostart
rm -rf autostart
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o autostart -ldflags "-s -w" && tar zcfv "autostart-linux-amd64.tar.gz" autostart
rm -rf autostart
#darwin
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o autostart -ldflags "-s -w" && tar zcfv "autostart-darwin-386.tar.gz" autostart
rm -rf autostart
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o autostart -ldflags "-s -w" && tar zcfv "autostart-darwin-amd64.tar.gz" autostart
rm -rf autostart
#windows
#apt-get install gcc-multilib 
#apt-get install gcc-mingw-w64
CC=i686-w64-mingw32-gcc-win32 CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o autostart.exe && tar zcfv "autostart-windows-386.tar.gz" autostart.exe
rm -rf autostart.exe
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o autostart.exe && tar zcfv "autostart-windows-amd64.tar.gz" autostart.exe
rm -rf autostart.exe