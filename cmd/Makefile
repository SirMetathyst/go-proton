LDFLAGS="-s -w"
INPUT="./main.go"
OUTPUT="./proton.exe"

build:
	packr.exe build -ldflags=$(LDFLAGS) -o $(OUTPUT) $(INPUT)
	upx.exe --brute $(OUTPUT)

install: build
	cp ./proton.exe ../../../../../bin/proton.exe

install_clean: install
	rm ./proton.exe