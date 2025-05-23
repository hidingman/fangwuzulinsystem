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
	orm.RegisterModel(new(models.Guke))
}

// 用户注册接口
func GukeRegister(c *gin.Context) {
	o := orm.NewOrm()
	guke := models.Guke{}
	json.NewDecoder(c.Request.Body).Decode(&guke)

	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	_, err := o.Insert(&guke)
	o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()

	var res models.ReturnMsg
	if err != nil {
		res = models.ReturnMsg{
			Code: -1,
			Msg:  "注册失败！",
			Data: nil,
		}
	} else {
		res = models.ReturnMsg{
			Code: 0,
			Msg:  "注册成功！",
			Data: nil,
		}
	}

	c.JSON(200, res)
}

// 用户登录接口
func GukeLogin(c *gin.Context) {
	o := orm.NewOrm()
	guke := models.Guke{}
	password := c.Query("password")

	err := o.QueryTable("guke").Filter("zhanghao", c.Query("username")).Filter("mima", password).One(&guke)
	var res interface{}
	if err != nil {
		res = models.ReturnMsg{
			Code: -1,
			Msg:  "用户名或密码错误！",
			Data: nil,
		}
	} else {
		session := models.Session{
			Id:              guke.Id,
			Username:        guke.Zhanghao,
			Tablename:       "guke",
			Role:            "顾客",
			LoginUserColumn: "zhanghao",
		}
		token, _ := utils.GenerateToken(&session, 0)
		res = models.ReturnToken{
			Code:  0,
			Token: token,
		}
	}

	c.JSON(200, res)
}

// 用户退出接口
func GukeLogout(c *gin.Context) {
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "退出成功！",
		Data: nil,
	})
}

// 获取session的接口
func GukeSession(c *gin.Context) {
	o := orm.NewOrm()
	token := c.GetHeader("Token")
	userinfo, err := utils.ValidateToken(token)
	guke := models.Guke{Id: userinfo.Id}
	o.Read(&guke)
	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: guke,
	})
}

// 忘记密码（找回密码）
func GukeResetPass(c *gin.Context) {
	o := orm.NewOrm()
	username := c.Query("username")
	pwd := "123456"
	o.Raw("UPDATE guke SET mima = ? WHERE zhanghao = ?", pwd, username).Exec()
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "密码已重置为：123456",
		Data: nil,
	})
}

// 分页接口（后端）
func GukePage(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var guke []models.Guke

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
	zhanghaoColumn := c.Query("zhanghao")
	if zhanghaoColumn != "" {
		if strings.Contains(zhanghaoColumn, "%") {
			cond = cond.And("zhanghao__contains", strings.Trim(zhanghaoColumn, "%"))
		} else {
			cond = cond.And("zhanghao", zhanghaoColumn)
		}
	}
	// 省略其他字段条件，格式类似 zhanghaoColumn
	mimaColumn := c.Query("mima")
	if mimaColumn != "" {
		if strings.Contains(mimaColumn, "%") {
			cond = cond.And("mima__contains", strings.Trim(mimaColumn, "%"))
		} else {
			cond = cond.And("mima", mimaColumn)
		}
	}

	total, _ := o.QueryTable("guke").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sort, ",")

	for index, value := range sortlist {
		if orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	o.QueryTable("guke").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&guke)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      guke,
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
func GukeLists(c *gin.Context) {
	o := orm.NewOrm()
	var guke []models.Guke
	cond := orm.NewCondition()
	o.QueryTable("guke").SetCond(cond).All(&guke)
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: guke,
	}
	c.JSON(200, res)
}

