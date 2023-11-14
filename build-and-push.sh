#!/bin/bash
set -ue

# Build
./build.sh

git add .
git commit -m "Manual Build"
git push
