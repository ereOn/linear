ifeq ($(OS),Windows_NT)
EXT:=.exe
else
EXT:=
endif

.PHONY: all
all: cmd

.PHONY: cmd
cmd:
	go build -o bin/linear${EXT} ./cmd/linear
