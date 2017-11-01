#!/bin/bash
docker-compose -f ../dist-linux/docker-compose.yml build --no-cache
docker image ls -a
