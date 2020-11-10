#!/bin/bash

PACKAGE="cloki-go"
VERSION="1.0.0"
ARCH="amd64"

# Create Distribution

# BUILD DEB PACKAGE
EXT="deb"
docker run --rm \
  -v $PWD:/tmp/pkg \
  -e VERSION="$VERSION" \
  goreleaser/nfpm pkg --config /tmp/pkg/$PACKAGE.yaml --target "/tmp/pkg/$PACKAGE-$VERSION-$ARCH.$EXT"

# BUILD RPM PACKAGE
EXT="rpm"
docker run --rm \
  -v $PWD:/tmp/pkg \
  -e VERSION="$VERSION" \
  goreleaser/nfpm pkg --config /tmp/pkg/$PACKAGE.yaml --target "/tmp/pkg/$PACKAGE-$VERSION-$ARCH.$EXT"

