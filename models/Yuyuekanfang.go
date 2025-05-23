package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 预约看房
type Yuyuekanfang struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 预约时间
	Yuyueshijian lib.Time `json:"yuyueshijian" orm:"auto_now_add;type(datetime);null;column(yuyueshijian)"`
	// 预约详情
	Yuyuexiangqing string `json:"yuyuexiangqing" orm:"column(yuyuexiangqing)"`
	// 预约内容
	Yuyueneirong string `json:"yuyueneirong" orm:"column(yuyueneirong)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	// 手机
	Shouji string `json:"shouji" orm:"column(shouji)"`
	// 是否审核
	Sfsh string `json:"sfsh" orm:"column(sfsh)"`
	// 审核回复
	Shhf string `json:"shhf" orm:"column(shhf)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
