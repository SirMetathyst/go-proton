LDFLAGS="-s -w"
INPUT="./main.go"
OUTPUT="./proton.exe"

build:
	go.exe build -ldflags=$(LDFLAGS) -o $(OUTPUT) $(INPUT)

buildadv:
	packr.exe build -ldflags=$(LDFLAGS) -o $(OUTPUT) $(INPUT)

buildadvfinal: build
	upx.exe --brute $(OUTPUT)

install: buildadvfinal
	cp ./proton.exe /mnt/c/Users/alexa/go/bin/proton.exe

install_clean: install
	rm ./proton.exe
