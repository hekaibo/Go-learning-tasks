package main

import (
	"log"
	"net"

	pb "mission/task2/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	id := in.GetId()
	if id == "1" {
		return &pb.UserReply{Name: "zhangsan ", Age: "21"}, nil
	} else if id == "2" {
		return &pb.UserReply{Name: "lisi ", Age: "25"}, nil
	} else if id == "3" {
		return &pb.UserReply{Name: "wangwu ", Age: "31"}, nil
	} else {
		return &pb.UserReply{Name: "zhaoliu ", Age: "11"}, nil
	}

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("grpc server listening at localhost:50051")

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
