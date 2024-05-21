.PHONY: all fmt test bench lint
all: fmt test bench lint clean

fmt:
	go fmt ./...

test:
	go test -v ./...

bench:
	go test -v -bench . ./... -benchmem -test.timeout 60m

lint:
	go vet ./...

clean:
	rm -rf ./coverage.out

deps:
	go mod tidy

