package models

type Session struct {
	Id        int `json:"id"`
	Username  string `json:"username"`
	Tablename string `json:"tablename"`
	Role      string `json:"role"`
	LoginUserColumn string `json:"loginUserColumn"`
}