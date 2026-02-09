BINDIR := bin
GOBIN := $(BINDIR)/$(SERVER)
GITIGNORE := .gitignore

.PHONY: all fcserver clean

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

test: lint fcmul_test

lint:
	golangci-lint run

clean:
	rm -rf $(BINDIR)
