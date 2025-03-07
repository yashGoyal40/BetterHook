#!/bin/bash

# Pre-commit hook script for BetterHook

echo "Running pre-commit checks..."

# Check for unstaged changes in tracked files
if ! git diff-index --quiet HEAD --; then
  echo "⚠️  You have unstaged changes. Please commit them or stash before proceeding."
  exit 1
fi

# Run formatting check (example: Go format check)
if command -v gofmt &> /dev/null; then
  unformatted=$(gofmt -l .)
  if [[ -n "$unformatted" ]]; then
    echo "⚠️  The following files need formatting:"
    echo "$unformatted"
    exit 1
  fi
fi

# Run linting (example: Go lint)
if command -v golint &> /dev/null; then
  golint ./...
fi

# Run tests
if command -v go test &> /dev/null; then
  echo "🛠 Running tests..."
  go test ./... || exit 1
fi

echo "✅ Pre-commit checks passed!"
exit 0
