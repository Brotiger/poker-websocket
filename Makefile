SHELL := /bin/bash

build:
	go build -o main ./cmd/main.go
.PHONY: build

up:
	./main
.PHONY: up

run:
	go run ./cmd
.PHONY: run