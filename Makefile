# Defined to avoid redundancy
pkg_name := snaparser_server
main_path := cmd/${pkg_name}/main.go
mkdir_cmd := mkdir -p bin/
build_cmd := go build -ldflags="-s -w" -o

# Run the program with sensible options for testing
run:
	go run ${main_path} --load config.toml

# Determine the OS
ifeq ($(OS), Windows_NT)
	OS_TARGET = Windows
else
	OS_TARGET := $(shell uname -s)
endif

# Build for the current platform
build:
	${mkdir_cmd}build
	${build_cmd} bin/build/ ${main_path}

# Build for specified OS
os: $(OS_TARGET)

# Build for Linux
Linux:
	${mkdir_cmd}linux
	GOOS=linux ${build_cmd} bin/linux/${pkg_name} ${main_path}

# Build for Darwin (MacOS)
Darwin:
	${mkdir_cmd}darwin
	GOOS=darwin ${build_cmd} bin/darwin/${pkg_name} ${main_path}

# Build for Windows
Windows:
	${mkdir_cmd}windows
	GOOS=windows ${build_cmd} bin/windows/${pkg_name}.exe ${main_path}

# Run linter
lint:
	golangci-lint run .

# Run tests
test:
	go test ./...

# Run tests, check for race conditions, report coverage
test-full:
	go test ./... -race -cover

# Run lint, test-full, build for all OS
all: lint test-full Linux Darwin Windows
