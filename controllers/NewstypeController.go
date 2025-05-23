package controllers

import (
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
	"net/url"
)

func init() {
	orm.RegisterModel(new(models.Newstype))
}

// Page 分页接口（后端）
func NewstypeControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var newstype []models.Newstype

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := orm.NewCondition()
	typenameColumn := c.Query("typename")
	if typenameColumn != "" {
		if strings.Contains(typenameColumn, "%") {
			cond = cond.And("typename__contains", strings.Trim(typenameColumn, "%"))
		} else {
			cond = cond.And("typename", typenameColumn)
		}
	}

	total, err := o.QueryTable("newstype").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: err,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sort, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	_, err = o.QueryTable("newstype").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&newstype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      newstype,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(http.StatusOK, res)
}

// Lists 分页接口（前端）
func NewstypeControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var newstype []models.Newstype
	cond := orm.NewCondition()
	_, err := o.QueryTable("newstype").SetCond(cond).All(&newstype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: newstype,
	}
	c.JSON(http.StatusOK, res)
}

// List 分页接口（前端）
func NewstypeControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var newstype []models.Newstype

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sort
	} else {
		order = sort
	}

	cond := orm.NewCondition()
	typenameColumn := c.Query("typename")
	if typenameColumn != "" {
		if strings.Contains(typenameColumn, "%") {
			cond = cond.And("typename__contains", strings.Trim(typenameColumn, "%"))
		} else {
			cond = cond.And("typename", typenameColumn)
		}
	}

	total, err := o.QueryTable("newstype").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: err,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	_, err = o.QueryTable("newstype").SetCond(cond).OrderBy(order).Limit(limit, start).All(&newstype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      newstype,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}

	c.JSON(http.StatusOK, res)
}

// Save 保存接口（后端）
func NewstypeControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var newstype models.Newstype
	err = c.ShouldBindJSON(&newstype)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&newstype)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// Add 保存接口（前端）
func NewstypeControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var newstype models.Newstype
	err := c.ShouldBindJSON(&newstype)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&newstype)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// Update 更新接口
func NewstypeControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var newstype models.Newstype
	err := c.ShouldBindJSON(&newstype)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	_, err = o.Update(&newstype)

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func NewstypeControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	_, err = o.QueryTable("newstype").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败",
			Data: err,
		})
		return
	}
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// Info 详情接口（后端）
func NewstypeControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误",
			Data: err,
		})
		return
	}
	newstype := models.Newstype{Id: id}
	err = o.Read(&newstype)

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: newstype,
		})
	}
}

// Query 查询单条数据
func NewstypeControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	typenameColumn := c.Query("typename")
	if typenameColumn != "" {
		if strings.Contains(typenameColumn, "%") {
			cond = cond.And("typename__contains", strings.Trim(typenameColumn, "%"))
		} else {
			cond = cond.And("typename", typenameColumn)
		}
	}
	var newstype models.Newstype
	err := o.QueryTable("newstype").SetCond(cond).One(&newstype)
	if err != nil && err != orm.ErrNoRows {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: newstype,
	})
}

// Detail 详情接口（前端）
func NewstypeControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误",
			Data: err,
		})
		return
	}
	newstype := models.Newstype{Id: id}
	err = o.Read(&newstype)

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: newstype,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func NewstypeControllerRemind(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	where := " 1=1 "

	sql := "SELECT 0 AS count"
	columnName := c.Param("columnName")
	typeP := c.Param("type")
	remindstart := c.Query("remindstart")
	remindend := c.Query("remindend")
	if typeP == "1" {
		if remindstart != "" {
			sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM newstype WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
		}
	}

	o := orm.NewOrm()
	var result []string
	_, err := o.Raw(sql).QueryRows(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	res := struct {
		Code  int         `json:"code"`
		Count int         `json:"count"`
		Data  interface{} `json:"data"`
	}{
		0,
		len(result),
		result,
	}
	c.JSON(http.StatusOK, res)
}

// Group 分组统计接口
func NewstypeControllerGroup(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	columnName := c.Param("columnName")
	sql := ""
	where := " WHERE 1 = 1 "

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM newstype " + where + " GROUP BY " + columnName
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")
	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI < totalJ
		})
	}
	if limit > 0 && limit < len(maps) {
		maps = maps[:limit]
	}
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// Value 统计指定字段
func NewstypeControllerValue(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	xColumnName := c.Param("xColumnName")
	yColumnName := c.Param("yColumnName")
	sql := ""
	where := " WHERE 1 = 1 "

	if "newstype" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM newstype " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")
	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI < totalJ
		})
	}
	if limit > 0 && limit < len(maps) {
		maps = maps[:limit]
	}
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueTime 按日期统计
func NewstypeControllerValueTime(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	xColumnName := c.Param("xColumnName")
	yColumnName := c.Param("yColumnName")
	timeStatType := c.Param("timeStatType")
	sql := ""
	where := " WHERE 1 = 1 "

	if "newstype" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM newstype " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM newstype " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM newstype " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")
	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI < totalJ
		})
	}
	if limit > 0 && limit < len(maps) {
		maps = maps[:limit]
	}
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueMul 统计指定字段(多)
func NewstypeControllerValueMul(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	xColumnName := c.Param("xColumnName")
	yColumnNameMul := c.Query("yColumnNameMul")

	sql := ""
	where := " WHERE 1 = 1 "
	if "newstype" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM newstype " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		_, err := o.Raw(sql).Values(&maps)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ReturnMsg{
				Code: 500,
				Msg:  "查询失败",
				Data: err,
			})
			return
		}
		result = append(result, maps)
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	})
}
