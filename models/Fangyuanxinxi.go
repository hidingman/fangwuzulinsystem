package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 房源信息
type Fangyuanxinxi struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋名称
	Fangwumingcheng string `json:"fangwumingcheng" orm:"column(fangwumingcheng)"`
	// 房屋图片
	Fangwutupian string `json:"fangwutupian" orm:"column(fangwutupian)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	// 房屋地址
	Fangwudizhi string `json:"fangwudizhi" orm:"column(fangwudizhi)"`
	// 面积
	Mianji string `json:"mianji" orm:"column(mianji)"`
	// 租金价格
	Zujinjiage float64 `json:"zujinjiage" orm:"column(zujinjiage)"`
	// 押金方式
	Yajinfangshi string `json:"yajinfangshi" orm:"column(yajinfangshi)"`
	// 付款方式
	Fukuanfangshi string `json:"fukuanfangshi" orm:"column(fukuanfangshi)"`
	// 房屋朝向
	Fangwuchaoxiang string `json:"fangwuchaoxiang" orm:"column(fangwuchaoxiang)"`
	// 押金
	Yajin float64 `json:"yajin" orm:"column(yajin)"`
	// 小区
	Xiaoqu string `json:"xiaoqu" orm:"column(xiaoqu)"`
	// 房屋视频
	Fangwushipin string `json:"fangwushipin" orm:"column(fangwushipin)"`
	// 楼栋单元
	Loudongdanyuan string `json:"loudongdanyuan" orm:"column(loudongdanyuan)"`
	// 房号
	Fanghao string `json:"fanghao" orm:"column(fanghao)"`
	// 积分
	Jifen int `json:"jifen" orm:"column(jifen)"`
	// 房屋结构
	Fangwujiegou string `json:"fangwujiegou" orm:"column(fangwujiegou)"`
	// 房屋状态
	Fangwuzhuangtai string `json:"fangwuzhuangtai" orm:"column(fangwuzhuangtai)"`
	// 房产证编号
	Fangchanzhengbianhao string `json:"fangchanzhengbianhao" orm:"column(fangchanzhengbianhao)"`
	// 房产证照片
	Fangchanzhengzhaopian string `json:"fangchanzhengzhaopian" orm:"column(fangchanzhengzhaopian)"`
	// 房主姓名
	Fangzhuxingming string `json:"fangzhuxingming" orm:"column(fangzhuxingming)"`
	// 房主身份证
	Fangzhushenfenzheng string `json:"fangzhushenfenzheng" orm:"column(fangzhushenfenzheng)"`
	// 房主电话
	Fangzhudianhua string `json:"fangzhudianhua" orm:"column(fangzhudianhua)"`
	// 房屋详情
	Fangwuxiangqing string `json:"fangwuxiangqing" orm:"column(fangwuxiangqing)"`
	// 发布时间
	Fabushijian lib.Time `json:"fabushijian" orm:"auto_now_add;type(datetime);null;column(fabushijian)"`
	// 工号
	Gonghao string `json:"gonghao" orm:"column(gonghao)"`
	// 员工姓名
	Yuangongxingming string `json:"yuangongxingming" orm:"column(yuangongxingming)"`
	// 赞
	Thumbsupnum int `json:"thumbsupnum" orm:"column(thumbsupnum)"`
	// 踩
	Crazilynum int `json:"crazilynum" orm:"column(crazilynum)"`
	// 最近点击时间
	Clicktime lib.Time `json:"clicktime" orm:"auto_now_add;type(datetime);null;column(clicktime)"`
	// 点击次数
	Clicknum int `json:"clicknum" orm:"column(clicknum)"`
	// 评论数
	Discussnum int `json:"discussnum" orm:"column(discussnum)"`
	// 收藏数
	Storeupnum int `json:"storeupnum" orm:"column(storeupnum)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
