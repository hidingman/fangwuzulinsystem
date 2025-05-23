package controllers

import (
	"log"
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
	orm.RegisterModel(new(models.Kaoqindaka))
}

// Page 分页接口（后端）
func KaoqindakaControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)
	o := orm.NewOrm()
	var kaoqindaka []models.Kaoqindaka

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")
	if sort == "" {
		sort = "id"
	}
	order := c.Query("order")

	cond := orm.NewCondition()
	dakashijianstart := c.Query("dakashijianstart")
	dakashijianend := c.Query("dakashijianend")
	if dakashijianstart != "" {
		cond = cond.And("dakashijian__gte", dakashijianstart)
	}
	if dakashijianend != "" {
		cond = cond.And("dakashijian__lte", dakashijianend)
	}
	dakaleixingColumn := c.Query("dakaleixing")
	if dakaleixingColumn != "" {
		if strings.Contains(dakaleixingColumn, "%") {
			cond = cond.And("dakaleixing__contains", strings.Trim(dakaleixingColumn, "%"))
		} else {
			cond = cond.And("dakaleixing", dakaleixingColumn)
		}
	}
	dakashijianColumn := c.Query("dakashijian")
	if dakashijianColumn != "" {
		if strings.Contains(dakashijianColumn, "%") {
			cond = cond.And("dakashijian__contains", strings.Trim(dakashijianColumn, "%"))
		} else {
			cond = cond.And("dakashijian", dakashijianColumn)
		}
	}
	dakaxiangqingColumn := c.Query("dakaxiangqing")
	if dakaxiangqingColumn != "" {
		if strings.Contains(dakaxiangqingColumn, "%") {
			cond = cond.And("dakaxiangqing__contains", strings.Trim(dakaxiangqingColumn, "%"))
		} else {
			cond = cond.And("dakaxiangqing", dakaxiangqingColumn)
		}
	}
	gonghaoColumn := c.Query("gonghao")
	if gonghaoColumn != "" {
		if strings.Contains(gonghaoColumn, "%") {
			cond = cond.And("gonghao__contains", strings.Trim(gonghaoColumn, "%"))
		} else {
			cond = cond.And("gonghao", gonghaoColumn)
		}
	}
	yuangongxingmingColumn := c.Query("yuangongxingming")
	if yuangongxingmingColumn != "" {
		if strings.Contains(yuangongxingmingColumn, "%") {
			cond = cond.And("yuangongxingming__contains", strings.Trim(yuangongxingmingColumn, "%"))
		} else {
			cond = cond.And("yuangongxingming", yuangongxingmingColumn)
		}
	}
	touxiangColumn := c.Query("touxiang")
	if touxiangColumn != "" {
		if strings.Contains(touxiangColumn, "%") {
			cond = cond.And("touxiang__contains", strings.Trim(touxiangColumn, "%"))
		} else {
			cond = cond.And("touxiang", touxiangColumn)
		}
	}

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)

	}

	total, _ := o.QueryTable("kaoqindaka").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sort, ",")

	for index, value := range sortlist {
		if orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("kaoqindaka").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&kaoqindaka)
	returnPage := models.ReturnPage{
		CurrPage:  int(page),
		List:      kaoqindaka,
		PageSize:  int(limit),
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		0, "请求成功！", returnPage,
	}
	c.JSON(http.StatusOK, res)
}

// Lists 分页接口（前端）
func KaoqindakaControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var kaoqindaka []models.Kaoqindaka
	cond := orm.NewCondition()
	o.QueryTable("kaoqindaka").SetCond(cond).All(&kaoqindaka)
	res := models.ReturnMsg{
		0, "请求成功！", kaoqindaka,
	}
	c.JSON(http.StatusOK, res)
}

