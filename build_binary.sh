#!/bin/bash
# BUILD GO BINARY
docker run --rm \
  -v $PWD:/app \
  golang:1.13 \
  bash -c "cd /app && make modules && make all"