run:
	go run main.go

test: generate
	go test library/usecases

testc: generate
	go test library/usecases -cover

clean-mocks:
	rm ./*/mock_*

generate:
	go generate ./...

.PHONY: test, clean_mocks, generate, run