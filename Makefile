ifeq ($(OS),Windows_NT)
EXT:=.exe
else
EXT:=
endif

.PHONY: all
all: pkg cmd

.PHONY: cmd
cmd:
	go build -o bin/linear${EXT} ./cmd/linear

.PHONY: pkg
pkg:
	@for pkg in command; do \
		go test -v --coverprofile coverage.$${pkg} ./pkg/$${pkg} --trace=trace.$${pkg}; \
		go tool cover -func=coverage.$${pkg}; \
	done
