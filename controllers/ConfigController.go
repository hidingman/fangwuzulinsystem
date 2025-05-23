package controllers

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
)

func init() {
	orm.RegisterModel(new(models.Config))
}

// ConfigControllerPage 分页接口（后端）
func ConfigControllerPage(c *gin.Context) {
	o := orm.NewOrm()
	var config []models.Config

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10 // 默认每页 10 条
	}
	sort := c.Query("sort")
	if sort == "" {
		sort = "id"
	}
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sort
	} else {
		order = sort
	}

	name := c.Query("name")
	cond := orm.NewCondition()
	if name != "" {
		cond = cond.And("name__contains", strings.Trim(name, "%"))
	}

	total, err := o.QueryTable("config").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: nil,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	_, err = o.QueryTable("config").SetCond(cond).OrderBy(order).Limit(limit, start).All(&config)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      config,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	})
}

// ConfigControllerList 分页接口（前端）
func ConfigControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var config []models.Config

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10 // 默认每页 10 条
	}
	sort := c.Query("sort")
	if sort == "" {
		sort = "id"
	}
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sort
	} else {
		order = sort
	}

	name := c.Query("name")
	cond := orm.NewCondition()
	if name != "" {
		cond = cond.And("name__contains", name)
	}

	total, err := o.QueryTable("config").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: nil,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	_, err = o.QueryTable("config").SetCond(cond).OrderBy(order).Limit(limit, start).All(&config)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      config,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	})
}

// ConfigControllerUpdate 更新接口
func ConfigControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	config := models.Config{}
	err := json.NewDecoder(c.Request.Body).Decode(&config)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: nil,
		})
		return
	}
	_, err = o.Update(&config)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "更新失败",
			Data: err.Error(),
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: nil,
	})
}

// ConfigControllerInfoid 详情接口
func ConfigControllerInfoid(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "无效的 ID 参数",
			Data: nil,
		})
		return
	}
	config := models.Config{Id: id}
	err = o.Read(&config)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: config,
	})
}

// ConfigControllerInfo 配置获取接口（后端）
func ConfigControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	name := c.Query("name")
	config := models.Config{Name: name}
	err := o.Read(&config, "Name")
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: config,
	})
}

// ConfigControllerDetail 配置获取接口（前端）
func ConfigControllerDetail(c *gin.Context) {
	o := orm.NewOrm()
	name := c.Query("name")
	config := models.Config{Name: name}
	err := o.Read(&config, "Name")
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: config,
	})
}
