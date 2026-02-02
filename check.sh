#!/bin/bash
# Pre-launch check script for OTUS homework
#
# Usage: Run from the repo root directory
#   ./check.sh
#
# The script automatically detects the homework folder from the current branch name.
# For example, if you're on branch 'hw01_hello_otus', it will check the hw01_hello_otus/ folder.
#
# Steps performed:
#   1. Remove .sync file if exists
#   2. Run go mod tidy
#   3. Run golangci-lint formatter
#   4. Run golangci-lint linter
#   5. Run tests
#   6. Run test.sh if exists

set -e

BRANCH=$(git rev-parse --abbrev-ref HEAD)
DIR="$BRANCH"

if [ ! -d "$DIR" ]; then
    echo "Error: Directory '$DIR' not found"
    exit 1
fi

echo "=== Checking $DIR ==="

echo ""
echo "=== Removing .sync if exists ==="
rm -f "$DIR/.sync"

echo ""
echo "=== Running go mod tidy ==="
(cd "$DIR" && go mod tidy)

echo ""
echo "=== Running formatter ==="
(cd "$DIR" && golangci-lint fmt .)

echo ""
echo "=== Running golangci-lint ==="
(cd "$DIR" && golangci-lint run .)

echo ""
echo "=== Running tests ==="
(cd "$DIR" && go test -v -count=1 -race -timeout=1m .)

if [ -f "$DIR/test.sh" ]; then
    echo ""
    echo "=== Running test.sh ==="
    (cd "$DIR" && ./test.sh)
fi

echo ""
echo "=== All checks passed ==="
