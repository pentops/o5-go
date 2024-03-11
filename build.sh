#!/bin/bash
set -ue

cd "$(dirname "$0")"

ls -d */ | xargs rm -r
buf generate ../o5-pb
