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
	orm.RegisterModel(new(models.Zulinpingjia))
}

// buildCondition 构建查询条件
func zulinpingjiabuildCondition(c *gin.Context) *orm.Condition {
	cond := orm.NewCondition()
	fields := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "pingjiapingfen",
		"gongsipingjia", "yuangongpingjia", "pingjiaxiangqing", "gonghao",
		"yuangongxingming", "zhanghao", "xingming",
	}

	for _, field := range fields {
		value := c.Query(field)
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(field+"__contains", strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}
	return cond
}

// Page 分页接口（后端）
func ZulinpingjiaControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var zulinpingjia []models.Zulinpingjia

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := zulinpingjiabuildCondition(c)
	tableName := userInfo.Tablename

	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)

	}
	if tableName == "guke" {
		cond = cond.And("zhanghao", userInfo.Username)

	}

	total, _ := o.QueryTable("zulinpingjia").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	sortlist := strings.Split(sort, ",")
	orderlist := strings.Split(order, ",")

	for index, value := range sortlist {
		if len(orderlist) > index && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("zulinpingjia").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&zulinpingjia)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      zulinpingjia,
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
func ZulinpingjiaControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var zulinpingjia []models.Zulinpingjia
	cond := orm.NewCondition()
	o.QueryTable("zulinpingjia").SetCond(cond).All(&zulinpingjia)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: zulinpingjia,
	}
	c.JSON(http.StatusOK, res)
}

// List 分页接口（前端）
func ZulinpingjiaControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var zulinpingjia []models.Zulinpingjia

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sort
	} else {
		order = sort
	}

	cond := zulinpingjiabuildCondition(c)

	total, _ := o.QueryTable("zulinpingjia").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("zulinpingjia").SetCond(cond).OrderBy(order).Limit(limit, start).All(&zulinpingjia)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      zulinpingjia,
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
func ZulinpingjiaControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var zulinpingjia models.Zulinpingjia
	if err := c.ShouldBindJSON(&zulinpingjia); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&zulinpingjia)
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
func ZulinpingjiaControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var zulinpingjia models.Zulinpingjia
	if err := c.ShouldBindJSON(&zulinpingjia); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&zulinpingjia)
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
func ZulinpingjiaControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var zulinpingjia models.Zulinpingjia
	if err := c.ShouldBindJSON(&zulinpingjia); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	_, err := o.Update(&zulinpingjia)

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func ZulinpingjiaControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.QueryTable("zulinpingjia").Filter("id__in", ids).Delete()
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// Info 详情接口（后端）
func ZulinpingjiaControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	zulinpingjia := models.Zulinpingjia{Id: id}
	err = o.Read(&zulinpingjia)

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
			Data: zulinpingjia,
		})
	}
}

// Query 查询单条数据
func ZulinpingjiaControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := zulinpingjiabuildCondition(c)
	var zulinpingjia models.Zulinpingjia
	o.QueryTable("zulinpingjia").SetCond(cond).One(&zulinpingjia)
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: zulinpingjia,
	})
}

// Detail 详情接口（前端）
func ZulinpingjiaControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	zulinpingjia := models.Zulinpingjia{Id: id}
	err = o.Read(&zulinpingjia)

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
			Data: zulinpingjia,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func ZulinpingjiaControllerRemind(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	where := " 1=1 "
	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}
	if tableName == "guke" {
		where += " AND zhanghao = '" + userInfo.Username + "' "
	}

	sql := "SELECT 0 AS count"
	columnName := c.Param("columnName")
	typeP := c.Param("type")
	remindstart := c.Query("remindstart")
	remindend := c.Query("remindend")
	if typeP == "1" {
		if remindstart != "" {
			sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM zulinpingjia WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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
	c.JSON(http.StatusOK, res)
}

// Group 分组统计接口
func ZulinpingjiaControllerGroup(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)

	if err != nil {
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

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}
	if tableName == "guke" {
		where += " AND zhanghao = '" + userInfo.Username + "' "
	}

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM zulinpingjia " + where + " GROUP BY " + columnName
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// Value 统计指定字段
func ZulinpingjiaControllerValue(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)

	if err != nil {
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

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}
	if tableName == "guke" {
		where += " AND zhanghao = '" + userInfo.Username + "' "
	}

	if "zulinpingjia" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM zulinpingjia " + where + " GROUP BY " + xColumnName
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueTime 按日期统计
func ZulinpingjiaControllerValueTime(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)

	if err != nil {
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

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}
	if tableName == "guke" {
		where += " AND zhanghao = '" + userInfo.Username + "' "
	}

	if "zulinpingjia" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM zulinpingjia " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM zulinpingjia " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM zulinpingjia " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueMul 统计指定字段(多)
func ZulinpingjiaControllerValueMul(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, err := utils.ValidateToken(token)

	if err != nil {
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
	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}
	if tableName == "guke" {
		where += " AND zhanghao = '" + userInfo.Username + "' "
	}
	if "zulinpingjia" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM zulinpingjia " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		o.Raw(sql).Values(&maps)
		result = append(result, maps)
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	})
}
