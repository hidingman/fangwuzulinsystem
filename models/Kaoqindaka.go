package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 考勤打卡
type Kaoqindaka struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 打卡类型
	Dakaleixing string `json:"dakaleixing" orm:"column(dakaleixing)"`
	// 打卡时间
	Dakashijian lib.Time `json:"dakashijian" orm:"auto_now_add;type(datetime);null;column(dakashijian)"`
	// 打卡详情
	Dakaxiangqing string `json:"dakaxiangqing" orm:"column(dakaxiangqing)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 头像
	Touxiang string `json:"touxiang" orm:"column(touxiang)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
