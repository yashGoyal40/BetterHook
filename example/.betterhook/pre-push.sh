#!/bin/bash

echo "🚀 Lefthook is running pre-commit"
branch_name=$(git rev-parse --abbrev-ref HEAD)
regex="^(feat|bug)/[A-Z]+-[0-9]+$"

if ! echo "$branch_name" | grep -Eq "$regex" && \
   [ "$branch_name" != "master" ] && \
   [ "$branch_name" != "@staging" ] && \
   [ "$branch_name" != "@develop" ]; then
    echo "❌ ERROR: Invalid branch name '$branch_name'"
    echo "✅ Allowed formats: feat/JIRA-123 or bug/JIRA-123"
    exit 1
fi

echo "✅ Branch name '$branch_name' is valid."
