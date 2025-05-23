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
	orm.RegisterModel(new(models.Fangwuzulin))
}

// Page 分页接口（后端）
func FangwuzulinControllerPage(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "Token 为空",
		})
		return
	}

	// 验证 Token
	userInfo, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "Token 验证失败",
		})
		return
	}
	o := orm.NewOrm()
	var fangwuzulin []models.Fangwuzulin

	// 获取分页参数
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

	// 定义查询字段映射
	fields := map[string]string{
		"fangwumingcheng":  "fangwumingcheng",
		"fangwutupian":     "fangwutupian",
		"fangwuhuxing":     "fangwuhuxing",
		"qishishijian":     "qishishijian",
		"jieshushijian":    "jieshushijian",
		"zulinshijian":     "zulinshijian",
		"zujinjiage":       "zujinjiage",
		"zongji":           "zongji",
		"yajin":            "yajin",
		"jifen":            "jifen",
		"zulinxiangqing":   "zulinxiangqing",
		"zulinbeizhu":      "zulinbeizhu",
		"gonghao":          "gonghao",
		"yuangongxingming": "yuangongxingming",
		"zhanghao":         "zhanghao",
		"xingming":         "xingming",
		"ispay":            "ispay",
	}

	// 处理查询条件
	for param, field := range fields {
		value := c.Query(param)
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(field+"__contains", strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)

	}
	if tableName == "guke" {
		cond = cond.And("zhanghao", userInfo.Username)

	}

	// 获取总记录数
	total, err := o.QueryTable("fangwuzulin").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总记录数失败",
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

	// 查询数据
	_, err = o.QueryTable("fangwuzulin").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      fangwuzulin,
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
func FangwuzulinControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuzulin []models.Fangwuzulin
	cond := orm.NewCondition()

	// 执行查询并处理错误
	_, err := o.QueryTable("fangwuzulin").SetCond(cond).All(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: nil,
		})
		return
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangwuzulin,
	}

	c.JSON(http.StatusOK, res)
}

// 分页接口（前端）
func FangwuzulinControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuzulin []models.Fangwuzulin

	// 获取分页参数
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

	// 定义查询字段映射
	fields := map[string]string{
		"fangwumingcheng":  "fangwumingcheng",
		"fangwutupian":     "fangwutupian",
		"fangwuhuxing":     "fangwuhuxing",
		"qishishijian":     "qishishijian",
		"jieshushijian":    "jieshushijian",
		"zulinshijian":     "zulinshijian",
		"zujinjiage":       "zujinjiage",
		"zongji":           "zongji",
		"yajin":            "yajin",
		"jifen":            "jifen",
		"zulinxiangqing":   "zulinxiangqing",
		"zulinbeizhu":      "zulinbeizhu",
		"gonghao":          "gonghao",
		"yuangongxingming": "yuangongxingming",
		"zhanghao":         "zhanghao",
		"xingming":         "xingming",
		"ispay":            "ispay",
	}

	// 处理查询条件
	for param, field := range fields {
		value := c.Query(param)
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(field+"__contains", strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	// 获取总记录数
	total, err := o.QueryTable("fangwuzulin").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总记录数失败",
			Data: nil,
		})
		return
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit

	// 查询数据
	_, err = o.QueryTable("fangwuzulin").SetCond(cond).OrderBy(order).Limit(limit, start).All(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      fangwuzulin,
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

func FangwuzulinControllerSave(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "Token 为空",
			Data: nil,
		})
		return
	}

	// 验证 Token
	_, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "Token 验证失败",
			Data: nil,
		})
		return
	}
	o := orm.NewOrm()
	var fangwuzulin models.Fangwuzulin

	// 解析请求体
	if err := c.ShouldBindJSON(&fangwuzulin); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 关闭外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "关闭外键检查失败",
			Data: err,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&fangwuzulin)

	// 开启外键检查
	if _, errExec := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); errExec != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "开启外键检查失败",
			Data: errExec,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "添加成功！",
		Data: id,
	})
}

// Add 保存接口（前端）
func FangwuzulinControllerAdd(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	var fangwuzulin models.Fangwuzulin

	// 解析请求体
	if err := c.ShouldBindJSON(&fangwuzulin); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 关闭外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "关闭外键检查失败",
			Data: err,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&fangwuzulin)

	// 开启外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "开启外键检查失败",
			Data: err,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！",
			Data: err,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "添加成功！",
		Data: id,
	})
}

// Update 更新接口
func FangwuzulinControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuzulin models.Fangwuzulin

	// 解析请求体
	if err := c.ShouldBindJSON(&fangwuzulin); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 执行更新操作
	num, err := o.Update(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "更新失败",
			Data: err,
		})
		return
	}

	if num == 0 {
		c.JSON(http.StatusNotFound, models.ReturnMsg{
			Code: 404,
			Msg:  "未找到要更新的记录",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: num,
	})
}

// Delete 删除接口
func FangwuzulinControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int

	// 解析请求体
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 执行删除操作
	num, err := o.QueryTable("fangwuzulin").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败",
			Data: err,
		})
		return
	}

	if num == 0 {
		c.JSON(http.StatusNotFound, models.ReturnMsg{
			Code: 404,
			Msg:  "未找到要删除的记录",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: num,
	})
}

// 详情接口（后端）
func FangwuzulinControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	// 从 URL 参数获取 id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "无效的 ID 参数",
			Data: nil,
		})
		return
	}

	fangwuzulin := models.Fangwuzulin{Id: id}
	// 读取数据
	err = o.Read(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangwuzulin,
	})
}

