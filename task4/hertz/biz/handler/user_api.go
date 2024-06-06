package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mission/task4/pb"
	"net/http"
)

func CreateUser(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	email := c.Param("email")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name, Email: email}
	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})

	//err := dbop.CreateUser(dao.DB, &model.User{Name: name, Email: email})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//c.JSON(consts.StatusOK, utils.H{"状态：": "新建 " + name + " 用户成功"})

}

func GetUserByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	req := &pb.UserRequest{Name: name}
	res, err := client.GetUserByName(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
		log.Fatal(err)
	}
	if res.GetUserDetail() == nil {
		c.JSON(http.StatusNotFound, utils.H{"msg": res.GetMsg()})
	} else {
		c.JSON(consts.StatusOK, utils.H{"user:": res.UserDetail})
	}
}

func GetAllUser(ctx context.Context, c *app.RequestContext) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{}
	res, err := client.GetAllUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
		log.Fatal(err)
	}
	c.JSON(consts.StatusOK, utils.H{"users:": res.UserDetail})
}

func UpdateUserByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	email := c.Param("email")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name, Email: email}
	res, err := client.UpdateUserByName(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}

func DeleteUserByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name}
	res, err := client.DeleteUserByName(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}
