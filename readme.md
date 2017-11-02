## Project
> This project illustrates the use of Docker with Go to build a RESTful service to transform data. It provides a service through a simple container.

## Running the Docker container
In the `dist-linux` folder, there is the `datawasherLinux` server and support config and test files. To run this service as a container, just use:

In the `docker` folder.
```
linux-docker-build.sh
linux-docker-up.sh
```
## Using the datawasher service

### GET endpoints
These endpoints will return JSON objects
```
http://localhost:8000/random_contact
http://localhost:8000/contacts
```

### POST endpoints
This endpoint will transform POST a JSON payload of contacts and return the same payload with transformed first_names.

See `testing/post-upload-contacts.sh` for a curl command to test the endpoint.
```
http://localhost:8000/json_contacts
```

## Development

| Docker Dev Environment                      |
|---------------------------------------------|
| `docker pull kyledinh/devlinux`             |
| `docker run -it --rm -v /Users/kyle/src:/opt/src -e GOPATH=/opt kyledinh/devlinux bash` |

The `$GOPATH` is set to `/opt`. I would attach my host's GOPATH to this directory and set a volume to `/opt/src`.

* `kyledinh/devlinux` is built from this repo: https://github.com/kyledinh/docker/tree/master/images
