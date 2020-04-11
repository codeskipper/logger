ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

.PHONY: test
test: ## Test with coverage
	@go test -short -cover $(ALL_PACKAGES)

.PHONY: test-cover-html
test-cover-html: ## Export test coverage to html
	mkdir -p out/
	@echo "mode: count" > ./out/coverage-all.out
	@$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=./out/coverage.out -covermode=count $(pkg);\
	tail -n +2 ./out/coverage.out >> ./out/coverage-all.out;)
	@go tool cover -html=./out/coverage-all.out -o out/coverage.html
	@open out/coverage.html

.PHONY: help
help: ## Shows help
	@echo
	@echo 'Usage:'
	@echo '    make <target>'
	@echo
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
