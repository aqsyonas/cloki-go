#!/bin/bash
#
# cLoki-go from qxip
# Docker Builder & Slimmer
#
# ENV OPTIONS:
#    PUSH     = push images to dockerhub
#    SLIM     = build slim docker image
#    REPO     = default image respository/name
#    TAG      = defailt image tag (server)

REPO=${REPO:-sipcapture/webapp}
TAG=${TAG:-latest}

echo "Building HEPIC docker ..."
docker build -t $REPO:$TAG .
if [ ! -z "$PUSH" ]; then
  echo "Pushing $REPO:$TAG ..."
  docker push $REPO:$TAG
fi
