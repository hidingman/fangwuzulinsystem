package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 智能客服
type Chat struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 用户id
	Userid int `json:"userid" orm:"column(userid)"`
	// 管理员id
	Adminid int `json:"adminid" orm:"column(adminid)"`
	// 提问
	Ask string `json:"ask" orm:"column(ask)"`
	// 回复
	Reply string `json:"reply" orm:"column(reply)"`
	// 是否回复
	Isreply int `json:"isreply" orm:"column(isreply)"`
	// 已读/未读(1:已读,0:未读)
	Isread int `json:"isread" orm:"column(isread)"`
	// 用户头像
	Uname string `json:"uname" orm:"column(uname)"`
	// 用户名
	Uimage string `json:"uimage" orm:"column(uimage)"`
	// 内容类型(1:文本,2:图片,3:视频,4:文件,5:表情)
	Type int `json:"type" orm:"column(type)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
