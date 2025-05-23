package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 退租申请
type Tuizushenqing struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 押金
	Yajin float64 `json:"yajin" orm:"column(yajin)"`
	// 退租时间
	Tuizushijian lib.Time `json:"tuizushijian" orm:"auto_now_add;type(datetime);null;column(tuizushijian)"`
	// 退租原因
	Tuizuyuanyin string `json:"tuizuyuanyin" orm:"column(tuizuyuanyin)"`
	// 退租详情
	Tuizuxiangqing string `json:"tuizuxiangqing" orm:"column(tuizuxiangqing)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	// 是否审核
	Sfsh string `json:"sfsh" orm:"column(sfsh)"`
	// 审核回复
	Shhf string `json:"shhf" orm:"column(shhf)"`
	// 是否支付
	Ispay string `json:"ispay" orm:"column(ispay)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
