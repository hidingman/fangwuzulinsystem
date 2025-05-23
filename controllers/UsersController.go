package controllers

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
)

func init() {
	orm.RegisterModel(new(models.Users))
}

// UsersControllerRegister 用户注册接口
func UsersControllerRegister(c *gin.Context) {
	o := orm.NewOrm()
	users := models.Users{}
	// 解析请求体并处理可能的错误
	if err := json.NewDecoder(c.Request.Body).Decode(&users); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: -1,
			Msg:  "请求体解析失败！",
			Data: nil,
		})
		return
	}
	users.Role = "管理员"
	_, err := o.Insert(&users)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "注册失败！",
			Data: nil,
		})
		return
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "注册成功！",
		Data: nil,
	})
}

// UsersControllerLogin 用户登录接口
func UsersControllerLogin(c *gin.Context) {
	o := orm.NewOrm()
	users := models.Users{}
	username := c.Query("username")
	password := c.Query("password")
	err := o.QueryTable("users").Filter("username", username).Filter("password", password).One(&users)
	var res interface{}
	if err != nil {
		c.JSON(401, models.ReturnMsg{
			Code: -1,
			Msg:  "用户名或密码错误！",
			Data: nil,
		})
		return
	}
	session := models.Session{
		Id:        users.Id,
		Username:  users.Username,
		Tablename: "users",
		Role:      "管理员",
	}
	token, err := utils.GenerateToken(&session, 0)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "生成令牌失败！",
			Data: nil,
		})
		return
	}
	res = models.ReturnToken{
		Code:  0,
		Token: token,
	}

	c.JSON(200, res)
}

// UsersControllerLogout 用户退出接口
func UsersControllerLogout(c *gin.Context) {
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "退出成功！",
		Data: nil,
	})
}

// UsersControllerSession 获取 session 的接口
func UsersControllerSession(c *gin.Context) {
	o := orm.NewOrm()
	// 从请求头中获取 Token
	token := c.GetHeader("Token")
	if token == "" {
		c.AbortWithStatusJSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	// 验证 Token
	userinfo, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}
	// 根据用户 ID 查询用户信息
	users := models.Users{Id: userinfo.Id}
	if err := o.Read(&users); err != nil {
		c.AbortWithStatusJSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询用户信息失败！",
			Data: nil,
		})
		return
	}
	// 返回成功响应
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: users,
	})
}

// UsersControllerPage 分页接口（后端）
func UsersControllerPage(c *gin.Context) {
	o := orm.NewOrm()
	var users []models.Users

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
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

	username := c.Query("username")
	cond := orm.NewCondition()
	if username != "" {
		cond = cond.And("username__contains", username)
	}

	total, err := o.QueryTable("users").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "获取总数失败！",
			Data: nil,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit
	// 查询用户列表并处理可能的错误
	if _, err := o.QueryTable("users").SetCond(cond).OrderBy(order).Limit(limit, start).All(&users); err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "查询用户列表失败！",
			Data: nil,
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      users,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	})
}

// UsersControllerSave 保存接口
func UsersControllerSave(c *gin.Context) {
	o := orm.NewOrm()
	users := models.Users{}
	// 解析请求体并处理可能的错误
	if err := json.NewDecoder(c.Request.Body).Decode(&users); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: -1,
			Msg:  "请求体解析失败！",
			Data: nil,
		})
		return
	}
	_, err := o.Insert(&users)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "添加失败！",
			Data: err,
		})
		return
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "添加成功！",
		Data: nil,
	})
}

// UsersControllerUpdate 更新接口
func UsersControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	users := models.Users{}
	// 解析请求体并处理可能的错误
	if err := json.NewDecoder(c.Request.Body).Decode(&users); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: -1,
			Msg:  "请求体解析失败！",
			Data: nil,
		})
		return
	}
	_, err := o.Update(&users)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "更新失败！",
			Data: err,
		})
		return
	}
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: nil,
	})
}

// UsersControllerDelete 删除接口
func UsersControllerDelete(c *gin.Context) {
	o := orm.NewOrm()
	var ids []int
	// 解析请求体并处理可能的错误
	if err := json.NewDecoder(c.Request.Body).Decode(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: -1,
			Msg:  "请求体解析失败！",
			Data: nil,
		})
		return
	}
	_, err := o.QueryTable("users").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: -1,
			Msg:  "删除失败！",
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

// UsersControllerInfo 详情接口
func UsersControllerInfo(c *gin.Context) {
	o := orm.NewOrm()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: -1,
			Msg:  "无效的 ID！",
			Data: nil,
		})
		return
	}
	users := models.Users{Id: id}
	// 读取用户信息并处理可能的错误
	if err := o.Read(&users); err != nil {
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
		Data: users,
	})
}
