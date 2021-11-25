#!/bin/bash

if [ "$(docker ps -q -f name=user-mongo-db)" ]; then
  docker rm -f user-mongo-db
fi