// List 分页接口（前端）
func KaoqindakaControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var kaoqindaka []models.Kaoqindaka

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
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

	cond := orm.NewCondition()
	dakashijianstart := c.Query("dakashijianstart")
	dakashijianend := c.Query("dakashijianend")
	if dakashijianstart != "" {
		cond = cond.And("dakashijian__gte", dakashijianstart)
	}
	if dakashijianend != "" {
		cond = cond.And("dakashijian__lte", dakashijianend)
	}
	dakaleixingColumn := c.Query("dakaleixing")
	if dakaleixingColumn != "" {
		if strings.Contains(dakaleixingColumn, "%") {
			cond = cond.And("dakaleixing__contains", strings.Trim(dakaleixingColumn, "%"))
		} else {
			cond = cond.And("dakaleixing", dakaleixingColumn)
		}
	}
	dakashijianColumn := c.Query("dakashijian")
	if dakashijianColumn != "" {
		if strings.Contains(dakashijianColumn, "%") {
			cond = cond.And("dakashijian__contains", strings.Trim(dakashijianColumn, "%"))
		} else {
			cond = cond.And("dakashijian", dakashijianColumn)
		}
	}
	dakaxiangqingColumn := c.Query("dakaxiangqing")
	if dakaxiangqingColumn != "" {
		if strings.Contains(dakaxiangqingColumn, "%") {
			cond = cond.And("dakaxiangqing__contains", strings.Trim(dakaxiangqingColumn, "%"))
		} else {
			cond = cond.And("dakaxiangqing", dakaxiangqingColumn)
		}
	}
	gonghaoColumn := c.Query("gonghao")
	if gonghaoColumn != "" {
		if strings.Contains(gonghaoColumn, "%") {
			cond = cond.And("gonghao__contains", strings.Trim(gonghaoColumn, "%"))
		} else {
			cond = cond.And("gonghao", gonghaoColumn)
		}
	}
	yuangongxingmingColumn := c.Query("yuangongxingming")
	if yuangongxingmingColumn != "" {
		if strings.Contains(yuangongxingmingColumn, "%") {
			cond = cond.And("yuangongxingming__contains", strings.Trim(yuangongxingmingColumn, "%"))
		} else {
			cond = cond.And("yuangongxingming", yuangongxingmingColumn)
		}
	}
	touxiangColumn := c.Query("touxiang")
	if touxiangColumn != "" {
		if strings.Contains(touxiangColumn, "%") {
			cond = cond.And("touxiang__contains", strings.Trim(touxiangColumn, "%"))
		} else {
			cond = cond.And("touxiang", touxiangColumn)
		}
	}

	total, _ := o.QueryTable("kaoqindaka").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("kaoqindaka").SetCond(cond).OrderBy(order).Limit(limit, start).All(&kaoqindaka)
	returnPage := models.ReturnPage{
		CurrPage:  int(page),
		List:      kaoqindaka,
		PageSize:  int(limit),
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		0, "请求成功！", returnPage,
	}

	c.JSON(http.StatusOK, res)
}

// Save 保存接口（后端）
func KaoqindakaControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var kaoqindaka models.Kaoqindaka
	c.ShouldBindJSON(&kaoqindaka)

	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&kaoqindaka)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			402, "添加失败！", err,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			0, "添加成功！", id,
		})
	}
}

// Add 保存接口（前端）
func KaoqindakaControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	_, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}

	o := orm.NewOrm()
	var kaoqindaka models.Kaoqindaka
	if err := c.ShouldBindJSON(&kaoqindaka); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "请求体解析失败", err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	log.Println(kaoqindaka)
	id, err := o.Insert(&kaoqindaka)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			402, "添加失败！", err,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			0, "添加成功！", id,
		})
	}
}

// Update 更新接口
func KaoqindakaControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var kaoqindaka models.Kaoqindaka
	if err := c.ShouldBindJSON(&kaoqindaka); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "请求体解析失败", err,
		})
		return
	}
	_, err := o.Update(&kaoqindaka)

	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "更新成功！", err,
	})
}

// Delete 删除接口
func KaoqindakaControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "请求体解析失败", err,
		})
		return
	}
	o.QueryTable("kaoqindaka").Filter("id__in", ids).Delete()
	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "删除成功！", nil,
	})
}

// Info 详情接口（后端）
func KaoqindakaControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "参数解析失败", nil,
		})
		return
	}
	kaoqindaka := models.Kaoqindaka{Id: id}
	err = o.Read(&kaoqindaka)

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			500, "服务器错误！", nil,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			0, "请求成功！", kaoqindaka,
		})
	}
}

