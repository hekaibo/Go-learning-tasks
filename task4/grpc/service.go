package main

import (
	"log"
	"mission/task4/gorm"
	"mission/task4/gorm/dbop"
	"mission/task4/gorm/model"
	"net"

	pb "mission/task4/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	email := in.GetEmail()
	err := dbop.CreateUser(gorm.DB, &model.User{Name: name, Email: email})
	if err != nil {
		return &pb.UserReply{Msg: "创建用户失败"}, err
	}
	return &pb.UserReply{Msg: "创建用户 " + name + " 成功"}, nil
}

func (s *server) GetUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserDetailReply, error) {
	name := in.GetName()
	user, err := dbop.GetUserByName(gorm.DB, name)
	if err != nil {
		log.Println(err)
		return &pb.UserDetailReply{Msg: "查询失败"}, err
	}
	if user == nil {
		return &pb.UserDetailReply{Msg: "未查询到用户"}, err
	}
	var userrequest []*pb.UserRequest
	userrequest = append(userrequest, &pb.UserRequest{Id: uint64(user.Model.ID), Name: user.Name, Email: user.Email})
	return &pb.UserDetailReply{
		UserDetail: userrequest,
	}, nil
}

func (s *server) GetAllUser(ctx context.Context, in *pb.UserRequest) (*pb.UserDetailReply, error) {
	users, err := dbop.GetAllUser(gorm.DB)
	if err != nil {
		log.Fatal(err)
	}
	var userrequest []*pb.UserRequest
	for _, user := range users {
		userrequest = append(userrequest, &pb.UserRequest{Id: uint64(user.Model.ID), Name: user.Name, Email: user.Email})
	}
	return &pb.UserDetailReply{
		UserDetail: userrequest,
	}, nil
}
func (s *server) UpdateUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	email := in.GetEmail()
	err := dbop.UpdateUser(gorm.DB, name, email)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.UserReply{Msg: "更新用户 " + name + " 成功"}, nil
}
func (s *server) DeleteUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	err := dbop.DeleteUser(gorm.DB, name)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.UserReply{Msg: "删除用户 " + name + " 成功"}, nil
}

func main() {
	gorm.InitDB()
	log.Print("gorm初始化成功")
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
