package models

type Config struct {
	Id int `json:"id" pk:"auto;column(id)"`
	Name string `json:"name" orm:"column(name)"`
	Value string `json:"value" orm:"column(value)"`
}
