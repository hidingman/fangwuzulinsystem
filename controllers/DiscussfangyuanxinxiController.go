package controllers

import (
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"go-schema/models"
	"go-schema/utils"
)

func init() {
	orm.RegisterModel(new(models.Discussfangyuanxinxi))
}

// Page 分页接口（后端）
func DiscussfangyuanxinxiControllerPage(c *gin.Context) {
	token := c.GetHeader("Token")
	_, err := utils.ValidateToken(token)

	o := orm.NewOrm()
	var discussfangyuanxinxi []models.Discussfangyuanxinxi

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sortField := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := orm.NewCondition()
	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	avatarurlColumn := c.Query("avatarurl")
	if avatarurlColumn != "" {
		if strings.Contains(avatarurlColumn, "%") {
			cond = cond.And("avatarurl__contains", strings.Trim(avatarurlColumn, "%"))
		} else {
			cond = cond.And("avatarurl", avatarurlColumn)
		}
	}
	nicknameColumn := c.Query("nickname")
	if nicknameColumn != "" {
		if strings.Contains(nicknameColumn, "%") {
			cond = cond.And("nickname__contains", strings.Trim(nicknameColumn, "%"))
		} else {
			cond = cond.And("nickname", nicknameColumn)
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
	replyColumn := c.Query("reply")
	if replyColumn != "" {
		if strings.Contains(replyColumn, "%") {
			cond = cond.And("reply__contains", strings.Trim(replyColumn, "%"))
		} else {
			cond = cond.And("reply", replyColumn)
		}
	}
	thumbsupnumColumn := c.Query("thumbsupnum")
	if thumbsupnumColumn != "" {
		if strings.Contains(thumbsupnumColumn, "%") {
			cond = cond.And("thumbsupnum__contains", strings.Trim(thumbsupnumColumn, "%"))
		} else {
			cond = cond.And("thumbsupnum", thumbsupnumColumn)
		}
	}
	crazilynumColumn := c.Query("crazilynum")
	if crazilynumColumn != "" {
		if strings.Contains(crazilynumColumn, "%") {
			cond = cond.And("crazilynum__contains", strings.Trim(crazilynumColumn, "%"))
		} else {
			cond = cond.And("crazilynum", crazilynumColumn)
		}
	}
	istopColumn := c.Query("istop")
	if istopColumn != "" {
		if strings.Contains(istopColumn, "%") {
			cond = cond.And("istop__contains", strings.Trim(istopColumn, "%"))
		} else {
			cond = cond.And("istop", istopColumn)
		}
	}
	tuseridsColumn := c.Query("tuserids")
	if tuseridsColumn != "" {
		if strings.Contains(tuseridsColumn, "%") {
			cond = cond.And("tuserids__contains", strings.Trim(tuseridsColumn, "%"))
		} else {
			cond = cond.And("tuserids", tuseridsColumn)
		}
	}
	cuseridsColumn := c.Query("cuserids")
	if cuseridsColumn != "" {
		if strings.Contains(cuseridsColumn, "%") {
			cond = cond.And("cuserids__contains", strings.Trim(cuseridsColumn, "%"))
		} else {
			cond = cond.And("cuserids", cuseridsColumn)
		}
	}

	total, err := o.QueryTable("discussfangyuanxinxi").SetCond(cond).Count()
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
	sortlist := strings.Split(sortField, ",")

	for index, value := range sortlist {
		if index < len(orderlist) && orderlist[index] == "desc" {
			sortlist[index] = "-" + value
		}
	}
	_, err = o.QueryTable("discussfangyuanxinxi").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&discussfangyuanxinxi)
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
		List:      discussfangyuanxinxi,
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
func DiscussfangyuanxinxiControllerLists(c *gin.Context) {
	o := orm.NewOrm()
	var discussfangyuanxinxi []models.Discussfangyuanxinxi
	cond := orm.NewCondition()
	_, err := o.QueryTable("discussfangyuanxinxi").SetCond(cond).All(&discussfangyuanxinxi)
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
		Data: discussfangyuanxinxi,
	}
	c.JSON(200, res)
}

// List 分页接口（前端）
func DiscussfangyuanxinxiControllerList(c *gin.Context) {
	o := orm.NewOrm()
	var discussfangyuanxinxi []models.Discussfangyuanxinxi

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	sortField := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		order = "-" + sortField
	} else {
		order = sortField
	}

	cond := orm.NewCondition()
	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	avatarurlColumn := c.Query("avatarurl")
	if avatarurlColumn != "" {
		if strings.Contains(avatarurlColumn, "%") {
			cond = cond.And("avatarurl__contains", strings.Trim(avatarurlColumn, "%"))
		} else {
			cond = cond.And("avatarurl", avatarurlColumn)
		}
	}
	nicknameColumn := c.Query("nickname")
	if nicknameColumn != "" {
		if strings.Contains(nicknameColumn, "%") {
			cond = cond.And("nickname__contains", strings.Trim(nicknameColumn, "%"))
		} else {
			cond = cond.And("nickname", nicknameColumn)
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
	replyColumn := c.Query("reply")
	if replyColumn != "" {
		if strings.Contains(replyColumn, "%") {
			cond = cond.And("reply__contains", strings.Trim(replyColumn, "%"))
		} else {
			cond = cond.And("reply", replyColumn)
		}
	}
	thumbsupnumColumn := c.Query("thumbsupnum")
	if thumbsupnumColumn != "" {
		if strings.Contains(thumbsupnumColumn, "%") {
			cond = cond.And("thumbsupnum__contains", strings.Trim(thumbsupnumColumn, "%"))
		} else {
			cond = cond.And("thumbsupnum", thumbsupnumColumn)
		}
	}
	crazilynumColumn := c.Query("crazilynum")
	if crazilynumColumn != "" {
		if strings.Contains(crazilynumColumn, "%") {
			cond = cond.And("crazilynum__contains", strings.Trim(crazilynumColumn, "%"))
		} else {
			cond = cond.And("crazilynum", crazilynumColumn)
		}
	}
	istopColumn := c.Query("istop")
	if istopColumn != "" {
		if strings.Contains(istopColumn, "%") {
			cond = cond.And("istop__contains", strings.Trim(istopColumn, "%"))
		} else {
			cond = cond.And("istop", istopColumn)
		}
	}
	tuseridsColumn := c.Query("tuserids")
	if tuseridsColumn != "" {
		if strings.Contains(tuseridsColumn, "%") {
			cond = cond.And("tuserids__contains", strings.Trim(tuseridsColumn, "%"))
		} else {
			cond = cond.And("tuserids", tuseridsColumn)
		}
	}
	cuseridsColumn := c.Query("cuserids")
	if cuseridsColumn != "" {
		if strings.Contains(cuseridsColumn, "%") {
			cond = cond.And("cuserids__contains", strings.Trim(cuseridsColumn, "%"))
		} else {
			cond = cond.And("cuserids", cuseridsColumn)
		}
	}

	total, err := o.QueryTable("discussfangyuanxinxi").SetCond(cond).Count()
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
	_, err = o.QueryTable("discussfangyuanxinxi").SetCond(cond).OrderBy(order).Limit(limit, start).All(&discussfangyuanxinxi)
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
		List:      discussfangyuanxinxi,
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
func DiscussfangyuanxinxiControllerSave(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	o := orm.NewOrm()
	var discussfangyuanxinxi models.Discussfangyuanxinxi
	err := c.ShouldBindJSON(&discussfangyuanxinxi)
	if err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "请求参数解析失败",
			Data: err,
		})
		return
	}
	discussfangyuanxinxi.Userid = userInfo.Id
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&discussfangyuanxinxi)
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
func DiscussfangyuanxinxiControllerAdd(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
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
	var discussfangyuanxinxi models.Discussfangyuanxinxi
	err = c.ShouldBindJSON(&discussfangyuanxinxi)

	if discussfangyuanxinxi.Userid == 0 {
		discussfangyuanxinxi.Userid = userInfo.Id
	}
	o.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec()
	id, err := o.Insert(&discussfangyuanxinxi)
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
func DiscussfangyuanxinxiControllerUpdate(c *gin.Context) {
	o := orm.NewOrm()
	var discussfangyuanxinxi models.Discussfangyuanxinxi
	err := c.ShouldBindJSON(&discussfangyuanxinxi)
	_, err = o.Update(&discussfangyuanxinxi)
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "更新成功！",
		Data: err,
	})
}

