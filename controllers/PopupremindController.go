package controllers

import (
	"encoding/json"
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
	orm.RegisterModel(new(models.Popupremind))
}

// Page 分页接口（后端）
func PopupremindControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var popupremind []models.Popupremind

	page, pageErr := strconv.Atoi(c.Query("page"))
	if pageErr != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	sortStr := c.Query("sort")
	if sortStr == "" {
		sortStr = "id"
	}
	order := c.Query("order")

	cond := orm.NewCondition()
	remindtimestart := c.Query("remindtimestart")
	remindtimeend := c.Query("remindtimeend")
	if remindtimestart != "" {
		cond = cond.And("remindtime__gte", remindtimestart)
	}
	if remindtimeend != "" {
		cond = cond.And("remindtime__lte", remindtimeend)
	}
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	titleColumn := c.Query("title")
	if titleColumn != "" {
		if strings.Contains(titleColumn, "%") {
			cond = cond.And("title__contains", strings.Trim(titleColumn, "%"))
		} else {
			cond = cond.And("title", titleColumn)
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
	briefColumn := c.Query("brief")
	if briefColumn != "" {
		if strings.Contains(briefColumn, "%") {
			cond = cond.And("brief__contains", strings.Trim(briefColumn, "%"))
		} else {
			cond = cond.And("brief", briefColumn)
		}
	}
	contentColumn := c.Query("content")
	if contentColumn != "" {
		if strings.Contains(contentColumn, "%") {
			cond = cond.And("content__contains", strings.Trim(contentColumn, "%"))
		} else {
			cond = cond.And("content", contentColumn)
		}
	}
	remindtimeColumn := c.Query("remindtime")
	if remindtimeColumn != "" {
		if strings.Contains(remindtimeColumn, "%") {
			cond = cond.And("remindtime__contains", strings.Trim(remindtimeColumn, "%"))
		} else {
			cond = cond.And("remindtime", remindtimeColumn)
		}
	}
	hasUserId := "是"

	if hasUserId == "是" && userInfo.Role != "管理员" {
		cond = cond.And("userid", userInfo.Id)
	}

	total, _ := o.QueryTable("popupremind").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sortStr, ",")

	for index, value := range sortlist {
		if orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("popupremind").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&popupremind)
	returnPage := models.ReturnPage{
		CurrPage:  int(page),
		List:      popupremind,
		PageSize:  int(limit),
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
func PopupremindControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var popupremind []models.Popupremind
	cond := orm.NewCondition()
	o.QueryTable("popupremind").SetCond(cond).All(&popupremind)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: popupremind,
	}
	c.JSON(200, res)
}

// List 分页接口（前端）
func PopupremindControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var popupremind []models.Popupremind

	page, pageErr := strconv.Atoi(c.Query("page"))
	if pageErr != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	sortStr := c.Query("sort")
	if sortStr == "" {
		sortStr = "id"
	}
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sortStr
	} else {
		order = sortStr
	}

	cond := orm.NewCondition()
	remindtimestart := c.Query("remindtimestart")
	remindtimeend := c.Query("remindtimeend")
	if remindtimestart != "" {
		cond = cond.And("remindtime__gte", remindtimestart)
	}
	if remindtimeend != "" {
		cond = cond.And("remindtime__lte", remindtimeend)
	}
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	titleColumn := c.Query("title")
	if titleColumn != "" {
		if strings.Contains(titleColumn, "%") {
			cond = cond.And("title__contains", strings.Trim(titleColumn, "%"))
		} else {
			cond = cond.And("title", titleColumn)
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
	briefColumn := c.Query("brief")
	if briefColumn != "" {
		if strings.Contains(briefColumn, "%") {
			cond = cond.And("brief__contains", strings.Trim(briefColumn, "%"))
		} else {
			cond = cond.And("brief", briefColumn)
		}
	}
	contentColumn := c.Query("content")
	if contentColumn != "" {
		if strings.Contains(contentColumn, "%") {
			cond = cond.And("content__contains", strings.Trim(contentColumn, "%"))
		} else {
			cond = cond.And("content", contentColumn)
		}
	}
	remindtimeColumn := c.Query("remindtime")
	if remindtimeColumn != "" {
		if strings.Contains(remindtimeColumn, "%") {
			cond = cond.And("remindtime__contains", strings.Trim(remindtimeColumn, "%"))
		} else {
			cond = cond.And("remindtime", remindtimeColumn)
		}
	}

	total, _ := o.QueryTable("popupremind").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("popupremind").SetCond(cond).OrderBy(order).Limit(limit, start).All(&popupremind)
	returnPage := models.ReturnPage{
		CurrPage:  int(page),
		List:      popupremind,
		PageSize:  int(limit),
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
func PopupremindControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var popupremind models.Popupremind
	if err := json.NewDecoder(c.Request.Body).Decode(&popupremind); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求数据解析失败",
			Data: err,
		})
		return
	}
	popupremind.Userid = userInfo.Id
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&popupremind)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(402, models.ReturnMsg{
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
func PopupremindControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var popupremind models.Popupremind
	json.NewDecoder(c.Request.Body).Decode(&popupremind)
	if popupremind.Userid == 0 {
		popupremind.Userid = userInfo.Id
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&popupremind)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(402, models.ReturnMsg{
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
func PopupremindControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var popupremind models.Popupremind
	if err := json.NewDecoder(c.Request.Body).Decode(&popupremind); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求数据解析失败",
			Data: err,
		})
		return
	}
	_, err := o.Update(&popupremind)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "更新失败！",
			Data: err,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "更新成功！",
			Data: nil,
		})
	}
}

