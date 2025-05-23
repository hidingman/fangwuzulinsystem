package controllers

import (
	"fmt"
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
	orm.RegisterModel(new(models.Yuangong))
}

// 用户注册接口
func YuangongRegister(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong models.Yuangong

	// 解析请求体
	if err := c.ShouldBindJSON(&yuangong); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 临时禁用外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "禁用外键检查失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 插入数据
	_, err := o.Insert(&yuangong)

	// 恢复外键检查
	if _, errExec := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); errExec != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "恢复外键检查失败: " + errExec.Error(),
			Data: nil,
		})
		return
	}

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
func YuangongLogin(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong models.Yuangong
	password := c.Query("password")

	err := o.QueryTable("yuangong").Filter("gonghao", c.Query("username")).Filter("mima", password).One(&yuangong)
	var res interface{}
	if err != nil {
		res = models.ReturnMsg{
			Code: -1,
			Msg:  "用户名或密码错误！",
			Data: nil,
		}
	} else {
		session := models.Session{
			Id:              yuangong.Id,
			Username:        yuangong.Gonghao,
			Tablename:       "yuangong",
			Role:            "员工",
			LoginUserColumn: "gonghao",
		}
		token, err := utils.GenerateToken(&session, 0)
		if err != nil {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "生成 Token 失败: " + err.Error(),
				Data: nil,
			})
			return
		}
		res = models.ReturnToken{
			Code:  0,
			Token: token,
		}
	}

	c.JSON(200, res)
}

// 用户退出接口
func YuangongLogout(c *gin.Context) {
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "退出成功！",
		Data: nil,
	})
}

// 获取session的接口
func YuangongSession(c *gin.Context) {
	o := orm.NewOrm()
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	userinfo, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	yuangong := models.Yuangong{Id: userinfo.Id}
	err = o.Read(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: yuangong,
	})
}

// 忘记密码（找回密码）
func YuangongResetPass(c *gin.Context) {
	o := orm.NewOrm()
	username := c.Query("username")
	pwd := "123456"

	// 执行更新操作并处理错误
	res, err := o.Raw("UPDATE yuangong SET mima = ? WHERE gonghao = ?", pwd, username).Exec()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "密码重置失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 检查是否有记录被更新
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "获取更新记录数失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(404, models.ReturnMsg{
			Code: 404,
			Msg:  "未找到该用户，密码重置失败",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "密码已重置为：123456",
		Data: nil,
	})
}

