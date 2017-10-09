COMMANDS=linear \
		 linear-init \
		 linear-add \
		 linear-add-api

PACKAGES=command \
		 database

ifeq ($(OS),Windows_NT)
EXT:=.exe
else
EXT:=
endif

.PHONY: all
all: pkg cmd

.PHONY: cmd
cmd:
	@for cmd in ${COMMANDS}; do \
		go build -o bin/$${cmd}${EXT} ./cmd/$${cmd}; \
	done

.PHONY: pkg
pkg:
	@for pkg in ${PACKAGES}; do \
		go test -v --coverprofile coverage.$${pkg} ./pkg/$${pkg} --trace=trace.$${pkg}; \
		go tool cover -func=coverage.$${pkg}; \
	done
