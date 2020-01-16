#!/usr/bin/env bash

GOLANGCI_VERSION=v1.22.2

mkdir -p bin
curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" \
  | sh -s ${GOLANGCI_VERSION}
