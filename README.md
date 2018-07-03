# Quick start

Run your *server* node microservice with:

```bash
cd $GOPATH/src/gokit-grpc/car-microservice
go run main.go
```

Run your *client* node microservice with:

```bash
cd $GOPATH/src/gokit-grpc/car-client
go run main.go
```

Magic happens on port `3000`

```bash
curl -X POST http://localhost:3000/car -H 'content-type: application/json' -d  '{"name": "Enzo", "manifacturer": "Ferrari"}'

curl -X GET http://localhost:3000/car/1
```

# Golang utilities

Some utilities for development

## Dependency manager

Download `Dep` from https://github.com/golang/dep/releases.
Extract and export bin folder into your `$PATH`.

```bash
#DEP
export PATH=${PATH}:~/dep/bin/
```

## Protocol buffer compiler

Download protocol buffer compiler from https://github.com/google/protobuf/releases. 
Extract and export bin folder into your `$PATH`.

```bash
#PROTOC
export PROTOC_HOME=~/opt/protoc-3.2.0-osx-x86_64
export PATH=${PATH}:$PROTOC_HOME/bin/
```

Now you can compile `.proto` files with:

```bash
protoc example.proto --go_out=plugins=grpc:.
```

This command generates an `example.pb.go` file.


# Links

1. https://matthewbrown.io/2016/01/23/factory-pattern-in-golang/
2. https://www.alexedwards.net/blog/organising-database-access


