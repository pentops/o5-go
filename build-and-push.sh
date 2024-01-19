#!/bin/bash
set -ue

# Build
./build.sh

go test ./...

git add .
git commit -m "Manual Build"
git push