// 分页接口（前端）
func GukeList(c *gin.Context) {
	o := orm.NewOrm()
	var guke []models.Guke

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
	// 省略字段条件构建，同 Page 方法
	zhanghaoColumn := c.Query("zhanghao")
	if zhanghaoColumn != "" {
		if strings.Contains(zhanghaoColumn, "%") {
			cond = cond.And("zhanghao__contains", strings.Trim(zhanghaoColumn, "%"))
		} else {
			cond = cond.And("zhanghao", zhanghaoColumn)
		}
	}

	total, _ := o.QueryTable("guke").SetCond(cond).Count()
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	o.QueryTable("guke").SetCond(cond).OrderBy(order).Limit(limit, start).All(&guke)
	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      guke,
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
func GukeSave(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	guke := models.Guke{}
	err = json.NewDecoder(c.Request.Body).Decode(&guke)
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&guke)
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

// 保存接口（前端）
func GukeAdd(c *gin.Context) {
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
	guke := models.Guke{}
	json.NewDecoder(c.Request.Body).Decode(&guke)
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&guke)
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

// 更新接口
func GukeUpdate(c *gin.Context) {
	o := orm.NewOrm()
	guke := models.Guke{}
	if err := json.NewDecoder(c.Request.Body).Decode(&guke); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求数据解析失败",
			Data: err,
		})
		return
	}
	_, err := o.Update(&guke)
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

// 删除接口
func GukeDelete(c *gin.Context) {
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
	_, err := o.QueryTable("guke").Filter("id__in", ids).Delete()
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

// 详情接口（后端）
func GukeInfo(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "ID 格式错误",
			Data: nil,
		})
		return
	}
	guke := models.Guke{Id: id}
	err = o.Read(&guke)

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
			Data: guke,
		})
	}
}

// 查询单条数据
func GukeQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	zhanghaoColumn := c.Query("zhanghao")
	if zhanghaoColumn != "" {
		if strings.Contains(zhanghaoColumn, "%") {
			cond = cond.And("zhanghao__contains", strings.Trim(zhanghaoColumn, "%"))
		} else {
			cond = cond.And("zhanghao", zhanghaoColumn)
		}
	}
	// 省略其他字段条件，格式类似 zhanghaoColumn
	mimaColumn := c.Query("mima")
	if mimaColumn != "" {
		if strings.Contains(mimaColumn, "%") {
			cond = cond.And("mima__contains", strings.Trim(mimaColumn, "%"))
		} else {
			cond = cond.And("mima", mimaColumn)
		}
	}

	var guke models.Guke
	o.QueryTable("guke").SetCond(cond).One(&guke)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: guke,
	})
}

// 详情接口（前端）
func GukeDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "ID 格式错误",
			Data: nil,
		})
		return
	}
	guke := models.Guke{Id: id}
	err = o.Read(&guke)

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
			Data: guke,
		})
	}
}

// 获取需要提醒的记录数接口
func GukeRemind(c *gin.Context) {
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
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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

// 分组统计接口
func GukeGroup(c *gin.Context) {
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM guke " + where + " GROUP BY " + columnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.Query("limit"))
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

// 统计指定字段
func GukeValue(c *gin.Context) {
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

	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM guke " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	o.Raw(sql).Values(&maps)

	limit, _ := strconv.Atoi(c.Query("limit"))
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

// 按日期统计
func GukeValueTime(c *gin.Context) {
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

	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	limit, _ := strconv.Atoi(c.Query("limit"))
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

// 统计指定字段(多)
func GukeValueMul(c *gin.Context) {
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
	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM guke " + where + " GROUP BY " + xColumnName + " LIMIT 10"
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

// 总数接口
func GukeCount(c *gin.Context) {
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

	cond := orm.NewCondition()
	zhanghaoColumn := c.Query("zhanghao")
	if zhanghaoColumn != "" {
		if strings.Contains(zhanghaoColumn, "%") {
			cond = cond.And("zhanghao__contains", strings.Trim(zhanghaoColumn, "%"))
		} else {
			cond = cond.And("zhanghao", zhanghaoColumn)
		}
	}
	xingmingColumn := c.Query("xingming")
	if xingmingColumn != "" {
		if strings.Contains(xingmingColumn, "%") {
			cond = cond.And("xingming__contains", strings.Trim(xingmingColumn, "%"))
		} else {
			cond = cond.And("xingming", xingmingColumn)
		}
	}

	o := orm.NewOrm()
	cnt, _ := o.QueryTable("guke").SetCond(cond).Count()

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: cnt,
	})
}
