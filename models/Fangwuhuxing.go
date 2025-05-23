package models

import (
	_ "time"
	"go-schema/utils/lib"
)

// 房屋户型
type Fangwuhuxing struct {
	Id int `json:"id" pk:"auto;column(id)"`
	// 房屋户型
	Fangwuhuxing string `json:"fangwuhuxing" orm:"column(fangwuhuxing)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
