package grpc

import (
	"awesomePet/api/debug"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial(RpcAddress, grpc.WithInsecure())
	debug.PanicErr(err)
	//defer conn.Close()
}
