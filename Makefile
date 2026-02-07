BINDIR := bin
GOBIN := $(BINDIR)/$(SERVER)
GITIGNORE := .gitignore

.PHONY: all fcserver clean

all: fcserver

$(BINDIR):
	mkdir -p $(BINDIR)

fcserver: $(BINDIR)
	go build -o $(GOBIN) ./cmd/fcserver

fcmul_test:
	go test ./pkg/fcmul/lexer

test: fcmul_test

clean:
	rm -rf $(BINDIR)
