BINARY=terraform-provider-confluence
INSTALL_DIRECTORY=${HOME}/.terraform.d/plugins

${BINARY}: *.go go.mod go.sum
	go build -o ${BINARY}

build: ${BINARY}

clean:
	rm ${BINARY}

install: ${BINARY}
	mkdir -p "${INSTALL_DIRECTORY}"
	cp ${BINARY} "${INSTALL_DIRECTORY}"

all: build
