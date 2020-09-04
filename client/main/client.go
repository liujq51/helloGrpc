package main

import (
	"fmt"
	"helloGrpc/etcdv3"
	pb "helloGrpc/proto" // 引入proto包
	"log"
	"time"

	"google.golang.org/grpc/resolver"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	EtcdEndpoints = []string{"localhost:2379"}
	SerName       = "simple_grpc"
	grpcClient    pb.ServeRouteClient
)

func main() {
	r := etcdv3.NewServiceDiscovery(EtcdEndpoints)
	resolver.Register(r)
fmt.Println(fmt.Sprintf("%s:///%s", r.Scheme(), SerName))
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", r.Scheme(), SerName), grpc.WithBalancerName("round_robin"), grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	grpcClient = pb.NewServeRouteClient(conn)
	for i := 0; i < 100; i++ {
		route(int32(i))
		time.Sleep(1 * time.Second)
	}

	//reqBody2 := &pb.Name{Name: "Hello"}
	//res2, err := rpc.GetActivity(context.Background(), reqBody2) //就像调用本地函数一样，通过serve2得到返回值
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("message from serve: ", res2.Name, res2.Tp)
	//
	//rpc2 := pb.NewDetailClient(conn)
	//reqBody3 := &pb.UserId{Id: 1}
	//res3, err := rpc2.GetUserInfo(context.Background(), reqBody3) //就像调用本地函数一样，通过serve3得到返回值
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("message from detail: ", res3.Name, res3.Id)
}

func route(i int32) {
	reqBody1 := &pb.Id{Id: i}
	res1, err := grpcClient.GetUser(context.Background(), reqBody1) //就像调用本地函数一样，通过serve1得到返回值
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from serve: ", res1.Name)
}
