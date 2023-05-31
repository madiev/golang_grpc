# grpc golang

## prepare

```
export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=$PWD --go_out=$PWD --go-grpc_out=$PWD $PWD/proto/myapp.proto
```

## build server and run client
```
go build -v ./cmd/server
go run ./cmd/client/main.go 5 9
```