# Build and run main file
run:
	go run main.go

# Run tests without coverage
t:
	go test library/usecases library/http/controller

# Run tests with coverage
tc:
	go test library/usecases library/http/controller -cover

# Clean "temp" files
clean:
	rm ./*/mock_*

# Generate files
gen:
	go generate ./...

.PHONY: t tc, clean, gen, run