package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 房屋租赁
type Fangwuzulin struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 起始时间
	Qishishijian lib.Time `json:"qishishijian" orm:"auto_now_add;type(datetime);null;column(qishishijian)"`
	// 结束时间
	Jieshushijian lib.Time `json:"jieshushijian" orm:"auto_now_add;type(datetime);null;column(jieshushijian)"`
	// 租赁时间
	Zulinshijian int `json:"zulinshijian" orm:"column(zulinshijian)"`
	// 租金价格
	Zujinjiage float64 `json:"zujinjiage" orm:"column(zujinjiage)"`
	// 总计
	Zongji float64 `json:"zongji" orm:"column(zongji)"`
	// 押金
	Yajin float64 `json:"yajin" orm:"column(yajin)"`
	// 积分
	Jifen int `json:"jifen" orm:"column(jifen)"`
	// 租赁详情
	Zulinxiangqing string `json:"zulinxiangqing" orm:"column(zulinxiangqing)"`
	// 租赁备注
	Zulinbeizhu string `json:"zulinbeizhu" orm:"column(zulinbeizhu)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	// 是否支付
	Ispay string `json:"ispay" orm:"column(ispay)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
