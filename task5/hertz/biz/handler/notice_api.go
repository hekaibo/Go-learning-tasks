package handler

import (
	context "context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
	pb "mission/task5/pb"
	"net/http"
	"strconv"
)

func CreateNotice(ctx context.Context, c *app.RequestContext) {
	title, _ := c.GetPostForm("title")
	content, _ := c.GetPostForm("content")
	pubuser, _ := c.GetPostForm("pubuser")

	client := pb.NewNoticeServiceClient(conn)

	req := &pb.NoticeRequest{Title: title, Content: content, Pubuser: pubuser}
	res, err := client.CreateNotice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"创建失败": err.Error()})
		log.Println(err)
	}
	c.JSON(http.StatusOK, utils.H{
		"msg": res.Msg,
	})
}

func ShowNotice(ctx context.Context, c *app.RequestContext) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	client := pb.NewNoticeServiceClient(conn)
	req := &pb.NoticeRequest{Id: id}
	res, err := client.ShowNotice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"msg": err.Error()})
	}
	if res.GetNoticeDetail() == nil {
		c.JSON(http.StatusInternalServerError, utils.H{"msg": res.Msg})
	} else {
		c.JSON(http.StatusOK, utils.H{"data": res.NoticeDetail})
	}
}

func ShowAllNotice(ctx context.Context, c *app.RequestContext) {
	client := pb.NewNoticeServiceClient(conn)
	req := &pb.NoticeRequest{}
	res, err := client.ShowAllNotice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"msg": err.Error()})
	}
	type SimpleNotices struct {
		Id         int    `json:"id"`
		Title      string `json:"title"`
		Createtime string `json:"createtime"`
		Updatetime string `json:"updatetime"`
	}
	var simpleNotices []*SimpleNotices
	for _, detail := range res.NoticeDetail {
		simpleNotices = append(simpleNotices, &SimpleNotices{
			Id:         int(detail.Id),
			Title:      detail.Title,
			Createtime: detail.Createtime[:10],
			Updatetime: detail.Updatetime[:10],
		})
	}
	c.JSON(http.StatusOK, utils.H{"data": simpleNotices})

}

func UpdateNotice(ctx context.Context, c *app.RequestContext) {
	id_s, _ := c.GetPostForm("id")
	id, _ := strconv.ParseUint(id_s, 10, 64)
	title, _ := c.GetPostForm("title")
	content, _ := c.GetPostForm("content")
	pubuser, _ := c.GetPostForm("pubuser")

	client := pb.NewNoticeServiceClient(conn)

	req := &pb.NoticeRequest{Id: id, Title: title, Content: content, Pubuser: pubuser}
	res, err := client.UpdateNotice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"更新失败": err.Error()})
		log.Println(err.Error())

	} else {
		c.JSON(http.StatusOK, utils.H{
			"msg": res.Msg,
		})
	}
}

func DeleteNotice(ctx context.Context, c *app.RequestContext) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	client := pb.NewNoticeServiceClient(conn)
	req := &pb.NoticeRequest{Id: id}
	res, err := client.DeleteNotice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, utils.H{"msg": res.Msg})
}
