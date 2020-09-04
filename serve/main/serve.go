package main

import (
	"context"
	"fmt"
	"helloGrpc/etcdv3"
	"log"
	"math/rand"
	"net"
	"time"

	pb "helloGrpc/proto"

	"google.golang.org/grpc"
)

const (
	// Address 监听地址
	Address string = "localhost:9001"
	// Network 网络通信协议
	Network string = "tcp"
	// SerName 服务名称
	SerName string = "simple_grpc"
)

var EtcdEndpoints = []string{"localhost:2379"}

func main() {
	listen, err := net.Listen(Network, Address) // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	pb.RegisterServeRouteServer(s, ServeRoute{})
	pb.RegisterDetailServer(s, DetailRoute{})

	//把服务注册到etcd
	ser, err := etcdv3.NewServiceRegister(EtcdEndpoints, SerName, Address, 5)
	if err != nil {
		log.Fatalf("register service err: %v", err)
	}
	defer ser.Close()

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
