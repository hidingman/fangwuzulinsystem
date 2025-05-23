package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func init() {
	orm.RegisterModel(new(models.Tuizushenqing))
}

// 构建查询条件
func buildCondition(c *gin.Context) *orm.Condition {
	cond := orm.NewCondition()
	fields := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "yajin", "tuizushijian",
		"tuizuyuanyin", "tuizuxiangqing", "gonghao", "yuangongxingming", "zhanghao",
		"xingming", "sfsh", "shhf", "ispay",
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

// 分页接口（后端）
func TuizushenqingControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var tuizushenqing []models.Tuizushenqing

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := buildCondition(c)
	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)
	}
	if tableName == "guke" {
		cond = cond.And("zhanghao", userInfo.Username)
	}

	total, _ := o.QueryTable("tuizushenqing").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	sortlist := strings.Split(sort, ",")
	orderlist := strings.Split(order, ",")

	for index, value := range sortlist {
		if len(orderlist) > index && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("tuizushenqing").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&tuizushenqing)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      tuizushenqing,
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

// 分页接口（前端）
func TuizushenqingControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var tuizushenqing []models.Tuizushenqing
	cond := orm.NewCondition()
	o.QueryTable("tuizushenqing").SetCond(cond).All(&tuizushenqing)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: tuizushenqing,
	}
	c.JSON(http.StatusOK, res)
}

// 分页接口（前端）
func TuizushenqingControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var tuizushenqing []models.Tuizushenqing

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

	cond := buildCondition(c)

	total, _ := o.QueryTable("tuizushenqing").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("tuizushenqing").SetCond(cond).OrderBy(order).Limit(limit, start).All(&tuizushenqing)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      tuizushenqing,
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

// 保存接口（后端）
func TuizushenqingControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var tuizushenqing models.Tuizushenqing
	if err := c.ShouldBindJSON(&tuizushenqing); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&tuizushenqing)
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

// 保存接口（前端）
func TuizushenqingControllerAdd(c *gin.Context) {
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
	var tuizushenqing models.Tuizushenqing
	if err := c.ShouldBindJSON(&tuizushenqing); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&tuizushenqing)
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

// 更新接口
func TuizushenqingControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var tuizushenqing models.Tuizushenqing
	if err := c.ShouldBindJSON(&tuizushenqing); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数错误！",
			Data: err,
		})
		return
	}
	_, err := o.Update(&tuizushenqing)
	if tuizushenqing.Sfsh == "是" {
		fangyuanxinxi := models.Fangyuanxinxi{}
		err := o.QueryTable(new(models.Fangyuanxinxi)).Filter("fangwumingcheng", tuizushenqing.Fangwumingcheng).One(&fangyuanxinxi)
		if err != nil {
			if err == orm.ErrNoRows {
				c.JSON(404, models.ReturnMsg{
					Code: 404,
					Msg:  "未找到对应的房源信息",
					Data: nil,
				})
			} else {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "查询房源信息失败",
					Data: err,
				})
			}
			return
		}

		if fangyuanxinxi.Fangwuzhuangtai == "已租" {
			// 更新租赁状态为已租
			fangyuanxinxi.Fangwuzhuangtai = "未租"
			o.Update(&fangyuanxinxi, "fangwuzhuangtai")

		}
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// 删除接口
func TuizushenqingControllerDelete(c *gin.Context) {
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
	o.QueryTable("tuizushenqing").Filter("id__in", ids).Delete()
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// 详情接口（后端）
func TuizushenqingControllerInfo(c *gin.Context) {
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
	tuizushenqing := models.Tuizushenqing{Id: id}
	err = o.Read(&tuizushenqing)

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
			Data: tuizushenqing,
		})
	}
}

// 查询单条数据
func TuizushenqingControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := buildCondition(c)
	var tuizushenqing models.Tuizushenqing
	o.QueryTable("tuizushenqing").SetCond(cond).One(&tuizushenqing)
	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: tuizushenqing,
	})
}

// 详情接口（前端）
func TuizushenqingControllerDetail(c *gin.Context) {
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
	tuizushenqing := models.Tuizushenqing{Id: id}
	err = o.Read(&tuizushenqing)

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
			Data: tuizushenqing,
		})
	}
}

// 获取需要提醒的记录数接口
func TuizushenqingControllerRemind(c *gin.Context) {
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
			sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM tuizushenqing WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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

// 分组统计接口
func TuizushenqingControllerGroup(c *gin.Context) {
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM tuizushenqing " + where + " GROUP BY " + columnName
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

// 统计指定字段
func TuizushenqingControllerValue(c *gin.Context) {
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

	if "tuizushenqing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM tuizushenqing " + where + " GROUP BY " + xColumnName
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

// 按日期统计
func TuizushenqingControllerValueTime(c *gin.Context) {
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

	if "tuizushenqing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM tuizushenqing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM tuizushenqing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM tuizushenqing " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
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

// 统计指定字段(多)
func TuizushenqingControllerValueMul(c *gin.Context) {
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
	if "tuizushenqing" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM tuizushenqing " + where + " GROUP BY " + xColumnName + " LIMIT 10"
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

// 批量审核接口
func TuizushenqingControllerShBatch(c *gin.Context) {
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
	o.QueryTable("tuizushenqing").Filter("id__in", ids).Update(orm.Params{
		"sfsh": c.Query("sfsh"),
		"shhf": c.Query("shhf"),
	})

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "审核成功！",
		Data: nil,
	})
}
