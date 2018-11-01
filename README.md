# grpc-golang-server
A gRPC server written in Golang

The purpose is to benchmark how many requests client can make to server considering REST and gRPC as protocols.

There are two applications: one server written in Go (golang-grpc-server) and other written in Node.js (node-grpc-client).

Get both from repositories so they can talk with each other and you can get some benchmarks.

## Golang install

First of all, clone repository into your ```$GOPATH/src``` folder

### Install gRPC lib
```
$ go get -u google.golang.org/grpc
```

### Install Protocol Buffers v3
Choose the last release of **protoc.zip** from https://github.com/google/protobuf/releases, unzip and put `bin` folder in your `PATH`:
```
$ protoc --version
libprotoc 3.6.1
```

### Install the **protoc plugin** for Go
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Compile proto file
```
$ cd grpc-golang-server
$ protoc -I calculator/ calculator/calculator.proto --go_out=plugins=grpc:calculator
```

### Run server
```
$ go run main.go
gRPC server listening at port :50051
```