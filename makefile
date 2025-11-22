APP := treecraft

.PHONY: build run clean test fmt vet install tool get

## Build the binary (local)
build:
	go build -o $(APP) ./cmd/treecraft

## Run the app locally
run: build
	./$(APP)

## Format code
fmt:
	go fmt ./...

## Lint
vet:
	go vet ./...

## Run tests
test:
	go test ./...

## Install YOUR cli globally (put in $HOME/go/bin)
install:
	go install ./cmd/treecraft

## Install external tools (air, goreleaser, etc.)
## Usage: make tool pkg=github.com/cosmtrek/air
tool:
ifndef pkg
	$(error You must provide a package. Example: make tool pkg=github.com/cosmtrek/air)
endif
	go install $(pkg)

## Add a dependency
## Usage: make get pkg=github.com/google/uuid
get:
ifndef pkg
	$(error You must provide a package name. Example: make get pkg=github.com/google/uuid)
endif
	go get $(pkg)

## Remove local binary
clean:
	rm -f $(APP)
