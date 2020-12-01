COMMANDS = $(notdir $(wildcard ./cmds/*))

.PHONY: all build $(COMMANDS)

all: build

build: $(COMMANDS)

$(COMMANDS): % : cmds/%/main.go utils/utils.go
	go build -o bin/$@ cmds/$@/main.go

clean:
	-rm bin/*
