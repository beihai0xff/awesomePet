package grpc

import (
	"awesomePet/api"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func init() {
	var err error
	conn, err = grpc.Dial(RpcAddress, grpc.WithInsecure())
	api.PanicErr(err)
	//defer conn.Close()
}
