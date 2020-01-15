#!/usr/bin/env bash

# Check gofmt
echo "==> Checking for linting errors..."

if ! which revive > /dev/null; then
    echo "==> Installing revive..."
    go get -u github.com/mgechev/revive
fi

err_files=$(revive -config revive.toml -formatter unix)

if [[ -n ${err_files} ]]; then
    echo 'Linting errors found in the following places:'
    echo "${err_files}"
    echo "Please handle returned errors. You can check directly with \`make lint\`"
    exit 1
fi

exit 0
