#!/bin/bash
set -e

echo "=== Running golangci-lint ==="
golangci-lint run .

echo ""
echo "=== Running tests ==="
go test -v -count=1 -race -timeout=1m .

if [ -f "./test.sh" ]; then
    echo ""
    echo "=== Running test.sh ==="
    ./test.sh
fi

echo ""
echo "=== All checks passed ==="
