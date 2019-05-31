# Sample rest api

This is sample rest api service with single endpoint written in go, to be used in [api-status-operator](https://github.com/pratikjagrut/api-status-operator) project.

## Quick start

### Prerequisites
- [go](https://golang.org/dl/) version v1.10+
- go get gopkg.in/src-d/go-git.v4

### Steps to run

```
$ mkdir $GOPATH/src/github.com/pratikjagrut
$ cd $GOPATH/src/github.com/pratikjagrut
$ git clone https://github.com/pratikjagrut/status-rest-api.git
$ cd status-rest-api
$ go build -o app
$ ./app
```

```
Output:
$ curl localhost:8080
[{"CommitID":"4281f56df39670db14a70e0435720d8ba6e74e03"}]
```

### Run using docker image

```
$ docker build -t apistatus .
$ docker run -p 8080:8080 -d apistatus
```

```
Output:
$ curl localhost:8080
[{"CommitID":"4281f56df39670db14a70e0435720d8ba6e74e03"}]
```