// 分页接口（后端）
func YuangongPage(c *gin.Context) {
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
	var yuangong []models.Yuangong

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	// 避免与导入包冲突，修改变量名
	sortStr := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := orm.NewCondition()
	// 定义查询字段映射
	fields := map[string]string{
		"gonghao":          c.Query("gonghao"),
		"mima":             c.Query("mima"),
		"yuangongxingming": c.Query("yuangongxingming"),
		"xingbie":          c.Query("xingbie"),
		"touxiang":         c.Query("touxiang"),
		"shenfenzheng":     c.Query("shenfenzheng"),
		"dianhua":          c.Query("dianhua"),
		"minzu":            c.Query("minzu"),
		"jiguan":           c.Query("jiguan"),
		"chushengriyue":    c.Query("chushengriyue"),
		"nianling":         c.Query("nianling"),
		"xueli":            c.Query("xueli"),
		"zhuzhi":           c.Query("zhuzhi"),
		"tezhang":          c.Query("tezhang"),
		"ziwopingjia":      c.Query("ziwopingjia"),
		"gongzuoanpai":     c.Query("gongzuoanpai"),
		"gongzixinxi":      c.Query("gongzixinxi"),
		"jifen":            c.Query("jifen"),
	}

	// 构建查询条件
	for field, value := range fields {
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(fmt.Sprintf("%s__contains", field), strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	// 计算总记录数并处理错误
	total, err := o.QueryTable("yuangong").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总记录数失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	orderlist := strings.Split(order, ",")
	sortlist := strings.Split(sortStr, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}

	// 查询数据并处理错误
	_, err = o.QueryTable("yuangong").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      yuangong,
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
func YuangongLists(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong []models.Yuangong
	cond := orm.NewCondition()

	// 执行查询并处理错误
	_, err := o.QueryTable("yuangong").SetCond(cond).All(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询员工信息失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: yuangong,
	}
	c.JSON(200, res)
}

// 分页接口（前端）
func YuangongList(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong []models.Yuangong

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	// 避免与导入包冲突，修改变量名
	sortStr := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sortStr
	} else {
		order = sortStr
	}

	cond := orm.NewCondition()
	// 定义查询字段映射
	fields := map[string]string{
		"gonghao":          c.Query("gonghao"),
		"mima":             c.Query("mima"),
		"yuangongxingming": c.Query("yuangongxingming"),
		"xingbie":          c.Query("xingbie"),
		"touxiang":         c.Query("touxiang"),
		"shenfenzheng":     c.Query("shenfenzheng"),
		"dianhua":          c.Query("dianhua"),
		"minzu":            c.Query("minzu"),
		"jiguan":           c.Query("jiguan"),
		"chushengriyue":    c.Query("chushengriyue"),
		"nianling":         c.Query("nianling"),
		"xueli":            c.Query("xueli"),
		"zhuzhi":           c.Query("zhuzhi"),
		"tezhang":          c.Query("tezhang"),
		"ziwopingjia":      c.Query("ziwopingjia"),
		"gongzuoanpai":     c.Query("gongzuoanpai"),
		"gongzixinxi":      c.Query("gongzixinxi"),
		"jifen":            c.Query("jifen"),
	}

	// 构建查询条件
	for field, value := range fields {
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(fmt.Sprintf("%s__contains", field), strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	// 计算总记录数并处理错误
	total, err := o.QueryTable("yuangong").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总记录数失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit

	// 查询数据并处理错误
	_, err = o.QueryTable("yuangong").SetCond(cond).OrderBy(order).Limit(limit, start).All(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      yuangong,
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
func YuangongSave(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong models.Yuangong

	// 解析请求体
	if err := c.ShouldBindJSON(&yuangong); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 临时禁用外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "禁用外键检查失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&yuangong)

	// 恢复外键检查
	if _, errExec := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); errExec != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "恢复外键检查失败: " + errExec.Error(),
			Data: nil,
		})
		return
	}

	if err != nil {
		c.JSON(402, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！" + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "添加成功！",
		Data: id,
	})
}

// 保存接口（前端）
func YuangongAdd(c *gin.Context) {
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
	var yuangong models.Yuangong

	// 解析请求体
	if err := c.ShouldBindJSON(&yuangong); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 临时禁用外键检查
	if _, err := o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "禁用外键检查失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 插入数据
	id, err := o.Insert(&yuangong)

	// 恢复外键检查
	if _, errExec := o.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(); errExec != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "恢复外键检查失败: " + errExec.Error(),
			Data: nil,
		})
		return
	}

	if err != nil {
		c.JSON(402, models.ReturnMsg{
			Code: 402,
			Msg:  "添加失败！" + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "添加成功！",
		Data: id,
	})
}

// 更新接口
func YuangongUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var yuangong models.Yuangong

	// 解析请求体
	if err := c.ShouldBindJSON(&yuangong); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 更新数据
	_, err := o.Update(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "更新失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: nil,
	})
}

// 删除接口
func YuangongDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int

	// 解析请求体
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求体解析失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 删除数据
	_, err := o.QueryTable("yuangong").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "删除失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "删除成功！",
		Data: nil,
	})
}

// 详情接口（后端）
func YuangongInfo(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "ID 参数错误: " + err.Error(),
			Data: nil,
		})
		return
	}

	yuangong := models.Yuangong{Id: id}
	err = o.Read(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: yuangong,
	})
}

