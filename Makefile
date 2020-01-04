BINARY=terraform-provider-confluence

build:
	go build -o ${BINARY}

clean:
	rm ${BINARY}

all: build
