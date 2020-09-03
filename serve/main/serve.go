package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "helloGrpc/proto"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000") // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	pb.RegisterServeRouteServer(s, ServeRoute{})
	pb.RegisterDetailServer(s, DetailRoute{})
	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

type ServeRoute struct{}

func (h ServeRoute) GetUser(ctx context.Context, in *pb.Id) (*pb.User, error) {
	resp := &pb.User{
		Time: time.Now().UnixNano() / 1e6,
		Name: fmt.Sprintf("%d,username", in.Id),
	}
	fmt.Println(resp)
	return resp, nil
}

func (h ServeRoute) GetActivity(ctx context.Context, in *pb.Name) (*pb.Activity, error) {
	tp := pb.Tp(rand.Intn(4))
	resp := &pb.Activity{
		Name: in.Name,
		Tp:   tp,
	}
	fmt.Println(resp)
	return resp, nil
}

type DetailRoute struct{}

func (d DetailRoute) GetUserInfo(ctx context.Context, in *pb.UserId) (*pb.UserInfo, error) {
	resp := &pb.UserInfo{
		Name: "hello grpc",
		Id:   in.Id,
	}

	fmt.Println(resp)
	return resp, nil
}
