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
	orm.RegisterModel(new(models.Storeup))
}

// Page 分页接口（后端）
func StoreupControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var storeup []models.Storeup

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
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	tablenameColumn := c.Query("tablename")
	if tablenameColumn != "" {
		if strings.Contains(tablenameColumn, "%") {
			cond = cond.And("tablename__contains", strings.Trim(tablenameColumn, "%"))
		} else {
			cond = cond.And("tablename", tablenameColumn)
		}
	}
	nameColumn := c.Query("name")
	if nameColumn != "" {
		if strings.Contains(nameColumn, "%") {
			cond = cond.And("name__contains", strings.Trim(nameColumn, "%"))
		} else {
			cond = cond.And("name", nameColumn)
		}
	}
	pictureColumn := c.Query("picture")
	if pictureColumn != "" {
		if strings.Contains(pictureColumn, "%") {
			cond = cond.And("picture__contains", strings.Trim(pictureColumn, "%"))
		} else {
			cond = cond.And("picture", pictureColumn)
		}
	}
	typeColumn := c.Query("type")
	if typeColumn != "" {
		if strings.Contains(typeColumn, "%") {
			cond = cond.And("type__contains", strings.Trim(typeColumn, "%"))
		} else {
			cond = cond.And("type", typeColumn)
		}
	}
	inteltypeColumn := c.Query("inteltype")
	if inteltypeColumn != "" {
		if strings.Contains(inteltypeColumn, "%") {
			cond = cond.And("inteltype__contains", strings.Trim(inteltypeColumn, "%"))
		} else {
			cond = cond.And("inteltype", inteltypeColumn)
		}
	}
	remarkColumn := c.Query("remark")
	if remarkColumn != "" {
		if strings.Contains(remarkColumn, "%") {
			cond = cond.And("remark__contains", strings.Trim(remarkColumn, "%"))
		} else {
			cond = cond.And("remark", remarkColumn)
		}
	}
	hasUserId := "是"

	if hasUserId == "是" && userInfo.Role != "管理员" {
		cond = cond.And("userid", userInfo.Id)
	}

	total, err := o.QueryTable("storeup").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
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
	_, err = o.QueryTable("storeup").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&storeup)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      storeup,
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
func StoreupControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var storeup []models.Storeup
	cond := orm.NewCondition()
	_, err := o.QueryTable("storeup").SetCond(cond).All(&storeup)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: storeup,
	}
	c.JSON(200, res)
}

// List 分页接口（前端）
func StoreupControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var storeup []models.Storeup

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
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	tablenameColumn := c.Query("tablename")
	if tablenameColumn != "" {
		if strings.Contains(tablenameColumn, "%") {
			cond = cond.And("tablename__contains", strings.Trim(tablenameColumn, "%"))
		} else {
			cond = cond.And("tablename", tablenameColumn)
		}
	}
	nameColumn := c.Query("name")
	if nameColumn != "" {
		if strings.Contains(nameColumn, "%") {
			cond = cond.And("name__contains", strings.Trim(nameColumn, "%"))
		} else {
			cond = cond.And("name", nameColumn)
		}
	}
	pictureColumn := c.Query("picture")
	if pictureColumn != "" {
		if strings.Contains(pictureColumn, "%") {
			cond = cond.And("picture__contains", strings.Trim(pictureColumn, "%"))
		} else {
			cond = cond.And("picture", pictureColumn)
		}
	}
	typeColumn := c.Query("type")
	if typeColumn != "" {
		if strings.Contains(typeColumn, "%") {
			cond = cond.And("type__contains", strings.Trim(typeColumn, "%"))
		} else {
			cond = cond.And("type", typeColumn)
		}
	}
	inteltypeColumn := c.Query("inteltype")
	if inteltypeColumn != "" {
		if strings.Contains(inteltypeColumn, "%") {
			cond = cond.And("inteltype__contains", strings.Trim(inteltypeColumn, "%"))
		} else {
			cond = cond.And("inteltype", inteltypeColumn)
		}
	}
	remarkColumn := c.Query("remark")
	if remarkColumn != "" {
		if strings.Contains(remarkColumn, "%") {
			cond = cond.And("remark__contains", strings.Trim(remarkColumn, "%"))
		} else {
			cond = cond.And("remark", remarkColumn)
		}
	}

	total, err := o.QueryTable("storeup").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: err,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	_, err = o.QueryTable("storeup").SetCond(cond).OrderBy(order).Limit(limit, start).All(&storeup)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取数据失败",
			Data: err,
		})
		return
	}
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      storeup,
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
func StoreupControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var storeup models.Storeup
	err := c.ShouldBindJSON(&storeup)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	storeup.Userid = userInfo.Id
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&storeup)
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
func StoreupControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您未登录，请登录后重试",
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
	var storeup models.Storeup
	err = c.ShouldBindJSON(&storeup)

	if storeup.Userid == 0 {
		storeup.Userid = userInfo.Id
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&storeup)
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
func StoreupControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var storeup models.Storeup
	err := c.ShouldBindJSON(&storeup)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	_, err = o.Update(&storeup)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func StoreupControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	_, err = o.QueryTable("storeup").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败",
			Data: err,
		})
		return
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// Info 详情接口（后端）
func StoreupControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误",
			Data: err,
		})
		return
	}
	storeup := models.Storeup{Id: id}
	err = o.Read(&storeup)

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
			Data: storeup,
		})
	}
}

