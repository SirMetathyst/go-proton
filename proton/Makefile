LDFLAGS="-s -w"
INPUT="./main.go"
OUTPUT="./proton.exe"

build:
	packr.exe build -ldflags=$(LDFLAGS) -o $(OUTPUT) $(INPUT)

buildfinal: build
	upx.exe --brute $(OUTPUT)

install: buildfinal
	cp ./proton.exe ../../../../../bin/proton.exe

install_clean: buildfinal
	rm ./proton.exe