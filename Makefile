VERSION := $(shell git describe --tags --always)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

LDFLAGS := -X 'goauth/internal/version.Version=$(VERSION)' \
		   -X 'goauth/internal/version.Commit=$(COMMIT)' \
		   -X 'goauth/internal/version.BuildTime=$(DATE)'

generate-secret:
	echo "APP_SECRET=$$(openssl rand -hex 32)" > .env

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o bin/goauth-linux

clear:
	rm -rf bin/*
