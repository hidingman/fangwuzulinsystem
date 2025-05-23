package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
	"log"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func init() {
	orm.RegisterModel(new(models.Zulinhetong))
}

// Page 分页接口（后端）
func ZulinhetongControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ReturnMsg{
			Code: 401,
			Msg:  "Token 验证失败",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var zulinhetong []models.Zulinhetong

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	cond := orm.NewCondition()
	// 处理查询条件
	fields := map[string]string{
		"fangwumingcheng":  "fangwumingcheng",
		"fangwutupian":     "fangwutupian",
		"fangwuhuxing":     "fangwuhuxing",
		"yajin":            "yajin",
		"jifen":            "jifen",
		"hetongbianhao":    "hetongbianhao",
		"hetongfujian":     "hetongfujian",
		"qiandingshijian":  "qiandingshijian",
		"hetongxiangqing":  "hetongxiangqing",
		"hetongshixiang":   "hetongshixiang",
		"gonghao":          "gonghao",
		"yuangongxingming": "yuangongxingming",
		"zhanghao":         "zhanghao",
		"xingming":         "xingming",
		"sfsh":             "sfsh",
		"shhf":             "shhf",
	}

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

	total, err := o.QueryTable("zulinhetong").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
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
		if len(orderlist) > index && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}

	_, err = o.QueryTable("zulinhetong").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&zulinhetong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: err,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      zulinhetong,
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
func ZulinhetongControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuzulin []models.Fangwuzulin
	cond := orm.NewCondition()

	_, err := o.QueryTable("fangwuzulin").SetCond(cond).All(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: err,
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

// List 分页接口（前端）
func ZulinhetongControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var fangwuzulin []models.Fangwuzulin

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	if order == "desc" {
		order = "-" + sort
	} else {
		order = sort
	}

	cond := orm.NewCondition()
	// 处理查询条件
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

	total, err := o.QueryTable("fangwuzulin").SetCond(cond).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "获取总数失败",
			Data: err,
		})
		return
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit

	_, err = o.QueryTable("fangwuzulin").SetCond(cond).OrderBy(order).Limit(limit, start).All(&fangwuzulin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败",
			Data: err,
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

// Save 保存接口（后端）
func ZulinhetongControllerSave(c *gin.Context) {
	// 检查 Token 头是否存在
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			401, "缺少 Token 信息！", nil,
		})
		return
	}

	// 验证 Token
	_, valiErr := utils.ValidateToken(token)
	if valiErr != nil {
		c.JSON(401, models.ReturnMsg{
			401, "Token 验证失败！", valiErr,
		})
		return
	}

	o := orm.NewOrm()
	zulinhetong := models.Zulinhetong{}

	// 反序列化请求体
	if err := json.NewDecoder(c.Request.Body).Decode(&zulinhetong); err != nil {
		c.JSON(400, models.ReturnMsg{
			400, "请求体解析失败！", err,
		})
		return
	}
	log.Println(zulinhetong)
	// 关闭外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(500, models.ReturnMsg{
			500, "关闭外键检查失败！", err,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&zulinhetong)

	// 恢复外键检查
	if _, err2 := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); err2 != nil {
		c.JSON(500, models.ReturnMsg{
			500, "恢复外键检查失败！", err2,
		})
		return
	}

	// 处理插入错误
	if err != nil {
		c.JSON(402, models.ReturnMsg{
			402, "添加失败！", err,
		})
	} else {
		c.JSON(0, models.ReturnMsg{
			0, "添加成功！", id,
		})
	}
}

// Add 保存接口（前端）
func ZulinhetongControllerAdd(c *gin.Context) {
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
	var zulinhetong models.Zulinhetong

	// 解析请求体
	if err := c.ShouldBindJSON(&zulinhetong); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 关闭外键检查
	_, err = o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "关闭外键检查失败",
			Data: err,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&zulinhetong)

	// 开启外键检查
	_, err = o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "开启外键检查失败",
			Data: err,
		})
		return
	}

	// 根据插入结果返回响应
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
func ZulinhetongControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	zulinhetong := models.Zulinhetong{}

	// 解析请求体
	if err := json.NewDecoder(c.Request.Body).Decode(&zulinhetong); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}

	// 执行更新操作
	_, err := o.Update(&zulinhetong)
	if zulinhetong.Sfsh == "是" {
		fangyuanxinxi := models.Fangyuanxinxi{}
		err := o.QueryTable(new(models.Fangyuanxinxi)).Filter("fangwumingcheng", zulinhetong.Fangwumingcheng).One(&fangyuanxinxi)
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

		if fangyuanxinxi.Fangwuzhuangtai == "未租" {
			// 更新租赁状态为已租
			fangyuanxinxi.Fangwuzhuangtai = "已租"
			o.Update(&fangyuanxinxi, "fangwuzhuangtai")

		}
	}
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "更新失败",
			Data: err,
		})
		return
	}

	// 更新成功
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功",
		Data: nil,
	})
}

