# helloworld

A simple Hello World using the [atlas app toolkit](https://github.com/infobloxopen/atlas-app-toolkit)

## Getting Started

Guide [here](https://github.com/infobloxopen/atlas-cli/wiki). 

### Prerequisites

In addition to the guide above, make sure to run the command

```
go get -u github.com/golang/protobuf/protoc-gen-go
```
Also, your $GOPATH should be set to ~/go, and $GOBIN should be set to ~/go/bin

### Installing

After the above steps are complete, go to your desired directory and run the following: 

```
atlas init-app -name helloworld -gateway
```

From there, you should be able to open the directory structure in your IDE of choice.

## Deployment

In the helloworld directory, execute the following to run the service:

```
go run cmd/server/*.go
```

## Running the tests

Open Postman, and send a `GET` request to `0.0.0.0:8080/helloworld/v1/version`. This should return the version of the project, which is `0.0.1`.

Sending a `POST` request to `0.0.0.0:8080/helloworld/v1/hello` with a JSON body consisting of something like `{ "name": "Broc" }` should return `"Hello Broc"`

## Versioning

I don't anticipate updating this code. It shall stay at 0.0.1
