#!/bin/bash
set -ue

ls -d */ | xargs rm -r
buf generate ../o5-pb/proto
