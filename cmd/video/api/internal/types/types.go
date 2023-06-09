// Code generated by goctl. DO NOT EDIT.
package types

type PublishActionReq struct {
	Token string `form:"token"`
	Title string `form:"title"`
}

type PublishActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishlistReq struct {
	Token  string `form:"token"`
	UserID int64  `form:"user_id"`
}

type PublishlistResp struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type FeedReq struct {
	LatestTime int64  `form:"latest_time,optional"`
	Token      string `form:"token,optional"`
}

type FeedResp struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	NextTime   int64   `json:"next_time"`
	VideoList  []Video `json:"video_list"`
}

type FavoriteActionReq struct {
	UserID     int64  `form:"user_id"`
	Token      string `form:"token"`
	VideoID    int64  `form:"video_id"`
	ActionType int32  `form:"action_type,options=1|2"`
}

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteVideoListReq struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type FavoriteVideoListResp struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int32  `json:"follow_count"`
	FollowerCount   int32  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int32  `json:"work_count"`
	FavoriteCount   int32  `json:"favorite_count"`
}

type Video struct {
	ID            int64  `json:"id,optional"`
	Author        User   `json:"author,optional"`
	CoverUrl      string `json:"cover_url,optional"`
	PlayUrl       string `json:"play_url,optional"`
	FavoriteCount int32  `json:"favorite_count,optional"`
	CommentCount  int32  `json:"comment_count,optional"`
	IsFavorite    bool   `json:"is_favorite,optional"`
	Title         string `json:"title,optional"`
}

type Comment struct {
	ID         int64  `json:"id,optional"`
	User       User   `json:"user,optional"`
	Content    string `json:"content,optional"`
	CreateDate string `json:"create_date,optional"`
}

type Msg struct {
	ID         int64  `json:"id"`
	Content    string `json:"content"`
	ToUserID   int64  `json:"to_user_id"`
	FromUserID int64  `json:"from_user_id"`
	CreateTime int64  `json:"create_time,omitempty"`
}
