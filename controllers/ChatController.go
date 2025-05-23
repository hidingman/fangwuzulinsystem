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
	orm.RegisterModel(new(models.Chat))
}

// buildCondition 构建查询条件
func chatbuildCondition(c *gin.Context) *orm.Condition {
	cond := orm.NewCondition()
	columns := []string{
		"userid", "adminid", "ask", "reply", "isreply", "isread", "uname", "uimage", "type",
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
func ChatControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var chat []models.Chat

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
	order = "addtime"

	cond := chatbuildCondition(c)
	hasUserId := "是"
	if hasUserId == "是" && userInfo.Role != "管理员" {
		cond = cond.And("userid", userInfo.Id)
	}

	total, _ := o.QueryTable("chat").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sortParam, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("chat").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&chat)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      chat,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}

	userid, iderr := strconv.Atoi(c.Query("userid"))
	if iderr == nil {
		o.Raw("UPDATE chat SET isread = 1 where userid = ?", userid).Exec()
	}

	c.JSON(200, res)
}

// Lists 分页接口（前端）
func ChatControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var chat []models.Chat
	cond := orm.NewCondition()
	o.QueryTable("chat").SetCond(cond).All(&chat)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: chat,
	}
	c.JSON(200, res)
}

// List 分页接口（前端）
func ChatControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var chat []models.Chat

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

	cond := chatbuildCondition(c)
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	hasUserId := "是"
	if hasUserId == "是" && userInfo.Role != "管理员" {
		cond = cond.And("userid", userInfo.Id)
	}

	total, _ := o.QueryTable("chat").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("chat").SetCond(cond).OrderBy(order).Limit(limit, start).All(&chat)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      chat,
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
func ChatControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	if userInfo.Tablename == "users" {
		chat.Adminid = userInfo.Id
	} else {
		if chat.Userid == 0 {
			chat.Userid = userInfo.Id
		}
	}
	chat.Isreply = 1
	if chat.Ask != "" {
		o.Raw("UPDATE chat SET isreply = 0 where userid = ?", chat.Userid).Exec()
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&chat)
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
func ChatControllerAdd(c *gin.Context) {
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
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	if chat.Userid == 0 {
		chat.Userid = userInfo.Id
	}
	chat.Isreply = 1
	o.Raw("UPDATE chat SET isreply = 0 where userid = ?", userInfo.Id).Exec()
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&chat)
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
func ChatControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	_, err := o.Update(&chat)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func ChatControllerDelete(c *gin.Context) {
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
	o.QueryTable("chat").Filter("id__in", ids).Delete()
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// Info 详情接口（后端）
func ChatControllerInfo(c *gin.Context) {
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
	chat := models.Chat{Id: id}
	err = o.Read(&chat)

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
			Data: chat,
		})
	}
}

// Query 查询单条数据
func ChatControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := chatbuildCondition(c)
	var chat models.Chat
	o.QueryTable("chat").SetCond(cond).One(&chat)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: chat,
	})
}

// Detail 详情接口（前端）
func ChatControllerDetail(c *gin.Context) {
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
	chat := models.Chat{Id: id}
	err = o.Read(&chat)

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
			Data: chat,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func ChatControllerRemind(c *gin.Context) {
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
			sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT * FROM chat WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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

// Group 分组统计接口
func ChatControllerGroup(c *gin.Context) {
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
	sql := "SELECT COUNT(*) AS total, " + columnName + " FROM chat WHERE 1 = 1 GROUP BY " + columnName
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
func ChatControllerValue(c *gin.Context) {
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
	if "chat" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	sql := "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM chat " + where + " GROUP BY " + xColumnName
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
func ChatControllerValueTime(c *gin.Context) {
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
	if "chat" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	timeStatType, _ = url.QueryUnescape(timeStatType)

	var sql string
	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM chat " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM chat " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM chat " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
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
func ChatControllerValueMul(c *gin.Context) {
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
	if "chat" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql := "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM chat " + where + " GROUP BY " + xColumnName + " LIMIT 10"
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
