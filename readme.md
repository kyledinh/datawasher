## Project
> This project illustrates the use of Docker with Go to build a RESTful service to transform data. It provides a service through a simple container.

#### Technologies used
* Docker - https://www.docker.com/
* Go - https://golang.org/
* Gin Go Webserver - https://github.com/gin-gonic/gin
* Go-SimpleJson - https://github.com/bitly/go-simplejson

## Running the Docker container
In the `dist-linux` folder, there is the `datawasherLinux` server and support config and test files. To run this service as a container, just use:

In the `docker` folder.
```
linux-docker-build.sh
linux-docker-up.sh
```
## Using the datawasher services

### GET Create endpoint

```
http://localhost:8000/create?limit=10&first_name=MOX_RFN&last_name=MOX_RLN&email=MOX_EMAIL&addr=MOX_RSA&code=MOX_RI_1000&state=MOX_STATE&sex=MOX_RSMF
```
The `create` endpoint will generate an array of where the you can name the fields and use generator via the query string. In the example, the fields in the url param maps to the JSON fields created.

<img src="https://raw.github.com/kyledinh/datawasher/master/assets/get-create.png" width="1200" />


### Other GET endpoints
These endpoints will return JSON objects. These contacts are predefined from the `data/contacts.csv` file.
```
http://localhost:8000/random_contact
http://localhost:8000/contacts
```


### POST endpoints
This endpoint will transform POST a JSON payload of contacts and return the same payload with transformed first_names.

```
http://localhost:8000/json_contacts
```

See `testing/post-upload-contacts.sh` for a curl command to test the endpoint.
<img src="https://raw.github.com/kyledinh/datawasher/master/assets/contact.json.png" width="800" />

After datawashing
<img src="https://raw.github.com/kyledinh/datawasher/master/assets/washed.contacts.png" width="800" />

```
http://localhost:8000/washer?last_name=MOX_RLN&first_name=MOX_RFN&email=MOX_EMAIL
```
See `testing/post-washer.sh` for a curl command to test the endpoint. This endpoint allows you to describe the fields used in your json payload. In the query string you can map your fields to presets of "washer" functions.

| Washer Presets | |
|----------------|-|
| MOX_RFN        | Transforms with a random first name |
| MOX_RLN        | Transforms with a random last name  |
| MOX_EMAIL      | Used with MOX_RFN and MOX_RLN to generate an email address |
| MOX_RSA        | Random street address like: 123 Elm Street |
| MOX_RI_100     | Random integer from 0-99 |
| MOX_RI_1000    | Random integre from 0-999 |
| MOX_STATE      | Random state code of 50 US states, ie: CA |
| MOX_RSMF       | Random sex flag, connected with first name |


## Development

| Docker Dev Container                        |
|---------------------------------------------|
| `docker pull kyledinh/devlinux`             |
| `docker/launch-dev-container.sh`            |

You will need to customize `docker/launch-dev-container` for where you place your source code and $GOPATH on your host:
* -v `/Users/kyle/src`:/opt/src
* -w "/opt/src/github.com/kyledinh/datawasher"

The `$GOPATH` is set to `/opt`. I would attach my host's GOPATH to this directory and set a volume to `/opt/src`.

* `kyledinh/devlinux` is built from this repo: https://github.com/kyledinh/docker/tree/master/images

### Dev Resources
* dataset sources - http://mbejda.github.io/
* http://www.outpost9.com/files/WordLists.html
* https://godoc.org/github.com/bitly/go-simplejson#Json.Encode
* https://gobyexample.com/json
* https://gin-gonic.github.io/gin/
