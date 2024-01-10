greet:
	echo "Hello from Tickit"

test:
	cd internal && go test

build:
	go build -o bin/tickit main.go

run:
	go run main.go