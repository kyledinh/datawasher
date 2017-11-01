| Docker Dev Environment                      |
|---------------------------------------------|
| `docker pull kyledinh/devlinux`             |
| `docker run -it --rm -v /Users/kyle/src:/opt/src -e GOPATH=/opt kyledinh/devlinux bash` |

The `$GOPATH` is set to `/opt`. I would attach my host's GOPATH to this directory and set a volume to `/opt/src`.
