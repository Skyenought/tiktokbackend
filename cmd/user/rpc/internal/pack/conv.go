package pack

import (
	"fmt"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/internal/model"
	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
)

// UserConvProto converts model.User to userservice.User
func UserConvProto(m model.User, isSub bool) userservice.User {
	return userservice.User{
		ID:              int64(m.ID),
		Name:            m.Username,
		FollowCount:     m.FollowCount,
		FollowerCount:   m.FanCount,
		Avatar:          m.Avatar,
		IsFollow:        isSub,
		Signature:       m.Signature,
		TotalFavorited:  fmt.Sprintf("%d", m.FavoriteCount),
		BackgroundImage: m.BackgroundImage,
	}
}

// UsersConvProto converts []model.User to []*userservice.User
func UsersConvProto(ms []model.User, isSub bool) []*userservice.User {
	var ps []*userservice.User
	for _, m := range ms {
		user := UserConvProto(m, isSub)
		ps = append(ps, &user)
	}
	return ps
}


// MsgConvProto converts model.Msg to userservice.Msg
func MsgConvProto(m model.Msg) userservice.Msg {
	return userservice.Msg{
		FromUserID: int64(m.FromUserID),
		ToUserID:   int64(m.ToUserID),
		Content:    m.Content,
	}
}

// MsgsConvProto converts []model.Msg to []*userservice.Msg
func MsgsConvProto(ms []model.Msg) []*userservice.Msg {
	var ps []*userservice.Msg
	for _, m := range ms {
		msg := MsgConvProto(m)
		ps = append(ps, &msg)
	}
	return ps
}
