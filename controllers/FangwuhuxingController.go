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
	orm.RegisterModel(new(models.Fangwuhuxing))
}

// 分页接口（后端）
func FangwuhuxingPage(c *gin.Context) {
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
	if valiErr != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var fangwuhuxing []models.Fangwuhuxing

	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageErr != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := orm.NewCondition()
	fangwuhuxingColumn := c.Query("fangwuhuxing")
	if fangwuhuxingColumn != "" {
		if strings.Contains(fangwuhuxingColumn, "%") {
			cond = cond.And("fangwuhuxing__contains", strings.Trim(fangwuhuxingColumn, "%"))
		} else {
			cond = cond.And("fangwuhuxing", fangwuhuxingColumn)
		}
	}

	total, _ := o.QueryTable("fangwuhuxing").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sort, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}

	o.QueryTable("fangwuhuxing").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&fangwuhuxing)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      fangwuhuxing,
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

// 分页接口（前端）
func FangwuhuxingLists(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuhuxing []models.Fangwuhuxing
	cond := orm.NewCondition()
	o.QueryTable("fangwuhuxing").SetCond(cond).All(&fangwuhuxing)
	var result []string
	for _, item := range fangwuhuxing {
		result = append(result, item.Fangwuhuxing)
	}
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	}
	c.JSON(200, res)
}

// 分页接口（前端）
func FangwuhuxingList(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuhuxing []models.Fangwuhuxing

	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageErr != nil {
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

	cond := orm.NewCondition()
	fangwuhuxingColumn := c.Query("fangwuhuxing")
	if fangwuhuxingColumn != "" {
		if strings.Contains(fangwuhuxingColumn, "%") {
			cond = cond.And("fangwuhuxing__contains", strings.Trim(fangwuhuxingColumn, "%"))
		} else {
			cond = cond.And("fangwuhuxing", fangwuhuxingColumn)
		}
	}

	total, _ := o.QueryTable("fangwuhuxing").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("fangwuhuxing").SetCond(cond).OrderBy(order).Limit(limit, start).All(&fangwuhuxing)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      fangwuhuxing,
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

// 保存接口（后端）
func FangwuhuxingSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var fangwuhuxing models.Fangwuhuxing
	if err := c.ShouldBindJSON(&fangwuhuxing); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&fangwuhuxing)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(402, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(0, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// 保存接口（前端）
func FangwuhuxingAdd(c *gin.Context) {
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
	if valiErr != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var fangwuhuxing models.Fangwuhuxing
	if err := c.ShouldBindJSON(&fangwuhuxing); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&fangwuhuxing)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(402, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
	} else {
		c.JSON(0, models.ReturnMsg{
			Code: 0,
			Msg:  "添加成功！",
			Data: id,
		})
	}
}

// 更新接口
func FangwuhuxingUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuhuxing models.Fangwuhuxing
	if err := c.ShouldBindJSON(&fangwuhuxing); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	_, err := o.Update(&fangwuhuxing)
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// 删除接口
func FangwuhuxingDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o.QueryTable("fangwuhuxing").Filter("id__in", ids).Delete()
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// 详情接口（后端）
func FangwuhuxingInfo(c *gin.Context) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	fangwuhuxing := models.Fangwuhuxing{Id: id}
	err = o.Read(&fangwuhuxing)

	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(0, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: fangwuhuxing,
		})
	}
}

// 查询单条数据
func FangwuhuxingQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	fangwuhuxingColumn := c.Query("fangwuhuxing")
	if fangwuhuxingColumn != "" {
		if strings.Contains(fangwuhuxingColumn, "%") {
			cond = cond.And("fangwuhuxing__contains", strings.Trim(fangwuhuxingColumn, "%"))
		} else {
			cond = cond.And("fangwuhuxing", fangwuhuxingColumn)
		}
	}

	var fangwuhuxing models.Fangwuhuxing
	o.QueryTable("fangwuhuxing").SetCond(cond).One(&fangwuhuxing)
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangwuhuxing,
	})
}

// 详情接口（前端）
func FangwuhuxingDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	fangwuhuxing := models.Fangwuhuxing{Id: id}
	err = o.Read(&fangwuhuxing)

	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
	} else {
		c.JSON(0, models.ReturnMsg{
			Code: 0,
			Msg:  "请求成功！",
			Data: fangwuhuxing,
		})
	}
}

// 获取需要提醒的记录数接口
func FangwuhuxingRemind(c *gin.Context) {
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
	if valiErr != nil {
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
			sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM fangwuhuxing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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

// 分组统计接口
func FangwuhuxingGroup(c *gin.Context) {
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
	if valiErr != nil {
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM fangwuhuxing " + where + " GROUP BY " + columnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
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
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// 统计指定字段
func FangwuhuxingValue(c *gin.Context) {
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
	if valiErr != nil {
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

	if "fangwuhuxing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangwuhuxing " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
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
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// 按日期统计
func FangwuhuxingValueTime(c *gin.Context) {
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
	if valiErr != nil {
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

	if "fangwuhuxing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuhuxing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuhuxing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuhuxing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
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
	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// 统计指定字段(多)
func FangwuhuxingValueMul(c *gin.Context) {
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
	if valiErr != nil {
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
	if "fangwuhuxing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM fangwuhuxing " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		o.Raw(sql).Values(&maps)
		result = append(result, maps)
	}

	c.JSON(0, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	})
}
