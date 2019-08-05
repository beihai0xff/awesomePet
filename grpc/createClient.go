package grpc

import (
	"awesomePet/grpc/proto"
	"context"
	"fmt"
)

func CreateUserClient(uid uint64) bool {
	client := proto.NewCreateUserClient(conn)
	Value, err := client.CreateUser(context.Background(), &proto.CreateUserRequest{Uid: uid})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return Value.Result
}
