package pack

import (
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/commentrpc"
	"github.com/Skyenought/tiktokbackend/cmd/comment/rpc/internal/model"
)

func CommentConvProto(m model.Comment) *commentrpc.Comment {
	return &commentrpc.Comment{
		ID:         int64(m.ID),
		UserID:     int64(m.UserID),
		Content:    m.Content,
		CreateDate: m.CreateTime.Format("01-02"),
	}
}

// Comments conv to proto
func CommentsConvProto(m []model.Comment) []*commentrpc.Comment {
	var comments []*commentrpc.Comment
	for _, v := range m {
		comments = append(comments, CommentConvProto(v))
	}
	return comments
}
