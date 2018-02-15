build:
	go build -i -o main -v

fmt:
	go fmt
	cd ./server && go fmt
	cd ./client && go fmt
	cd ./core && go fmt

.PHONY: fmt build