// Delete 删除接口
func DiscussfangyuanxinxiControllerDelete(c *gin.Context) {
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
	_, err = o.QueryTable("discussfangyuanxinxi").Filter("id__in", ids).Delete()
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
func DiscussfangyuanxinxiControllerInfo(c *gin.Context) {
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
	discussfangyuanxinxi := models.Discussfangyuanxinxi{Id: id}
	err = o.Read(&discussfangyuanxinxi)

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
			Data: discussfangyuanxinxi,
		})
	}
}

// Query 查询单条数据
func DiscussfangyuanxinxiControllerQuery(c *gin.Context) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	refidColumn := c.Query("refid")
	if refidColumn != "" {
		if strings.Contains(refidColumn, "%") {
			cond = cond.And("refid__contains", strings.Trim(refidColumn, "%"))
		} else {
			cond = cond.And("refid", refidColumn)
		}
	}
	useridColumn := c.Query("userid")
	if useridColumn != "" {
		if strings.Contains(useridColumn, "%") {
			cond = cond.And("userid__contains", strings.Trim(useridColumn, "%"))
		} else {
			cond = cond.And("userid", useridColumn)
		}
	}
	avatarurlColumn := c.Query("avatarurl")
	if avatarurlColumn != "" {
		if strings.Contains(avatarurlColumn, "%") {
			cond = cond.And("avatarurl__contains", strings.Trim(avatarurlColumn, "%"))
		} else {
			cond = cond.And("avatarurl", avatarurlColumn)
		}
	}
	nicknameColumn := c.Query("nickname")
	if nicknameColumn != "" {
		if strings.Contains(nicknameColumn, "%") {
			cond = cond.And("nickname__contains", strings.Trim(nicknameColumn, "%"))
		} else {
			cond = cond.And("nickname", nicknameColumn)
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
	replyColumn := c.Query("reply")
	if replyColumn != "" {
		if strings.Contains(replyColumn, "%") {
			cond = cond.And("reply__contains", strings.Trim(replyColumn, "%"))
		} else {
			cond = cond.And("reply", replyColumn)
		}
	}
	thumbsupnumColumn := c.Query("thumbsupnum")
	if thumbsupnumColumn != "" {
		if strings.Contains(thumbsupnumColumn, "%") {
			cond = cond.And("thumbsupnum__contains", strings.Trim(thumbsupnumColumn, "%"))
		} else {
			cond = cond.And("thumbsupnum", thumbsupnumColumn)
		}
	}
	crazilynumColumn := c.Query("crazilynum")
	if crazilynumColumn != "" {
		if strings.Contains(crazilynumColumn, "%") {
			cond = cond.And("crazilynum__contains", strings.Trim(crazilynumColumn, "%"))
		} else {
			cond = cond.And("crazilynum", crazilynumColumn)
		}
	}
	istopColumn := c.Query("istop")
	if istopColumn != "" {
		if strings.Contains(istopColumn, "%") {
			cond = cond.And("istop__contains", strings.Trim(istopColumn, "%"))
		} else {
			cond = cond.And("istop", istopColumn)
		}
	}
	tuseridsColumn := c.Query("tuserids")
	if tuseridsColumn != "" {
		if strings.Contains(tuseridsColumn, "%") {
			cond = cond.And("tuserids__contains", strings.Trim(tuseridsColumn, "%"))
		} else {
			cond = cond.And("tuserids", tuseridsColumn)
		}
	}
	cuseridsColumn := c.Query("cuserids")
	if cuseridsColumn != "" {
		if strings.Contains(cuseridsColumn, "%") {
			cond = cond.And("cuserids__contains", strings.Trim(cuseridsColumn, "%"))
		} else {
			cond = cond.And("cuserids", cuseridsColumn)
		}
	}
	var discussfangyuanxinxi models.Discussfangyuanxinxi
	err := o.QueryTable("discussfangyuanxinxi").SetCond(cond).One(&discussfangyuanxinxi)
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
		Data: discussfangyuanxinxi,
	})
}

