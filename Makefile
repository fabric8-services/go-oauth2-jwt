.PHONY: tools
tools: ## Installs required go tools
	@go get -u github.com/dgrijalva/jwt-go
	@go get -u github.com/codegangsta/negroni
	@go get -u github.com/gorilla/mux
	@go get -u github.com/stretchr/testify/assert

.PHONY: install
install: ## Fetches all dependencies using glide
	glide install -v

.PHONY: up
update: ## Updates all dependencies defined for glide
	glide up -v

.PHONY: compile
test:
	@go build

.PHONY: format ## Removes unneeded imports and formats source code
format:
	@goimports -l -w pkg
