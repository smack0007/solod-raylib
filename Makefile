.PHONY: bind example help
.DEFAULT_GOAL := help

INC_DIR := $(shell pkg-config --variable=includedir raylib 2>/dev/null)

ifeq ($(INC_DIR),)
  INC_DIR := /opt/homebrew/include
  FLAGS  := -I/opt/homebrew/include -L/opt/homebrew/lib
else
  FLAGS  := $(shell pkg-config --cflags-only-I --libs-only-L raylib)
endif

CFLAGS ?= $(FLAGS)

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  bind                - Generate the bindings"
	@echo "  example name=<name> - Build an example program"

bind:
	@sobind -o libraylib/extern.go -pkg=libraylib -I $(INC_DIR) -style=cap $(INC_DIR)/raylib.h $(INC_DIR)/raymath.h

example:
	@CFLAGS="$(CFLAGS)" so build -check=sanitize -o ./build/$(name) ./example/$(name)
