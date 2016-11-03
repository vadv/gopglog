BINARY=./bin/gopglog
SOURCEDIR=./src
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
GOPATH := ${PWD}:${GOPATH}

export GOPATH

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o ${BINARY} $(SOURCEDIR)/main.go

run: clean $(BINARY)
	${BINARY} -log-file logs/postgresql.log

clean:
	rm -f $(BINARY)
