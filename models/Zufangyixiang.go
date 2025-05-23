package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 租房意向
type Zufangyixiang struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 意向标题
	Yixiangbiaoti string `json:"yixiangbiaoti" orm:"column(yixiangbiaoti)"`
	// 意向图片
	Yixiangtupian string `json:"yixiangtupian" orm:"column(yixiangtupian)"`
	// 租房地点
	Zufangdidian string `json:"zufangdidian" orm:"column(zufangdidian)"`
	// 理想价位
	Lixiangjiawei string `json:"lixiangjiawei" orm:"column(lixiangjiawei)"`
	// 期望户型
	Qiwanghuxing string `json:"qiwanghuxing" orm:"column(qiwanghuxing)"`
	// 登记时间
	Dengjishijian lib.Time `json:"dengjishijian" orm:"auto_now_add;type(datetime);null;column(dengjishijian)"`
	// 意向详情
	Yixiangxiangqing string `json:"yixiangxiangqing" orm:"column(yixiangxiangqing)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
