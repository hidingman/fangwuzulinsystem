package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 租赁合同
type Zulinhetong struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 押金
	Yajin float64 `json:"yajin" orm:"column(yajin)"`
	// 积分
	Jifen int `json:"jifen" orm:"column(jifen)"`
	// 合同编号
	Hetongbianhao string `json:"hetongbianhao" orm:"column(hetongbianhao)"`
	// 合同附件
	Hetongfujian string `json:"hetongfujian" orm:"column(hetongfujian)"`
	// 签订时间
	Qiandingshijian lib.Time `json:"qiandingshijian" orm:"auto_now_add;type(datetime);null;column(qiandingshijian)"`
	// 合同详情
	Hetongxiangqing string `json:"hetongxiangqing" orm:"column(hetongxiangqing)"`
	// 合同事项
	Hetongshixiang string `json:"hetongshixiang" orm:"column(hetongshixiang)"`
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
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
