#!/bin/bash
env CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o ./build/chat ./main.go
