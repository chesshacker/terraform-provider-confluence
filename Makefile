PKG_NAME=confluence
BINARY_NAME=terraform-provider-confluence
INSTALL_DIR=$(HOME)/.terraform.d/plugins
VERSION=$$(cat VERSION)
TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

all: check test build

check: bin/golangci-lint
	bin/golangci-lint run

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

build: $(BINARY_NAME)

$(BINARY_NAME):
	go build -v -o $(BINARY_NAME)

testacc:
	TF_ACC=1 source secrets.env && go test $(TEST) -v $(TESTARGS) -timeout 5m

fmt:
	gofmt -s -w $(GOFMT_FILES)

clean:
	rm -rf bin
	rm -rf site
	rm -f $(BINARY_NAME)

install: $(BINARY_NAME)
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)_v$(VERSION)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

bin/golangci-lint:
	scripts/get-golangci.sh

.PHONY: all build check clean fmt install test testacc uninstall
