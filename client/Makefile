VERSION    := 0.0.1
#COMMIT_SHA := $(shell git rev-parse --short HEAD)

all: build

build:
	go build -buildmode=c-shared -o passage_connect.so

#install:
#	cp passage_connect.so /usr/lib/security/

.PHONY: deps build install test fmt
