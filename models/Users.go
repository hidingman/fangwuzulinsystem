package models

import (
	_ "time"
	"go-schema/utils/lib"
)

type Users struct {
	Id       int `json:"id" pk:"auto;column(id)"`
	Username string `json:"username" orm:"column(username)"`
	Password string `json:"password" orm:"column(password)"`
	Image string `json:"image" orm:"column(image)"`
	Role string `json:"role" orm:"column(role)"`
	Addtime lib.Time `json:"addtime" orm:"auto_now_add;type(datetime);null;column(addtime)"`
}