// 查询单条数据
func YuangongQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	// 定义查询字段
	fields := map[string]string{
		"gonghao":          c.Query("gonghao"),
		"mima":             c.Query("mima"),
		"yuangongxingming": c.Query("yuangongxingming"),
		"xingbie":          c.Query("xingbie"),
		"touxiang":         c.Query("touxiang"),
		"shenfenzheng":     c.Query("shenfenzheng"),
		"dianhua":          c.Query("dianhua"),
		"minzu":            c.Query("minzu"),
		"jiguan":           c.Query("jiguan"),
		"chushengriyue":    c.Query("chushengriyue"),
		"nianling":         c.Query("nianling"),
		"xueli":            c.Query("xueli"),
		"zhuzhi":           c.Query("zhuzhi"),
		"tezhang":          c.Query("tezhang"),
		"ziwopingjia":      c.Query("ziwopingjia"),
		"gongzuoanpai":     c.Query("gongzuoanpai"),
		"gongzixinxi":      c.Query("gongzixinxi"),
		"jifen":            c.Query("jifen"),
	}

	// 构建查询条件
	for field, value := range fields {
		if value != "" {
			if strings.Contains(value, "%") {
				cond = cond.And(fmt.Sprintf("%s__contains", field), strings.Trim(value, "%"))
			} else {
				cond = cond.And(field, value)
			}
		}
	}

	var yuangong models.Yuangong
	err := o.QueryTable("yuangong").SetCond(cond).One(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: yuangong,
	})
}

// 详情接口（前端）
func YuangongDetail(c *gin.Context) {
	utils.GetMd5("123456")
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "ID 参数错误: " + err.Error(),
			Data: nil,
		})
		return
	}

	yuangong := models.Yuangong{Id: id}
	err = o.Read(&yuangong)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "服务器错误！",
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: yuangong,
	})
}

// 获取需要提醒的记录数接口
func YuangongRemind(c *gin.Context) {
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
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, err := strconv.Atoi(remindstart)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindstart 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			end, err := strconv.Atoi(remindend)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindend 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			if start > end {
				sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, err := strconv.Atoi(remindstart)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindstart 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if c.Query("remindend") != "" {
			end, err := strconv.Atoi(remindend)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindend 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if c.Query("remindstart") != "" && c.Query("remindend") != "" {
			start, err := strconv.Atoi(remindstart)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindstart 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			end, err := strconv.Atoi(remindend)
			if err != nil {
				c.JSON(400, models.ReturnMsg{
					Code: 400,
					Msg:  "remindend 参数转换错误: " + err.Error(),
					Data: nil,
				})
				return
			}
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT xingming FROM guke WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
		}
	}

	o := orm.NewOrm()
	var result []string
	// 执行 SQL 查询并处理错误
	if _, err := o.Raw(sql).QueryRows(&result); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

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
func YuangongGroup(c *gin.Context) {
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
	sql := ""
	where := " WHERE 1 = 1 "

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM guke " + where + " GROUP BY " + columnName
	var maps []orm.Params
	// 执行 SQL 查询并处理错误
	if _, err := o.Raw(sql).Values(&maps); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")

	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			return totalI > totalJ
		})
	} else if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
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
func YuangongValue(c *gin.Context) {
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
	sql := ""
	where := " WHERE 1 = 1 "

	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM guke " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	// 执行 SQL 查询并处理错误
	if _, err := o.Raw(sql).Values(&maps); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")

	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			return totalI > totalJ
		})
	} else if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
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
func YuangongValueTime(c *gin.Context) {
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
	sql := ""
	where := " WHERE 1 = 1 "

	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM guke " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	// 执行 SQL 查询并处理错误
	if _, err := o.Raw(sql).Values(&maps); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	order := c.Query("order")

	if order == "desc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			return totalI > totalJ
		})
	} else if order == "asc" {
		sort.Slice(maps, func(i, j int) bool {
			totalI, err := strconv.Atoi(maps[i]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalI 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
			totalJ, err := strconv.Atoi(maps[j]["total"].(string))
			if err != nil {
				c.JSON(500, models.ReturnMsg{
					Code: 500,
					Msg:  "转换 totalJ 失败: " + err.Error(),
					Data: nil,
				})
				return false
			}
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
func YuangongValueMul(c *gin.Context) {
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

	sql := ""
	where := " WHERE 1 = 1 "
	if "guke" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM guke " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		// 执行 SQL 查询并处理错误
		if _, err := o.Raw(sql).Values(&maps); err != nil {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "查询失败: " + err.Error(),
				Data: nil,
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

// 总数接口
func YuangongCount(c *gin.Context) {
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
	// 执行计数查询并处理错误
	cnt, err := o.QueryTable("guke").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总数失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: cnt,
	})
}
