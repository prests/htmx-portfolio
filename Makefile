# ==================================================================================== #
# DEV
# ==================================================================================== #

.PHONY: dev/app
dev/app:
	@echo 'Starting app in development mode'
	air

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/app: build application
.PHONY: build/app
build/app: build/template
	@echo 'Building app'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

.PHONY: build/template
build/template:
	@echo 'Building templates'
	templ generate

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...


## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...' go mod tidy
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'