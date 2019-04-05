test:
	go.exe test -coverprofile coverage.out ./...
	go.exe tool cover -html=coverage.out -o coverage.html
	go.exe test -v