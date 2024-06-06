package service

import (
	"log"
	"mission/task5/dao"
	"mission/task5/dao/dbop"
	"mission/task5/dao/model"
	pb "mission/task5/pb"

	"golang.org/x/net/context"
)

type userserver struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() pb.UserServiceServer {
	return &userserver{}
}

func (s *userserver) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	email := in.GetEmail()
	err := dbop.CreateUser(dao.DB, &model.User{Name: name, Email: email})
	if err != nil {
		return &pb.UserReply{Msg: "创建用户失败"}, err
	}
	return &pb.UserReply{Msg: "创建用户 " + name + " 成功"}, nil
}

func (s *userserver) UpdateUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	email := in.GetEmail()
	err := dbop.UpdateUser(dao.DB, name, email)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.UserReply{Msg: "更新用户 " + name + " 成功"}, nil
}

func (s *userserver) DeleteUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	name := in.GetName()
	err := dbop.DeleteUser(dao.DB, name)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.UserReply{Msg: "删除用户 " + name + " 成功"}, nil
}
