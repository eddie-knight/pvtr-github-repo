PACKNAME=pvtr-github-repo
BUILD_FLAGS=-X 'main.GitCommitHash=`git rev-parse --short HEAD`' -X 'main.BuiltAt=`date +%FT%T%z`'
BUILD_WIN=@env GOOS=windows GOARCH=amd64 go build -o $(PACKNAME).exe
BUILD_LINUX=@env GOOS=linux GOARCH=amd64 go build -o $(PACKNAME)
BUILD_MAC=@env GOOS=darwin GOARCH=amd64 go build -o $(PACKNAME)-darwin

default: help
release: package release bin
release-candidate: package release-candidate
binary: package build
release: release-nix release-win release-mac

help: # Display help
	@awk -F ':|##' \
		'/^[^\t].+?:.*?##/ {\
			printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
		}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "  >  Building binary ..."
	@go build -o $(PACKNAME) -ldflags="$(BUILD_FLAGS)"

package: tidy test ## Package static files
	@echo "  >  Packaging static files..."

test: ## Run tests
	@echo "  >  Validating code ..."
	@go vet ./...
	@go clean -testcache
	@go test ./...

tidy: ## Tidy
	@echo "  >  Tidying go.mod ..."
	@go mod tidy

test-cov: ## Run tests and generate coverage output
	@echo "Running tests and generating coverage output ..."
	@go test ./... -coverprofile coverage.out -covermode count
	@sleep 2 # Sleeping to allow for coverage.out file to get generated
	@echo "Current test coverage : $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') %"

release-candidate: tidy test ## Build the release candidate
	@echo "  >  Building release candidate for Linux..."
	$(BUILD_LINUX) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=nix-rc'"
	@echo "  >  Building release candidate for Windows..."
	$(BUILD_WIN) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=win-rc'"
	@echo "  >  Building release for Darwin..."
	$(BUILD_MAC) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=darwin-rc'"

release-nix: ## Build the release for Linux
	@echo "  >  Building release for Linux..."
	$(BUILD_LINUX) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=linux'"

release-win: ## Build the release for Windows
	@echo "  >  Building release for Windows..."
	$(BUILD_WIN) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=windows'"

release-mac: ## Build the release for OSX
	@echo "  >  Building release for Darwin..."
	$(BUILD_MAC) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=darwin'"

docker-build: ## Build the container image
	@echo "  >  Building docker image ..."
	@docker build -t $(PACKNAME) .

docker-run: ## Run the container image
	@echo "  >  Running docker image ..."
	@docker run -it -v ${PWD}/config.yml:/.privateer/config.yml ${PACKNAME}

bin: ## Move the binary to the privateer bin directory
	@mv $(PACKNAME)* ~/privateer/bin
