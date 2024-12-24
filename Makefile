.DEFAULT_GOAL := by_protoc

.PHONY: by_protoc
by_protoc:  ## build by protoc and protoc-gen-go
	protoc -I=. --go_out=. *.proto

.PHONY: by_buf
by_buf:  ## build by buf
	buf generate

.PHONY: help
help: ## prints this help message
	@echo "Usage: \n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

