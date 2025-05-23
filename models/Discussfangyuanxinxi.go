package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 房源信息评论表
type Discussfangyuanxinxi struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 关联表id
	Refid int `json:"refid" orm:"column(refid)"`
	// 用户id
	Userid int `json:"userid" orm:"column(userid)"`
	// 头像
	Avatarurl string `json:"avatarurl" orm:"column(avatarurl)"`
	// 用户名
	Nickname string `json:"nickname" orm:"column(nickname)"`
	// 评论内容
	Content string `json:"content" orm:"column(content)"`
	// 回复内容
	Reply string `json:"reply" orm:"column(reply)"`
	// 赞
	Thumbsupnum int `json:"thumbsupnum" orm:"column(thumbsupnum)"`
	// 踩
	Crazilynum int `json:"crazilynum" orm:"column(crazilynum)"`
	// 置顶(1:置顶,0:非置顶)
	Istop int `json:"istop" orm:"column(istop)"`
	// 赞用户ids
	Tuserids string `json:"tuserids" orm:"column(tuserids)"`
	// 踩用户ids
	Cuserids string `json:"cuserids" orm:"column(cuserids)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
