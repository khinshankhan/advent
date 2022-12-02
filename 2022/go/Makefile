COMMANDS = $(notdir $(wildcard ./cmds/*))

.PHONY: all build $(COMMANDS)

all: build

build: $(COMMANDS)

$(COMMANDS): % : cmds/%/main.go
	go build -o bin/$@ cmds/$@/main.go

clean:
	-rm bin/*
