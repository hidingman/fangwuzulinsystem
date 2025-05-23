package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 顾客
type Guke struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 账号
	Zhanghao string `json:"zhanghao" orm:"column(zhanghao)"`
	// 密码
	Mima string `json:"mima" orm:"column(mima)"`
	// 姓名
	Xingming string `json:"xingming" orm:"column(xingming)"`
	// 性别
	Xingbie string `json:"xingbie" orm:"column(xingbie)"`
	// 手机
	Shouji string `json:"shouji" orm:"column(shouji)"`
	// 头像
	Touxiang string `json:"touxiang" orm:"column(touxiang)"`
	// 身份证
	Shenfenzheng string `json:"shenfenzheng" orm:"column(shenfenzheng)"`
	// 民族
	Minzu string `json:"minzu" orm:"column(minzu)"`
	// 籍贯
	Jiguan string `json:"jiguan" orm:"column(jiguan)"`
	// 出生日期
	Chushengriqi lib.Date `json:"chushengriqi" orm:"type(date);null;column(chushengriqi)"`
	// 年龄
	Nianling int `json:"nianling" orm:"column(nianling)"`
	// 学历
	Xueli string `json:"xueli" orm:"column(xueli)"`
	// 住址
	Zhuzhi string `json:"zhuzhi" orm:"column(zhuzhi)"`
	// 需求
	Xuqiu string `json:"xuqiu" orm:"column(xuqiu)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
