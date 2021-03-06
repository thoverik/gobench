GO ?= go

PACKAGES := $(shell go list ./...)
EXAMPLES := $(shell go list ./... | grep "examples")

.PHONY: lint build examples tools ent statik

# run lint
lint-pkgs:
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/staticcheck
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell

lint:
	$(exit $(go fmt ./... | wc -l))
	go vet ./...
	# ignore statik.go file
	find . -type f -name "*.go" ! -name "statik.go" | xargs misspell -error -locale US
	staticcheck $(go list ./... | grep -v ent/privacy)

build:
	go build -o gobench ./

examples:
	$(foreach var, $(EXAMPLES), go build -o ./.bin/${var} $(var);)

tools:
	go build -o ./.bin/github.com/thoverik/gobench/tools/gobench-viewer github.com/thoverik/gobench/tools/gobench-viewer

# generate ent models
ent:
	entc generate ./ent/schema

# generate statik file for web ui
statik:
	statik -src=./web/ui/react-app/build -dest=./web -f
