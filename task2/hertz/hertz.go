package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mission/task2/pb"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 用于与grpc服务器通信
	client := pb.NewUserServiceClient(conn)

	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.GET("userinfo/:id", func(c context.Context, ctx *app.RequestContext) {
		id := ctx.Param("id")

		req := &pb.UserRequest{Id: id}
		res, err := client.GetUserInfo(context.Background(), req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, utils.H{
			"name": res.Name,
			"age":  res.Age,
		})
	})

	h.Spin()
}
