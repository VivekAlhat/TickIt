greet:
	echo "Hello from Tickit"

test:
	cd internal && go test

build:
	go build -o bin/tickit main.go

compile:
	echo "Compiling for Mac, Linux and Windows platform"
	GOOS=linux GOARCH=amd64 go build -o bin/tickit-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/tickit-mac main.go
	GOOS=windows GOARCH=amd64 go build -o bin/tickit-windows.exe main.go

run:
	go run main.go