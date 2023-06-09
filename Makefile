default: install app/server

VERSION := $(shell git rev-parse --short HEAD)

install:
	go install ./...

fmt:
	gofmt -w -s .

fmt/list:
	gofmt -l -s .

build: install
	go build .

run:
	go run -tags=development .

test: install
	go test ./...

app:
	mkdir -p app

app/server: app **/*.go
	# Creates a linux build for copying into the alpine container
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app/mock-backend-demo .

docker: app/server
	docker build -t mock-backend-demo .

clean:
	rm -rf ./app ./main

.PHONY: install lint test app clean docker
