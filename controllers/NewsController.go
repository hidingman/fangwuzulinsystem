package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func init() {
	orm.RegisterModel(new(models.News))
}

// buildCondition 构建查询条件
func newsbuildCondition(c *gin.Context) *orm.Condition {
	cond := orm.NewCondition()
	columns := []string{
		"title", "introduction", "typename", "name", "headportrait", "clicknum",
		"clicktime", "thumbsupnum", "crazilynum", "storeupnum", "picture", "content",
	}

	for _, col := range columns {
		value := c.Query(col)
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(col+"__contains", strings.Trim(value, "%"))
			} else {
				cond = cond.And(col, value)
			}
		}
	}
	return cond
}

// Page 分页接口（后端）
func NewsControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var news []models.News

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sortParam := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := newsbuildCondition(c)

	total, _ := o.QueryTable("news").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sortParam, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("news").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&news)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      news,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(200, res)
}

// Lists 分页接口（前端）
func NewsControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var news []models.News
	cond := orm.NewCondition()
	o.QueryTable("news").SetCond(cond).All(&news)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: news,
	}
	c.JSON(200, res)
}

// List 分页接口（前端）
func NewsControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var news []models.News

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sortParam := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sortParam
	} else {
		order = sortParam
	}

	cond := newsbuildCondition(c)

	total, _ := o.QueryTable("news").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("news").SetCond(cond).OrderBy(order).Limit(limit, start).All(&news)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      news,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(200, res)
}

// Save 保存接口（后端）
func NewsControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&news)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(200, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// Add 保存接口（前端）
func NewsControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&news)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(200, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// Update 更新接口
func NewsControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	_, err := o.Update(&news)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func NewsControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.QueryTable("news").Filter("id__in", ids).Delete()
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// Info 详情接口（后端）
func NewsControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	news := models.News{Id: id}
	err = o.Read(&news)

	sql := "update news set clicknum = clicknum + 1 where id = " + c.Param("id")
	o.Raw(sql).Exec()

	if err != nil {
		c.JSON(200, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: news,
		})
	}
}

// Query 查询单条数据
func NewsControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := newsbuildCondition(c)
	var news models.News
	o.QueryTable("news").SetCond(cond).One(&news)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: news,
	})
}

// Detail 详情接口（前端）
func NewsControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	news := models.News{Id: id}
	err = o.Read(&news)

	sql := "update news set clicknum = clicknum + 1 where id = " + c.Param("id")
	o.Raw(sql).Exec()

	if err != nil {
		c.JSON(200, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: news,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func NewsControllerRemind(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
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
			sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM news WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
		}
	}

	o := orm.NewOrm()
	var result []string
	o.Raw(sql).QueryRows(&result)
	res := struct {
		Code  int         `json:"code"`
		Count int         `json:"count"`
		Data  interface{} `json:"data"`
	}{
		0,
		len(result),
		result,
	}
	c.JSON(200, res)
}

// Thumbsup 赞、踩接口
func NewsControllerThumbsup(c *gin.Context) {
	o := orm.NewOrm()
	typeP := c.Query("type")
	id := c.Param("id")
	sql := ""
	if typeP == "1" {
		sql = "update news set thumbsupnum = thumbsupnum + 1 where id = " + id
	}
	if typeP == "2" {
		sql = "update news set crazilynum = crazilynum + 1 where id = " + id
	}

	o.Raw(sql).Exec()

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "操作成功！",
		Data: nil,
	})
}

// AutoSort 智能推荐接口
func NewsControllerAutoSort(c *gin.Context) {
	o := orm.NewOrm()
	var news []models.News

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		limit = 5
	}
	sortParam := c.DefaultQuery("sort", "clicktime")
	order := c.Query("order")
	if order == "desc" {
		order = sortParam
	} else {
		order = "-" + sortParam
	}

	sortParam = "-clicknum"
	order = sortParam

	cond := orm.NewCondition()
	sfsh := c.Query("sfsh")
	if sfsh != "" {
		cond = cond.And("sfsh", sfsh)
	}

	total, _ := o.QueryTable("news").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("news").SetCond(cond).OrderBy(order).Limit(limit, start).All(&news)

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      news,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(200, res)
}

// AutoSort2 智能推荐接口
func NewsControllerAutoSort2(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var news []models.News
	var news2 []models.News

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		limit = 5
	}

	condStoreup := orm.NewCondition()
	condStoreup = condStoreup.And("userid", userInfo.Id)
	condStoreup = condStoreup.And("tablename", "news")
	var storeupLists []orm.ParamsList
	o.QueryTable("storeup").SetCond(condStoreup).OrderBy("-addtime").ValuesList(&storeupLists, "inteltype")

	cond := orm.NewCondition()
	if len(storeupLists) == 0 {
		cond = cond.And("typename__in", []string{"null"})
	} else {
		cond = cond.And("typename__in", storeupLists)
	}
	total, _ := o.QueryTable("news").Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("news").SetCond(cond).Limit(limit, start).All(&news)

	cond2 := orm.NewCondition()
	if len(storeupLists) == 0 {
		cond2 = cond2.AndNot("typename__in", []string{"null"})
	} else {
		cond2 = cond2.AndNot("typename__in", storeupLists)
	}
	o.QueryTable("news").SetCond(cond2).Limit(5-len(news), 0).All(&news2)

	result := make([]models.News, len(news)+len(news2))
	copy(result, news)
	copy(result[len(news):], news2)

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      result,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(200, res)
}

// Group 分组统计接口
func NewsControllerGroup(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	columnName := c.Param("columnName")
	sql := "SELECT COUNT(*) AS total, " + columnName + " FROM news WHERE 1 = 1 GROUP BY " + columnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
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
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// Value 统计指定字段
func NewsControllerValue(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	xColumnName := c.Param("xColumnName")
	yColumnName := c.Param("yColumnName")
	where := " WHERE 1 = 1 "
	if "news" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	sql := "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM news " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
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
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueTime 按日期统计
func NewsControllerValueTime(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
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
	where := " WHERE 1 = 1 "
	if "news" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	timeStatType, _ = url.QueryUnescape(timeStatType)

	var sql string
	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM news " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM news " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM news " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	o.Raw(sql).Values(&maps)
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
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueMul 统计指定字段(多)
func NewsControllerValueMul(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	xColumnName := c.Param("xColumnName")
	yColumnNameMul := c.Query("yColumnNameMul")
	where := " WHERE 1 = 1 "
	if "news" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql := "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM news " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		o.Raw(sql).Values(&maps)
		result = append(result, maps)
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	})
}