// Delete 删除接口
func ZulinhetongControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	ids := make([]int, 0, 2)
	// 解析请求体
	if err := json.NewDecoder(c.Request.Body).Decode(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败",
			Data: err,
		})
		return
	}
	// 执行删除操作
	_, err := o.QueryTable("zulinhetong").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败",
			Data: err,
		})
		return
	}

	// 删除成功
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功",
		Data: nil,
	})
}

// Info 详情接口（后端）
func ZulinhetongControllerInfo(c *gin.Context) {
	o := orm.NewOrm()

	// 从 URL 参数中获取 id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "无效的 ID 参数",
			Data: err,
		})
		return
	}

	zulinhetong := models.Zulinhetong{Id: id}
	err = o.Read(&zulinhetong)

	if err != nil {
		if err == orm.ErrNoRows {
			c.JSON(404, models.ReturnMsg{
				Code: 404,
				Msg:  "未找到对应的租赁合同信息",
				Data: nil,
			})
		} else {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "服务器错误",
				Data: err,
			})
		}
		return
	}

	// 成功获取数据
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: zulinhetong,
	})
}

// Query 查询单条数据
func ZulinhetongControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	// 处理查询条件
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

// Detail 详情接口（前端）
func ZulinhetongControllerDetail(c *gin.Context) {
	// 调用 MD5 方法，目前硬编码参数，可按需调整
	utils.GetMd5("123456")

	o := orm.NewOrm()

	// 从 URL 参数中获取 id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "无效的 ID 参数",
			Data: err,
		})
		return
	}

	zulinhetong := models.Zulinhetong{Id: id}
	err = o.Read(&zulinhetong)

	if err != nil {
		if err == orm.ErrNoRows {
			c.JSON(404, models.ReturnMsg{
				Code: 404,
				Msg:  "未找到对应的租赁合同信息",
				Data: nil,
			})
		} else {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "服务器错误",
				Data: err,
			})
		}
		return
	}

	// 成功获取数据
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: zulinhetong,
	})
}

// Remind 获取需要提醒的记录数接口
func ZulinhetongControllerRemind(c *gin.Context) {
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
func ZulinhetongControllerGroup(c *gin.Context) {
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM fangwuzulin " + where + " GROUP BY " + columnName
	var maps []orm.Params
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
func ZulinhetongControllerValue(c *gin.Context) {
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

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangwuzulin " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
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
func ZulinhetongControllerValueTime(c *gin.Context) {
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

	timeStatType, err = url.QueryUnescape(timeStatType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: 400,
			Msg:  "时间统计类型参数解码失败",
			Data: err,
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
func ZulinhetongControllerValueMul(c *gin.Context) {
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

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM fangwuzulin " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
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
func ZulinhetongControllerCount(c *gin.Context) {
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

	cond := orm.NewCondition()
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

// ShBatch 批量审核接口
func ZulinhetongControllerShBatch(c *gin.Context) {
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

	sfsh := c.Query("sfsh")
	shhf := c.Query("shhf")

	// 执行批量更新操作
	num, err := o.QueryTable("zulinhetong").Filter("id__in", ids).Update(orm.Params{
		"sfsh": sfsh,
		"shhf": shhf,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: 500,
			Msg:  "审核失败",
			Data: err,
		})
		return
	}

	if num == 0 {
		c.JSON(http.StatusNotFound, models.ReturnMsg{
			Code: 404,
			Msg:  "未找到要审核的记录",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.ReturnMsg{
		Code: 0,
		Msg:  "审核成功！",
		Data: nil,
	})
}
