package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mission/task5/dao"
	"mission/task5/grpc/service"
	pb "mission/task5/pb"
	"net"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

type noticeServer struct {
	pb.UnimplementedNoticeServiceServer
}

func main() {
	dao.InitDB()
	log.Print("gorm初始化成功")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserServer())
	pb.RegisterNoticeServiceServer(s, service.NewNoticeServer())

	log.Println("grpc server listening at localhost:50051")

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
