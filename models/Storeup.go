package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 收藏表
type Storeup struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 用户id
	Userid int `json:"userid" orm:"column(userid)"`
	// 商品id
	Refid int `json:"refid" orm:"column(refid)"`
	// 表名
	Tablename string `json:"tablename" orm:"column(tablename)"`
	// 名称
	Name string `json:"name" orm:"column(name)"`
	// 图片
	Picture string `json:"picture" orm:"column(picture)"`
	// 类型
	Type string `json:"type" orm:"column(type)"`
	// 推荐类型
	Inteltype string `json:"inteltype" orm:"column(inteltype)"`
	// 备注
	Remark string `json:"remark" orm:"column(remark)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
