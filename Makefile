.PHONY: default

deps:
	go get -t -v

test:
	go test -covermode=count -coverprofile=count.out -v
	go tool cover -html=count.out -o coverage.html

build:
	go generate -v
	go build -o three .


default: deps build