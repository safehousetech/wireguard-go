PREFIX ?= /usr
DESTDIR ?=
BINDIR ?= $(PREFIX)/bin
export GO111MODULE := on

all: generate-version-and-build

MAKEFLAGS += --no-print-directory

generate-version-and-build:
	@export GIT_CEILING_DIRECTORIES="$(realpath $(CURDIR)/..)" && \
	tag="$$(git describe --dirty 2>/dev/null)" && \
	ver="$$(printf 'package main\n\nconst Version = "%s"\n' "$$tag")" && \
	[ "$$(cat version.go 2>/dev/null)" != "$$ver" ] && \
	echo "$$ver" > version.go && \
	git update-index --assume-unchanged version.go || true
	@$(MAKE) wireguard-go

wireguard-go: $(wildcard *.go) $(wildcard */*.go)
	go build -v -o "$@"

install: wireguard-go
	@install -v -d "$(DESTDIR)$(BINDIR)" && install -v -m 0755 "$<" "$(DESTDIR)$(BINDIR)/wireguard-go"

test:
	go test ./...

clean:
	rm -f wireguard-go

.PHONY: all clean test install generate-version-and-build

up: WG_QUICK_USERSPACE_IMPLEMENTATION=$(CURDIR)/wireguard-go
up: ENV_WG_PROCESS_FOREGROUND=1
up: LOG_LEVEL=debug
up:
	sudo WG_QUICK_USERSPACE_IMPLEMENTATION=$(WG_QUICK_USERSPACE_IMPLEMENTATION) \
	ENV_WG_PROCESS_FOREGROUND=$(ENV_WG_PROCESS_FOREGROUND) \
	LOG_LEVEL=$(LOG_LEVEL) \
	wg-quick up wg0

down:
	 sudo wg-quick down wg0

.PHONY: up down
