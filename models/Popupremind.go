package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 弹窗提醒
type Popupremind struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 发布人id
	Userid int `json:"userid" orm:"column(userid)"`
	// 标题
	Title string `json:"title" orm:"column(title)"`
	// 类型
	Type string `json:"type" orm:"column(type)"`
	// 简介
	Brief string `json:"brief" orm:"column(brief)"`
	// 内容
	Content string `json:"content" orm:"column(content)"`
	// 提醒时间
	Remindtime lib.Time `json:"remindtime" orm:"auto_now_add;type(datetime);null;column(remindtime)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