// Query 查询单条数据
func StoreupControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	tablenameColumn := c.Query("tablename")
	if tablenameColumn != "" {
		if strings.Contains(tablenameColumn, "%") {
			cond = cond.And("tablename__contains", strings.Trim(tablenameColumn, "%"))
		} else {
			cond = cond.And("tablename", tablenameColumn)
		}
	}
	nameColumn := c.Query("name")
	if nameColumn != "" {
		if strings.Contains(nameColumn, "%") {
			cond = cond.And("name__contains", strings.Trim(nameColumn, "%"))
		} else {
			cond = cond.And("name", nameColumn)
		}
	}
	pictureColumn := c.Query("picture")
	if pictureColumn != "" {
		if strings.Contains(pictureColumn, "%") {
			cond = cond.And("picture__contains", strings.Trim(pictureColumn, "%"))
		} else {
			cond = cond.And("picture", pictureColumn)
		}
	}
	typeColumn := c.Query("type")
	if typeColumn != "" {
		if strings.Contains(typeColumn, "%") {
			cond = cond.And("type__contains", strings.Trim(typeColumn, "%"))
		} else {
			cond = cond.And("type", typeColumn)
		}
	}
	inteltypeColumn := c.Query("inteltype")
	if inteltypeColumn != "" {
		if strings.Contains(inteltypeColumn, "%") {
			cond = cond.And("inteltype__contains", strings.Trim(inteltypeColumn, "%"))
		} else {
			cond = cond.And("inteltype", inteltypeColumn)
		}
	}
	remarkColumn := c.Query("remark")
	if remarkColumn != "" {
		if strings.Contains(remarkColumn, "%") {
			cond = cond.And("remark__contains", strings.Trim(remarkColumn, "%"))
		} else {
			cond = cond.And("remark", remarkColumn)
		}
	}
	var storeup models.Storeup
	err := o.QueryTable("storeup").SetCond(cond).One(&storeup)
	if err != nil && err != orm.ErrNoRows {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: storeup,
	})
}

// Detail 详情接口（前端）
func StoreupControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误",
			Data: err,
		})
		return
	}
	storeup := models.Storeup{Id: id}
	err = o.Read(&storeup)

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
			Data: storeup,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func StoreupControllerRemind(c *gin.Context) {
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
			sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM storeup WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
		}
	}

	o := orm.NewOrm()
	var result []string
	_, err = o.Raw(sql).QueryRows(&result)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
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
	c.JSON(200, res)
}

// Group 分组统计接口
func StoreupControllerGroup(c *gin.Context) {
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
	sql := "SELECT COUNT(*) AS total, " + columnName + " FROM storeup WHERE 1 = 1 GROUP BY " + columnName
	var maps []orm.Params
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		limit = 0
	}
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
func StoreupControllerValue(c *gin.Context) {
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
	if "storeup" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	sql := "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM storeup " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		limit = 0
	}
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
func StoreupControllerValueTime(c *gin.Context) {
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
	if "storeup" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	timeStatType, _ = url.QueryUnescape(timeStatType)

	var sql string
	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM storeup " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM storeup " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM storeup " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		limit = 0
	}
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
func StoreupControllerValueMul(c *gin.Context) {
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
	if "storeup" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql := "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM storeup " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		_, err = o.Raw(sql).Values(&maps)
		if err != nil {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "查询失败",
				Data: err,
			})
			return
		}
		result = append(result, maps)
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	})
}
