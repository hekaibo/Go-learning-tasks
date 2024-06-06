package service

import (
	"golang.org/x/net/context"
	"log"
	"mission/task5/dao"
	"mission/task5/dao/dbop"
	"mission/task5/dao/model"
	pb "mission/task5/pb"
)

type noticeserver struct {
	pb.UnimplementedNoticeServiceServer
}

func NewNoticeServer() pb.NoticeServiceServer {
	return &noticeserver{}
}

func (s *noticeserver) CreateNotice(ctx context.Context, in *pb.NoticeRequest) (*pb.NoticeReply, error) {
	title := in.GetTitle()
	content := in.GetContent()
	pubuser := in.GetPubuser()

	err := dbop.CreateNotice(dao.DB, &model.Notice{
		Title:       title,
		Content:     content,
		PublishUser: pubuser,
	})
	if err != nil {
		return &pb.NoticeReply{Msg: "创建公告失败"}, err
	}
	return &pb.NoticeReply{Msg: "用户 " + pubuser + " 创建公告《" + title + "》成功"}, nil
}

func (s *noticeserver) ShowNotice(ctx context.Context, in *pb.NoticeRequest) (*pb.NoticeDetailReply, error) {

	id := in.GetId()
	notices, err := dbop.ShowNotice(dao.DB, id)
	if err != nil {
		return &pb.NoticeDetailReply{Msg: err.Error()}, err
	}
	var noticerequest []*pb.NoticeModel
	if len(notices) == 0 {
		return &pb.NoticeDetailReply{NoticeDetail: noticerequest, Msg: "未找到公告"}, err
	}
	noticerequest = append(noticerequest, &pb.NoticeModel{
		Id:         uint64(notices[0].Model.ID),
		Title:      notices[0].Title,
		Content:    notices[0].Content,
		Pubuser:    notices[0].PublishUser,
		Createtime: notices[0].Model.CreatedAt.String()[:19],
		Updatetime: notices[0].Model.UpdatedAt.String()[:19],
	})
	return &pb.NoticeDetailReply{NoticeDetail: noticerequest}, nil
}

func (s *noticeserver) ShowAllNotice(ctx context.Context, in *pb.NoticeRequest) (*pb.NoticeDetailReply, error) {
	notices, err := dbop.ShowAllNotice(dao.DB)
	if err != nil {
		log.Println(err)
		return &pb.NoticeDetailReply{Msg: "查询失败"}, err
	}
	var noticerequest []*pb.NoticeModel
	for _, notice := range notices {
		noticerequest = append(noticerequest, &pb.NoticeModel{
			Id:         uint64(notice.Model.ID),
			Title:      notice.Title,
			Content:    notice.Content,
			Pubuser:    notice.PublishUser,
			Createtime: notices[0].Model.CreatedAt.String()[:19],
			Updatetime: notices[0].Model.UpdatedAt.String()[:19],
		})
	}
	return &pb.NoticeDetailReply{
		NoticeDetail: noticerequest,
	}, nil
}

func (s *noticeserver) UpdateNotice(ctx context.Context, in *pb.NoticeRequest) (*pb.NoticeReply, error) {
	id := in.GetId()
	title := in.GetTitle()
	content := in.GetContent()
	pubuser := in.GetPubuser()

	err := dbop.UpdateNotice(dao.DB, &model.Notice{
		Title:       title,
		Content:     content,
		PublishUser: pubuser,
	}, id)
	if err != nil {
		return &pb.NoticeReply{Msg: "更新公告失败,未找到公告"}, err
	}
	return &pb.NoticeReply{Msg: "用户 " + pubuser + " 更新公告《" + title + "》成功"}, nil
}

func (s *noticeserver) DeleteNotice(ctx context.Context, in *pb.NoticeRequest) (*pb.NoticeReply, error) {
	id := in.GetId()
	notice, err := dbop.ShowNotice(dao.DB, id)
	if len(notice) == 0 {
		return &pb.NoticeReply{Msg: "未找到公告"}, err
	}
	if err != nil {
		return &pb.NoticeReply{Msg: err.Error()}, err
	}
	err = dbop.DeleteNotice(dao.DB, id)
	if err != nil {
		return &pb.NoticeReply{Msg: "删除公告失败"}, err
	}
	title := notice[0].Title
	return &pb.NoticeReply{Msg: "删除公告《" + title + "》成功"}, nil
}
