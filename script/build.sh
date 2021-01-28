#!/usr/bin/env bash
cd /app/src/ \
&& go mod download \
&& go build -o main
