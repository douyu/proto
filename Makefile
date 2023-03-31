
.PHONY: init
# init tools
init:
	@echo "Installing tools from tools/tools.go"
	@cd tools && cat tools.go |grep _|awk -F '"' '{print $$2}' | xargs -tI % go install % && cd -

# update buf mod
update:
	cd api && buf mod update

.PHONY: generate
# generate code
generate:
	buf generate
	cd api && buf generate --template buf.gen.tag.yaml
	cd api && find . -type f -path '*.pb.go'|xargs -r rm

.PHONY: lint
# lint
lint:
	buf lint

# breaking
breaking:
	buf breaking --against https://github.com/douyu/proto/.git#branch=main,ref=HEAD~1,subdir=api

# test
test:
	go test -v -cover ./...

# validate openapi docs
validate:
	swagger validate gen/go/api/helloworld/v1/helloworld.swagger.json

# serve openapi docs
serve:
	swagger serve gen/go/api/helloworld/v1/helloworld.swagger.json

.PHONY: all

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help