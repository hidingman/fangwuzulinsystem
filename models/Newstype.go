package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 公告资讯分类
type Newstype struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 分类名称
	Typename string `json:"typename" orm:"column(typename)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