// Query 查询单条数据
func KaoqindakaControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	dakaleixingColumn := c.Query("dakaleixing")
	if dakaleixingColumn != "" {
		if strings.Contains(dakaleixingColumn, "%") {
			cond = cond.And("dakaleixing__contains", strings.Trim(dakaleixingColumn, "%"))
		} else {
			cond = cond.And("dakaleixing", dakaleixingColumn)
		}
	}
	dakashijianColumn := c.Query("dakashijian")
	if dakashijianColumn != "" {
		if strings.Contains(dakashijianColumn, "%") {
			cond = cond.And("dakashijian__contains", strings.Trim(dakashijianColumn, "%"))
		} else {
			cond = cond.And("dakashijian", dakashijianColumn)
		}
	}
	dakaxiangqingColumn := c.Query("dakaxiangqing")
	if dakaxiangqingColumn != "" {
		if strings.Contains(dakaxiangqingColumn, "%") {
			cond = cond.And("dakaxiangqing__contains", strings.Trim(dakaxiangqingColumn, "%"))
		} else {
			cond = cond.And("dakaxiangqing", dakaxiangqingColumn)
		}
	}
	gonghaoColumn := c.Query("gonghao")
	if gonghaoColumn != "" {
		if strings.Contains(gonghaoColumn, "%") {
			cond = cond.And("gonghao__contains", strings.Trim(gonghaoColumn, "%"))
		} else {
			cond = cond.And("gonghao", gonghaoColumn)
		}
	}
	yuangongxingmingColumn := c.Query("yuangongxingming")
	if yuangongxingmingColumn != "" {
		if strings.Contains(yuangongxingmingColumn, "%") {
			cond = cond.And("yuangongxingming__contains", strings.Trim(yuangongxingmingColumn, "%"))
		} else {
			cond = cond.And("yuangongxingming", yuangongxingmingColumn)
		}
	}
	touxiangColumn := c.Query("touxiang")
	if touxiangColumn != "" {
		if strings.Contains(touxiangColumn, "%") {
			cond = cond.And("touxiang__contains", strings.Trim(touxiangColumn, "%"))
		} else {
			cond = cond.And("touxiang", touxiangColumn)
		}
	}
	var kaoqindaka models.Kaoqindaka
	o.QueryTable("kaoqindaka").SetCond(cond).One(&kaoqindaka)
	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "请求成功！", kaoqindaka,
	})
}

// Detail 详情接口（前端）
func KaoqindakaControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "参数解析失败", nil,
		})
		return
	}
	kaoqindaka := models.Kaoqindaka{Id: id}
	err = o.Read(&kaoqindaka)

	if err != nil {
		c.JSON(http.StatusOK, models.ReturnMsg{
			500, "服务器错误！", nil,
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			0, "请求成功！", kaoqindaka,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func KaoqindakaControllerRemind(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}

	where := " 1=1 "
	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		where += " AND gonghao = '" + userInfo.Username + "' "
	}

	sql := "SELECT 0 AS count"
	columnName := c.Param("columnName")
	typeP := c.Param("type")
	remindstart := c.Query("remindstart")
	remindend := c.Query("remindend")
	if typeP == "1" {
		if remindstart != "" {
			sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM kaoqindaka WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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
func KaoqindakaControllerGroup(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM kaoqindaka " + where + " GROUP BY " + columnName
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "请求成功！", maps,
	})
}

// Value 统计指定字段
func KaoqindakaControllerValue(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
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

	if "kaoqindaka" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM kaoqindaka " + where + " GROUP BY " + xColumnName
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "请求成功！", maps,
	})
}

// ValueTime 按日期统计
func KaoqindakaControllerValueTime(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
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

	if "kaoqindaka" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM kaoqindaka " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM kaoqindaka " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM kaoqindaka " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
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
	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "请求成功！", maps,
	})
}

// ValueMul 统计指定字段(多)
func KaoqindakaControllerValueMul(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
		})
		return
	}
	userInfo, valiErr := utils.ValidateToken(token)

	if token == "" || valiErr != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			401, "您的权限不够！", nil,
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
	if "kaoqindaka" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM kaoqindaka " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		o.Raw(sql).Values(&maps)
		result = append(result, maps)
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		0, "请求成功！", result,
	})
}
