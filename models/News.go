package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 公告资讯
type News struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 标题
	Title string `json:"title" orm:"column(title)"`
	// 简介
	Introduction string `json:"introduction" orm:"column(introduction)"`
	// 分类名称
	Typename string `json:"typename" orm:"column(typename)"`
	// 发布人
	Name string `json:"name" orm:"column(name)"`
	// 头像
	Headportrait string `json:"headportrait" orm:"column(headportrait)"`
	// 点击次数
	Clicknum int `json:"clicknum" orm:"column(clicknum)"`
	// 最近点击时间
	Clicktime lib.Time `json:"clicktime" orm:"auto_now_add;type(datetime);null;column(clicktime)"`
	// 赞
	Thumbsupnum int `json:"thumbsupnum" orm:"column(thumbsupnum)"`
	// 踩
	Crazilynum int `json:"crazilynum" orm:"column(crazilynum)"`
	// 收藏数
	Storeupnum int `json:"storeupnum" orm:"column(storeupnum)"`
	// 图片
	Picture string `json:"picture" orm:"column(picture)"`
	// 内容
	Content string `json:"content" orm:"column(content)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
