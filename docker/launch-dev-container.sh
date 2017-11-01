#!/bin/bash

# customize for Kyle Dinh's Mac Book
# Uses this dev image: https://github.com/kyledinh/docker/tree/master/images
docker run -it --rm -p 8000:8000 -v /Users/kyle/src:/opt/src -e GOPATH=/opt kyledinh/devlinux bash
