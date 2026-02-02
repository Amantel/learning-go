#!/bin/bash
# Push current branch and output PR creation URL with template
#
# Usage: Run from the repo root directory
#   ./push.sh

set -e

BRANCH=$(git rev-parse --abbrev-ref HEAD)
REMOTE_URL=$(git config --get remote.origin.url)

# Convert SSH URL to HTTPS URL for browser
# git@github.com:user/repo.git -> https://github.com/user/repo
HTTPS_URL=$(echo "$REMOTE_URL" | sed -e 's/git@github.com:/https:\/\/github.com\//' -e 's/\.git$//')

echo "=== Pushing $BRANCH ==="
git push -u origin "$BRANCH"

PR_URL="${HTTPS_URL}/compare/${BRANCH}?expand=1&template=${BRANCH}.md"

# Extract owner/repo from URL
REPO_PATH=$(echo "$REMOTE_URL" | sed -e 's/git@github.com://' -e 's/\.git$//')

# Check if PR already exists using GitHub API
PR_EXISTS=$(curl -s "https://api.github.com/repos/${REPO_PATH}/pulls?head=${REPO_PATH%%/*}:${BRANCH}&state=open" | grep -c '"number"' || true)

if [ "$PR_EXISTS" -gt 0 ]; then
    echo ""
    echo "=== PR already exists ==="
else
    echo ""
    echo "=== Opening PR creation page ==="
    echo "$PR_URL"
    open "$PR_URL"
fi
