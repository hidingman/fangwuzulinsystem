package models

type ReturnMsg struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

type ReturnToken struct {
	Code  int `json:"code"`
	Token string `json:"token"`
}

type ReturnPage struct {
	CurrPage int `json:"currPage"`
	List interface{} `json:"list"`
	PageSize int `json:"pageSize"`
	Total int `json:"total"`
	TotalPage int `json:"totalPage"`
}

type ReturnFile struct {
	Code  int `json:"code"`
	File string `json:"file"`
}

type ReturnCount struct {
	Code  int `json:"code"`
	Count int `json:"count"`
}

type ReturnScore struct {
	Code  int `json:"code"`
	Score interface{} `json:"score"`
}
