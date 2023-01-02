# Build and run main file
run:
	go run main.go

# Run tests without coverage
t: gen
	go test library/usecases

# Run tests with coverage
tc: gen
	go test library/usecases -cover

# Clean "temp" files
clean:
	rm ./*/mock_*

# Generate files
gen:
	go generate ./...

.PHONY: t tc, clean, gen, run