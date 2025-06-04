GOCMD 	:= go
GOBUILD := $(GOCMD) build
GORUN 	:= $(GOCMD) run
GOCLEAN := $(GOCMD) clean
GOTIDY  := $(GOCMD) mod tidy
BINARY  := bin/worker-pool

run: build start clean

build:
	$(GOBUILD) -o $(BINARY) -v ./cmd/app

start:
	$(BINARY)

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

update:
	$(GOTIDY)

.PHONY: run build start clean test update

