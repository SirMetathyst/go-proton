build:
	packr.exe build -o proton.exe ./**.go 
install: build
	cp ./proton.exe ../../../../../bin/proton.exe
install_clean: install
	rm ./proton.exe