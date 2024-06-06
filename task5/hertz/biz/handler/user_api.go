package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mission/task5/pb"
	"net/http"
)

var conn *grpc.ClientConn

func InitConn() error {
	connect, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("did not connect: %v", err)
		return err
	}
	log.Println("connected to 'localhost:50051' grpc service")
	conn = connect
	return nil
}

func GetConn() *grpc.ClientConn {
	return conn
}

func CreateUser(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	email := c.Param("email")

	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name, Email: email}
	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"msg": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}

func UpdateUserByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	email := c.Param("email")

	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name, Email: email}
	res, err := client.UpdateUserByName(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"错误信息": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}

func DeleteUserByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Name: name}
	res, err := client.DeleteUserByName(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"错误信息": err.Error()})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}
