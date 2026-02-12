BINDIR := bin
GOBIN := $(BINDIR)/$(SERVER)
GITIGNORE := .gitignore

.PHONY: all fcserver fcclient clean test lint fcmul_test

all: lint fcserver fcclient

$(BINDIR):
	mkdir -p $(BINDIR)

fcserver: $(BINDIR)
	go build -o $(GOBIN) ./cmd/fcserver

fcclient: $(BINDIR)
	go build -o $(GOBIN) ./cmd/fcclient

fcmul_test:
	go test ./pkg/fcmul/lexer
	go test ./pkg/fcmul/parser
	go test ./pkg/fcmul

test: fcmul_test

bench: fcdb_bench

examples: fcdb_examples

lint:
	golangci-lint run

clean:
	rm -rf $(BINDIR)