// Detail 详情接口（前端）
func DiscussfangyuanxinxiControllerDetail(c *gin.Context) {
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
	discussfangyuanxinxi := models.Discussfangyuanxinxi{Id: id}
	err = o.Read(&discussfangyuanxinxi)

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
			Data: discussfangyuanxinxi,
		})
	}
}

// Remind 获取需要提醒的记录数接口
func DiscussfangyuanxinxiControllerRemind(c *gin.Context) {
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
			sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			if start > end {
				sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' OR " + columnName + " <= '" + remindend + "'"
			} else {
				sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
			}
		}
	}

	if typeP == "2" {
		if remindstart != "" {
			start, _ := strconv.Atoi(remindstart)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " >= '" + remindstart + "'"
		}
		if remindend != "" {
			end, _ := strconv.Atoi(remindend)
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " <= '" + remindend + "'"
		}
		if remindstart != "" && remindend != "" {
			start, _ := strconv.Atoi(remindstart)
			end, _ := strconv.Atoi(remindend)
			remindstart = utils.GetDateTimeFormat(start, "yyyy-MM-dd")
			remindend = utils.GetDateTimeFormat(end, "yyyy-MM-dd")
			sql = "SELECT  FROM discussfangyuanxinxi WHERE " + where + " AND " + columnName + " >= '" + remindstart + "' AND " + columnName + " <= '" + remindend + "'"
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
func DiscussfangyuanxinxiControllerGroup(c *gin.Context) {
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
	sql := "SELECT COUNT(*) AS total, " + columnName + " FROM discussfangyuanxinxi WHERE 1 = 1 GROUP BY " + columnName
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
func DiscussfangyuanxinxiControllerValue(c *gin.Context) {
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
	if "discussfangyuanxinxi" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	sql := "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM discussfangyuanxinxi " + where + " GROUP BY " + xColumnName
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
func DiscussfangyuanxinxiControllerValueTime(c *gin.Context) {
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
	if "discussfangyuanxinxi" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}
	timeStatType, _ = url.QueryUnescape(timeStatType)

	var sql string
	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM discussfangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM discussfangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM discussfangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
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
func DiscussfangyuanxinxiControllerValueMul(c *gin.Context) {
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
	if "discussfangyuanxinxi" == "orders" {
		where += " AND status IN ('已支付', '已发货', '已完成') "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql := "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM discussfangyuanxinxi " + where + " GROUP BY " + xColumnName + " LIMIT 10"
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