// Query 查询单条数据
func FangwuzulinControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	// 定义查询字段映射
	fields := map[string]string{
		"fangwumingcheng":  "fangwumingcheng",
		"fangwutupian":     "fangwutupian",
		"fangwuhuxing":     "fangwuhuxing",
		"qishishijian":     "qishishijian",
		"jieshushijian":    "jieshushijian",
		"zulinshijian":     "zulinshijian",
		"zujinjiage":       "zujinjiage",
		"zongji":           "zongji",
		"yajin":            "yajin",
		"jifen":            "jifen",
		"zulinxiangqing":   "zulinxiangqing",
		"zulinbeizhu":      "zulinbeizhu",
		"gonghao":          "gonghao",
		"yuangongxingming": "yuangongxingming",
		"zhanghao":         "zhanghao",
		"xingming":         "xingming",
		"ispay":            "ispay",
	}

	// 处理查询条件
	for param, field := range fields {
		value := c.Query(param)
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(field+"__contains", strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	var fangwuzulin models.Fangwuzulin
	// 执行查询
	err := o.QueryTable("fangwuzulin").SetCond(cond).One(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败！",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangwuzulin,
	})
}

func FangwuzulinControllerDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	// 从路径参数获取 id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "无效的 ID 参数",
			Data: nil,
		})
		return
	}

	fangwuzulin := models.Fangwuzulin{Id: id}
	// 读取数据
	err = o.Read(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangwuzulin,
	})
}

// Remind 获取需要提醒的记录数接口
func FangwuzulinControllerRemind(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	// 从路径参数获取 columnName 和 type
	columnName := c.Param("columnName")
	typeP := c.Param("type")
	remindstart := c.Query("remindstart")
	remindend := c.Query("remindend")

	if typeP == "1" {
		if remindstart != "" {
			sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, errStart := strconv.Atoi(remindstart)
			end, errEnd := strconv.Atoi(remindend)
			if errStart != nil || errEnd != nil {
				c.JSON(http.StatusBadRequest, models.ReturnMsg{
					Code: 400,
					Msg:  "日期参数转换失败",
					Data: nil,
				})
				return
			}
			if start > end {
				sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, err := strconv.Atoi(remindstart)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ReturnMsg{
					Code: 400,
					Msg:  "开始日期参数转换失败",
					Data: nil,
				})
				return
			}
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, err := strconv.Atoi(remindend)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ReturnMsg{
					Code: 400,
					Msg:  "结束日期参数转换失败",
					Data: nil,
				})
				return
			}
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, errStart := strconv.Atoi(remindstart)
			end, errEnd := strconv.Atoi(remindend)
			if errStart != nil || errEnd != nil {
				c.JSON(http.StatusBadRequest, models.ReturnMsg{
					Code: 400,
					Msg:  "日期参数转换失败",
					Data: nil,
				})
				return
			}
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT fangwumingcheng FROM fangwuzulin WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
		}
	}

	o := orm.NewOrm()
	var result []string
	// 执行查询并处理错误
	_, err = o.Raw(sql).QueryRows(&result)
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
func FangwuzulinControllerGroup(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	// 从路径参数获取 columnName
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM fangwuzulin " + where + " GROUP BY " + columnName
	var maps []orm.Params
	// 执行查询并处理错误
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
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

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// Value 统计指定字段
func FangwuzulinControllerValue(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	// 从路径参数获取 xColumnName 和 yColumnName
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

	// 此条件恒为 false，可移除
	// if "fangwuzulin" == "orders" {
	// 	where += " AND status IN ('已支付', '已发货', '已完成') "
	// }

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangwuzulin " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	// 执行查询并处理错误
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
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

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueTime 按日期统计
func FangwuzulinControllerValueTime(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	// 从路径参数获取 xColumnName、yColumnName 和 timeStatType
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

	// 此条件恒为 false，可移除
	// if "fangwuzulin" == "orders" {
	// 	where += " AND status IN ('已支付', '已发货', '已完成') "
	// }

	var errUnescape error
	timeStatType, errUnescape = url.QueryUnescape(timeStatType)
	if errUnescape != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "时间统计类型参数解码失败",
			Data: errUnescape,
		})
		return
	}

	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuzulin " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuzulin " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM fangwuzulin " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	// 执行查询并处理错误
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败",
			Data: err,
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
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

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	})
}

// ValueMul 统计指定字段(多)
func FangwuzulinControllerValueMul(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
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
	// 从路径参数获取 xColumnName
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

	// 此条件恒为 false，可移除
	// if "fangwuzulin" == "orders" {
	// 	where += " AND status IN ('已支付', '已发货', '已完成') "
	// }

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM fangwuzulin " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		// 执行查询并处理错误
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

// Count 总数接口
func FangwuzulinControllerCount(c *gin.Context) {
	// 从请求头获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	// 验证 Token
	userInfo, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	cond := orm.NewCondition()
	// 从查询参数获取 fangwumingcheng
	fangwumingchengColumn := c.Query("fangwumingcheng")
	if fangwumingchengColumn != "" {
		if strings.Contains(fangwumingchengColumn, "%") {
			cond = cond.And("fangwumingcheng__contains", strings.Trim(fangwumingchengColumn, "%"))
		} else {
			cond = cond.And("fangwumingcheng", fangwumingchengColumn)
		}
	}

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)
	}
	if tableName == "guke" {
		cond = cond.And("zhanghao", userInfo.Username)
	}

	o := orm.NewOrm()
	// 执行计数操作并处理错误
	cnt, err := o.QueryTable("fangwuzulin").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "统计总数失败",
			Data: err,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: cnt,
	})
}
