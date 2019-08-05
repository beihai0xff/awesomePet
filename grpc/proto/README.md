#  gRPC 远程程序调用
## Usage
####Requirements
* install [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases) and add to path
* install protoc-golang
```bash
$ go get -u github.com/grpc/grpc-go
$ go get -u github.com/golang/protobuf/protoc-gen-go 
```
#### compile proto source code
```bash
$ protoc --go_out=plugins=grpc:. hello.proto
```
