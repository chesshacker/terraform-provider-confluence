TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')
PKG_NAME=confluence
BINARY_NAME=terraform-provider-confluence
WEBSITE_DIR=site
INSTALL_DIR=$(HOME)/.terraform.d/plugins

all: check test build

check: fmtcheck lintcheck errcheck vet

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lintcheck:
	@sh -c "'$(CURDIR)/scripts/lintcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vet:
	@echo "go vet ."
	@go vet $$(go list ./...) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

build: $(BINARY_NAME)

$(BINARY_NAME):
	go build -v -o $(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf $(WEBSITE_DIR)

install: $(BINARY_NAME)
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

fmt:
	gofmt -w $(GOFMT_FILES)

.PHONY: all build check clean errcheck fmt fmtcheck install lintcheck test testacc uninstall vet
