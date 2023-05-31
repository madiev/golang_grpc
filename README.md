# grpc golang

## prepare

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go install github.com/ktr0731/evans@latest (debug)

export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=$PWD --go_out=$PWD --go-grpc_out=$PWD $PWD/proto/myapp.proto
```

## build and run server
```
go build -v ./cmd/server
./server
```

## run client
```
go run ./cmd/client/main.go 5 9
```

## run evans
```
evans proto/myapp.proto -p 8090
```