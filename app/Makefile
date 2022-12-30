run:
	go run main.go

build:
	go build

test: generate
	go test library/usecases

clean_mocks:
	rm ./*/mock_*

generate:
	go generate ./...

.PHONY: test, clean_mocks, generate, run, build