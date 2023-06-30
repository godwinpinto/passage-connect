VERSION    := 0.0.1
#COMMIT_SHA := $(shell git rev-parse --short HEAD)

all: build

build:
	go build -buildmode=c-shared -o build/pam_passage_connect.so

.PHONY: deps build install test fmt
