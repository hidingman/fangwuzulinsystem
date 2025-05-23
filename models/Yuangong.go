package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 员工
type Yuangong struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 密码
	Mima string `json:"mima" orm:"column(mima)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 性别
	Xingbie string `json:"xingbie" orm:"column(xingbie)"`
	// 头像
	Touxiang string `json:"touxiang" orm:"column(touxiang)"`
	// 身份证
	Shenfenzheng string `json:"shenfenzheng" orm:"column(shenfenzheng)"`
	// 电话
	Dianhua string `json:"dianhua" orm:"column(dianhua)"`
	// 民族
	Minzu string `json:"minzu" orm:"column(minzu)"`
	// 籍贯
	Jiguan string `json:"jiguan" orm:"column(jiguan)"`
	// 出生日月
	Chushengriyue lib.Date `json:"chushengriyue" orm:"type(date);null;column(chushengriyue)"`
	// 年龄
	Nianling int `json:"nianling" orm:"column(nianling)"`
	// 学历
	Xueli string `json:"xueli" orm:"column(xueli)"`
	// 住址
	Zhuzhi string `json:"zhuzhi" orm:"column(zhuzhi)"`
	// 特长
	Tezhang string `json:"tezhang" orm:"column(tezhang)"`
	// 自我评价
	Ziwopingjia string `json:"ziwopingjia" orm:"column(ziwopingjia)"`
	// 工作安排
	Gongzuoanpai string `json:"gongzuoanpai" orm:"column(gongzuoanpai)"`
	// 工资信息
	Gongzixinxi string `json:"gongzixinxi" orm:"column(gongzixinxi)"`
	// 积分
	Jifen int `json:"jifen" orm:"column(jifen)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