// Delete 删除接口
func PopupremindControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	if err := json.NewDecoder(c.Request.Body).Decode(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求数据解析失败",
			Data: err,
		})
		return
	}
	_, err := o.QueryTable("popupremind").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败！",
			Data: err,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "删除成功！",
			Data: nil,
		})
	}
}

// Info 详情接口（后端）
func PopupremindControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	popupremind := models.Popupremind{Id: id}
	err := o.Read(&popupremind)

	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: popupremind,
		})
	}
}

// Detail 详情接口（前端）
func PopupremindControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	popupremind := models.Popupremind{Id: id}
	err := o.Read(&popupremind)

	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(200, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: popupremind,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func PopupremindControllerRemind(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
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
			sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM popupremind WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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
		Code:  0,
		Count: len(result),
		Data:  result,
	}
	c.JSON(200, res)
}

// Group 分组统计接口
func PopupremindControllerGroup(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(401, models.ReturnMsg{
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM popupremind " + where + " GROUP BY " + columnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.Query("limit"))
	order := c.Query("order")
	if "desc" == order {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if "asc" == order {
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
func PopupremindControllerValue(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
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
	sql := ""
	where := " WHERE 1 = 1 "

	if "popupremind" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM popupremind " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)

	limit, _ := strconv.Atoi(c.Query("limit"))
	order := c.Query("order")
	if "desc" == order {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if "asc" == order {
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
func PopupremindControllerValueTime(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
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
	sql := ""
	where := " WHERE 1 = 1 "

	if "popupremind" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM popupremind " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM popupremind " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM popupremind " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.Query("limit"))
	order := c.Query("order")
	if "desc" == order {
		sort.Slice(maps, func(i, j int) bool {
			totalI, _ := strconv.Atoi(maps[i]["total"].(string))
			totalJ, _ := strconv.Atoi(maps[j]["total"].(string))
			return totalI > totalJ
		})
	}
	if "asc" == order {
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
func PopupremindControllerValueMul(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
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

	sql := ""
	where := " WHERE 1 = 1 "
	if "popupremind" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM popupremind " + where + " GROUP BY " + xColumnName + " LIMIT 10"
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

// Query 查询单条数据
func PopupremindControllerQuery(c *gin.Context) {
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
	titleColumn := c.Query("title")
	if titleColumn != "" {
		if strings.Contains(titleColumn, "%") {
			cond = cond.And("title__contains", strings.Trim(titleColumn, "%"))
		} else {
			cond = cond.And("title", titleColumn)
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
	briefColumn := c.Query("brief")
	if briefColumn != "" {
		if strings.Contains(briefColumn, "%") {
			cond = cond.And("brief__contains", strings.Trim(briefColumn, "%"))
		} else {
			cond = cond.And("brief", briefColumn)
		}
	}
	contentColumn := c.Query("content")
	if contentColumn != "" {
		if strings.Contains(contentColumn, "%") {
			cond = cond.And("content__contains", strings.Trim(contentColumn, "%"))
		} else {
			cond = cond.And("content", contentColumn)
		}
	}
	remindtimeColumn := c.Query("remindtime")
	if remindtimeColumn != "" {
		if strings.Contains(remindtimeColumn, "%") {
			cond = cond.And("remindtime__contains", strings.Trim(remindtimeColumn, "%"))
		} else {
			cond = cond.And("remindtime", remindtimeColumn)
		}
	}
	var popupremind models.Popupremind
	o.QueryTable("popupremind").SetCond(cond).One(&popupremind)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: popupremind,
	})
}
