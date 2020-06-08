# .PHONY, where we define all the targets that are not files. 
#make will run its recipe regardless of whether a file with that name exists or what its last modification time is.
.PHONY: all run build test fmt

# Default take for make without argumets
.DEFAULT_GOAL := all

# Run all listed targets
all: fmt build test

run: 
	go run main.go get ns --contexts=docker-desktop 
build:
	go build -v
test:
	go test -v ./...
fmt:
	go fmt 
	go fmt github.com/daftping/kubeagg/pkg/kubeagg
	go fmt github.com/daftping/kubeagg/cmd