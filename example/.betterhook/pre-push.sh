#!/bin/bash

# Pre-push hook script for BetterHook

echo "Running pre-push checks..."

# Run tests before pushing
if command -v go test &> /dev/null; then
  echo "🛠 Running tests..."
  go test ./... || { echo "❌ Tests failed! Fix errors before pushing."; exit 1; }
fi

# Check for TODO comments
TODO_COUNT=$(grep -r "TODO" --exclude-dir={.git,vendor} . | wc -l)
if [[ "$TODO_COUNT" -gt 0 ]]; then
  echo "⚠️  Found $TODO_COUNT TODO comments. Consider resolving them before pushing."
fi

# Check for large files (e.g., >5MB)
LARGE_FILES=$(git diff --cached --name-only | xargs -I {} find {} -size +5M 2>/dev/null)
if [[ -n "$LARGE_FILES" ]]; then
  echo "⚠️  The following files are larger than 5MB and may slow down the repository:"
  echo "$LARGE_FILES"
  exit 1
fi

echo "✅ Pre-push checks passed! Pushing..."
exit 0
