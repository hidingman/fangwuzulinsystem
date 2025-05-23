package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 租赁评价
type Zulinpingjia struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 评价评分
	Pingjiapingfen string `json:"pingjiapingfen" orm:"column(pingjiapingfen)"`
	// 公司评价
	Gongsipingjia string `json:"gongsipingjia" orm:"column(gongsipingjia)"`
	// 员工评价
	Yuangongpingjia string `json:"yuangongpingjia" orm:"column(yuangongpingjia)"`
	// 评价详情
	Pingjiaxiangqing string `json:"pingjiaxiangqing" orm:"column(pingjiaxiangqing)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
