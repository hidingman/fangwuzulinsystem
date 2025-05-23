package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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
	orm.RegisterModel(new(models.Fangyuanxinxi))
}

// 分页接口（后端）
func FangyuanxinxiPage(c *gin.Context) {
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
	if valiErr != nil {
		c.JSON(401, models.ReturnMsg{
			Code: 401,
			Msg:  "您的权限不够！",
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	var fangyuanxinxi []models.Fangyuanxinxi

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")

	cond := orm.NewCondition()

	// 处理 zujinjiagestart 和 zujinjiageend
	zujinjiagestart := c.Query("zujinjiagestart")
	zujinjiageend := c.Query("zujinjiageend")
	if zujinjiagestart != "" {
		cond = cond.And("zujinjiage__gte", zujinjiagestart)
	}
	if zujinjiageend != "" {
		cond = cond.And("zujinjiage__lte", zujinjiageend)
	}

	// 封装查询条件处理函数
	handleQueryCondition := func(cond *orm.Condition, columnName, param string) *orm.Condition {
		if param != "" {
			if strings.Contains(param, "%") {
				cond = cond.And(columnName+"__contains", strings.Trim(param, "%"))
			} else {
				cond = cond.And(columnName, param)
			}
		}
		return cond
	}

	// 处理其他查询条件
	columns := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "fangwudizhi", "mianji",
		"zujinjiage", "yajinfangshi", "fukuanfangshi", "fangwuchaoxiang", "yajin",
		"xiaoqu", "fangwushipin", "loudongdanyuan", "fanghao", "jifen",
		"fangwujiegou", "fangwuzhuangtai", "fangchanzhengbianhao", "fangchanzhengzhaopian",
		"fangzhuxingming", "fangzhushenfenzheng", "fangzhudianhua", "fangwuxiangqing",
		"fabushijian", "gonghao", "yuangongxingming", "thumbsupnum", "crazilynum",
		"clicktime", "clicknum", "discussnum", "storeupnum",
	}
	for _, col := range columns {
		param := c.Query(col)
		cond = handleQueryCondition(cond, col, param)
	}

	tableName := userInfo.Tablename
	if tableName == "yuangong" {
		cond = cond.And("gonghao", userInfo.Username)
	}

	total, err := o.QueryTable("fangyuanxinxi").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总数失败: " + err.Error(),
			Data: nil,
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

	_, err = o.QueryTable("fangyuanxinxi").SetCond(cond).OrderBy(sortlist...).Limit(limit, start).All(&fangyuanxinxi)
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
		List:      fangyuanxinxi,
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
func FangyuanxinxiLists(c *gin.Context) {
	o := orm.NewOrm()
	var fangyuanxinxi []models.Fangyuanxinxi

	cond := orm.NewCondition()

	_, err := o.QueryTable("fangyuanxinxi").SetCond(cond).All(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangyuanxinxi,
	}
	c.JSON(200, res)
}

// 分页接口（前端）
func FangyuanxinxiList(c *gin.Context) {
	o := orm.NewOrm()
	var fangyuanxinxi []models.Fangyuanxinxi

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	sort := c.DefaultQuery("sort", "id")
	order := c.Query("order")
	if order == "desc" {
		sort = "-" + sort
	}

	cond := orm.NewCondition()

	// 处理 zujinjiagestart 和 zujinjiageend
	zujinjiagestart := c.Query("zujinjiagestart")
	zujinjiageend := c.Query("zujinjiageend")
	if zujinjiagestart != "" {
		cond = cond.And("zujinjiage__gte", zujinjiagestart)
	}
	if zujinjiageend != "" {
		cond = cond.And("zujinjiage__lte", zujinjiageend)
	}

	// 封装查询条件处理函数
	handleQueryCondition := func(cond *orm.Condition, columnName, param string) *orm.Condition {
		if param != "" {
			if strings.Contains(param, "%") {
				cond = cond.And(columnName+"__contains", strings.Trim(param, "%"))
			} else {
				cond = cond.And(columnName, param)
			}
		}
		return cond
	}

	// 处理其他查询条件
	columns := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "fangwudizhi", "mianji",
		"zujinjiage", "yajinfangshi", "fukuanfangshi", "fangwuchaoxiang", "yajin",
		"xiaoqu", "fangwushipin", "loudongdanyuan", "fanghao", "jifen",
		"fangwujiegou", "fangwuzhuangtai", "fangchanzhengbianhao", "fangchanzhengzhaopian",
		"fangzhuxingming", "fangzhushenfenzheng", "fangzhudianhua", "fangwuxiangqing",
		"fabushijian", "gonghao", "yuangongxingming", "thumbsupnum", "crazilynum",
		"clicktime", "clicknum", "discussnum", "storeupnum",
	}
	for _, col := range columns {
		param := c.Query(col)
		cond = handleQueryCondition(cond, col, param)
	}

	total, err := o.QueryTable("fangyuanxinxi").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总数失败: " + err.Error(),
			Data: nil,
		})
		return
	}
	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit

	_, err = o.QueryTable("fangyuanxinxi").SetCond(cond).OrderBy(sort).Limit(limit, start).All(&fangyuanxinxi)
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
		List:      fangyuanxinxi,
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
func FangyuanxinxiSave(c *gin.Context) {
	var fangyuanxinxi models.Fangyuanxinxi
	if err := c.ShouldBindJSON(&fangyuanxinxi); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "保存数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "保存成功！",
		Data: nil,
	})
}

// 保存接口（前端）
func FangyuanxinxiAdd(c *gin.Context) {
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

	var fangyuanxinxi models.Fangyuanxinxi
	if err := c.ShouldBindJSON(&fangyuanxinxi); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "保存数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "保存成功！",
		Data: nil,
	})
}

// 更新接口
func FangyuanxinxiUpdate(c *gin.Context) {
	var fangyuanxinxi models.Fangyuanxinxi
	if err := c.ShouldBindJSON(&fangyuanxinxi); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "解析请求体失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	o := orm.NewOrm()
	_, err := o.Update(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "更新数据失败: " + err.Error(),
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
func FangyuanxinxiDelete(c *gin.Context) {
	o := orm.NewOrm()
	ids := make([]int, 0, 2)
	// 解析请求体
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, models.ReturnMsg{
			Code: 1,
			Data: "请求体解析失败",
		})
		return
	}

	// 检查 ID 列表是否为空
	if len(ids) == 0 {
		c.JSON(400, models.ReturnMsg{
			Code: 1,
			Data: "ID 列表不能为空",
		})
		return
	}

	// 先删除 storeup 表中的关联记录
	_, err := o.QueryTable("storeup").Filter("refid__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 1,
			Data: "删除关联数据失败",
		})
		return
	}

	// 执行删除操作
	_, err = o.QueryTable("fangyuanxinxi").Filter("id__in", ids).Delete()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 1,
			Data: "删除失败",
		})
		return
	}

	// 返回成功响应
	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Data: "删除成功！",
	})
}

// 详情接口（后端）
func FangyuanxinxiInfo(c *gin.Context) {
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

	o := orm.NewOrm()
	fangyuanxinxi := models.Fangyuanxinxi{Id: id}
	err = o.Read(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangyuanxinxi,
	})
}

// 查询单条数据
func FangyuanxinxiQuery(c *gin.Context) {
	o := orm.NewOrm()
	var fangyuanxinxi models.Fangyuanxinxi

	cond := orm.NewCondition()

	// 处理 zujinjiagestart 和 zujinjiageend
	zujinjiagestart := c.Query("zujinjiagestart")
	zujinjiageend := c.Query("zujinjiageend")
	if zujinjiagestart != "" {
		cond = cond.And("zujinjiage__gte", zujinjiagestart)
	}
	if zujinjiageend != "" {
		cond = cond.And("zujinjiage__lte", zujinjiageend)
	}

	// 封装查询条件处理函数
	handleQueryCondition := func(cond *orm.Condition, columnName, param string) *orm.Condition {
		if param != "" {
			if strings.Contains(param, "%") {
				cond = cond.And(columnName+"__contains", strings.Trim(param, "%"))
			} else {
				cond = cond.And(columnName, param)
			}
		}
		return cond
	}

	// 处理其他查询条件
	columns := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "fangwudizhi", "mianji",
		"zujinjiage", "yajinfangshi", "fukuanfangshi", "fangwuchaoxiang", "yajin",
		"xiaoqu", "fangwushipin", "loudongdanyuan", "fanghao", "jifen",
		"fangwujiegou", "fangwuzhuangtai", "fangchanzhengbianhao", "fangchanzhengzhaopian",
		"fangzhuxingming", "fangzhushenfenzheng", "fangzhudianhua", "fangwuxiangqing",
		"fabushijian", "gonghao", "yuangongxingming", "thumbsupnum", "crazilynum",
		"clicktime", "clicknum", "discussnum", "storeupnum",
	}
	for _, col := range columns {
		param := c.Query(col)
		cond = handleQueryCondition(cond, col, param)
	}

	err := o.QueryTable("fangyuanxinxi").SetCond(cond).One(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangyuanxinxi,
	})
}

// 详情接口（前端）
func FangyuanxinxiDetail(c *gin.Context) {
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

	o := orm.NewOrm()
	fangyuanxinxi := models.Fangyuanxinxi{Id: id}
	err = o.Read(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询数据失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: fangyuanxinxi,
	})
}

// 获取需要提醒的记录数接口
func FangyuanxinxiRemind(c *gin.Context) {
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

	// 此处需要实现具体的提醒记录数查询逻辑
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: 0,
	}
	c.JSON(200, res)
}

// 赞、踩接口
func FangyuanxinxiThumbsup(c *gin.Context) {
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

	// 此处需要实现具体的赞、踩逻辑
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: nil,
	}
	c.JSON(200, res)
}

// 智能推荐接口
func FangyuanxinxiAutoSort(c *gin.Context) {
	o := orm.NewOrm()
	var fangyuanxinxi []models.Fangyuanxinxi

	// 获取分页参数
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	// 获取排序参数
	sort := c.Query("sort")
	if sort == "" {
		sort = "clicktime"
	}

	order := c.Query("order")
	if order == "desc" {
		order = sort
	} else {
		order = "-" + sort
	}

	sort = "-clicknum"
	order = sort
	gin.DefaultWriter.Write([]byte(fmt.Sprintf("Sort: %s, Order: %s\n", sort, order)))
	cond := orm.NewCondition()
	sfsh := c.Query("sfsh")
	if sfsh != "" {
		cond = cond.And("sfsh", sfsh)
	}

	total, err := o.QueryTable("fangyuanxinxi").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Data: "",
		})
		return
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	start := (page - 1) * limit

	// 查询数据并处理错误
	_, err = o.QueryTable("fangyuanxinxi").SetCond(cond).OrderBy(order).Limit(limit, start).All(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Data: "",
		})
		return
	}

	returnPage := models.ReturnPage{
		CurrPage:  page,
		List:      fangyuanxinxi,
		PageSize:  limit,
		Total:     int(total),
		TotalPage: int(totalPage),
	}

	res := models.ReturnMsg{
		Code: 0,
		Data: returnPage,
	}

	c.JSON(200, res)
}

// 智能推荐接口
func FangyuanxinxiAutoSort2(c *gin.Context) {
	token := c.GetHeader("Token")
	userInfo, _ := utils.ValidateToken(token)

	var idListStr string
	o := orm.NewOrm()
	client := utils.NewCozeClient(utils.BaseURL, utils.ApiKey)

	payload := map[string]interface{}{
		"workflow_id": "",
		"parameters": map[string]string{
			"input": strconv.Itoa(userInfo.Id),
		},
	}

	endpoint := "/v1/workflow/run"

	body, err := client.SendRequest(endpoint, payload)
	if err != nil {
		log.Printf("请求出错: %v\n", err)
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "请求出错: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 定义一个 map 来存储解析后的 JSON 数据
	var results map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Printf("解析 JSON 出错: %v\n", err)
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "解析 JSON 出错: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 尝试取出 data 字段
	if dataStr, ok := results["data"].(string); ok {
		var innerData map[string]interface{}
		err = json.Unmarshal([]byte(dataStr), &innerData)
		if err != nil {
			log.Printf("解析内层 JSON 出错: %v\n", err)
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "解析内层 JSON 出错: " + err.Error(),
				Data: nil,
			})
			return
		}
		idListStr = innerData["data"].(string)
	} else {
		log.Println("响应中不存在 data 字段")
		c.JSON(400, models.ReturnMsg{
			Code: 400,
			Msg:  "响应中不存在 data 字段",
			Data: nil,
		})
		return
	}

	idListStr = strings.Trim(idListStr, "[]")
	idStrs := strings.Split(idListStr, ",")
	var idList []int
	for _, idStr := range idStrs {
		// 去除空格
		idStr = strings.TrimSpace(idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("解析 ID 失败:", err)
			res := models.ReturnMsg{
				Code: 1,
				Msg:  "解析 ID 失败",
				Data: nil,
			}
			c.JSON(400, res)
			return
		}
		idList = append(idList, id)
	}

	var fangyuanxinxi []models.Fangyuanxinxi
	var result []models.Fangyuanxinxi
	// 处理查询错误
	_, err = o.QueryTable("fangyuanxinxi").Filter("ID__in", idList).All(&fangyuanxinxi)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询房源信息失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	for index, value := range fangyuanxinxi {
		if value.Id != idList[index] {
			for i := 0; i < len(idList); i++ {
				if idList[index] == fangyuanxinxi[i].Id {
					fangyuanxinxi[i], fangyuanxinxi[index] = fangyuanxinxi[index], fangyuanxinxi[i]
					break
				}
			}
		}
	}
	result = fangyuanxinxi
	returnPage := models.ReturnPage{
		CurrPage:  1,
		List:      result,
		PageSize:  1,
		Total:     len(result),
		TotalPage: 1,
	}
	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: returnPage,
	}
	c.JSON(200, res)
}

// 分组统计接口
func FangyuanxinxiGroup(c *gin.Context) {
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

	sql = "SELECT COUNT(*) AS total, " + columnName + " FROM fangyuanxinxi " + where + " GROUP BY " + columnName
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
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

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	}
	c.JSON(200, res)
}

// 统计指定字段
func FangyuanxinxiValue(c *gin.Context) {
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

	// 示例条件，可根据实际需求修改
	if "fangyuanxinxi" == "someCondition" {
		where += " AND someField = 'someValue' "
	}

	sql = "SELECT " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangyuanxinxi " + where + " GROUP BY " + xColumnName
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
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

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	}
	c.JSON(200, res)
}

// 按日期统计
func FangyuanxinxiValueTime(c *gin.Context) {
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

	// 示例条件，可根据实际需求修改
	if "fangyuanxinxi" == "someCondition" {
		where += " AND someField = 'someValue' "
	}

	timeStatType, _ = url.QueryUnescape(timeStatType)

	switch timeStatType {
	case "日":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') AS " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	case "月":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') AS " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	case "年":
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') AS " + xColumnName + ", SUM(" + yColumnName + ") AS total FROM fangyuanxinxi " + where + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
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

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: maps,
	}
	c.JSON(200, res)
}

// 统计指定字段(多)
func FangyuanxinxiValueMul(c *gin.Context) {
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

	// 示例条件，可根据实际需求修改
	if "fangyuanxinxi" == "someCondition" {
		where += " AND someField = 'someValue' "
	}

	result := make([][]orm.Params, 0)
	yColumnNameMulArr := strings.Split(yColumnNameMul, ",")
	for _, v := range yColumnNameMulArr {
		sql = "SELECT " + xColumnName + ", SUM(" + v + ") AS total FROM fangyuanxinxi " + where + " GROUP BY " + xColumnName + " LIMIT 10"
		var maps []orm.Params
		_, err := o.Raw(sql).Values(&maps)
		if err != nil {
			c.JSON(500, models.ReturnMsg{
				Code: 500,
				Msg:  "查询失败: " + err.Error(),
				Data: nil,
			})
			return
		}
		result = append(result, maps)
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: result,
	}
	c.JSON(200, res)
}

// 总数接口
func FangyuanxinxiCount(c *gin.Context) {
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
	cond := orm.NewCondition()

	// 处理 zujinjiagestart 和 zujinjiageend
	zujinjiagestart := c.Query("zujinjiagestart")
	zujinjiageend := c.Query("zujinjiageend")
	if zujinjiagestart != "" {
		cond = cond.And("zujinjiage__gte", zujinjiagestart)
	}
	if zujinjiageend != "" {
		cond = cond.And("zujinjiage__lte", zujinjiageend)
	}

	// 封装查询条件处理函数
	handleQueryCondition := func(cond *orm.Condition, columnName, param string) *orm.Condition {
		if param != "" {
			if strings.Contains(param, "%") {
				cond = cond.And(columnName+"__contains", strings.Trim(param, "%"))
			} else {
				cond = cond.And(columnName, param)
			}
		}
		return cond
	}

	// 处理其他查询条件
	columns := []string{
		"fangwumingcheng", "fangwutupian", "fangwuhuxing", "fangwudizhi", "mianji",
		"zujinjiage", "yajinfangshi", "fukuanfangshi", "fangwuchaoxiang", "yajin",
		"xiaoqu", "fangwushipin", "loudongdanyuan", "fanghao", "jifen",
		"fangwujiegou", "fangwuz庄tai", "fangchanzhengbianhao", "fangchanzhengzhaopian",
		"fangzhuxingming", "fangzhushenfenzheng", "fangzhudianhua", "fangwuxiangqing",
		"fabushijian", "gonghao", "yuangongxingming", "thumbsupnum", "crazilynum",
		"clicktime", "clicknum", "discussnum", "storeupnum",
	}
	for _, col := range columns {
		param := c.Query(col)
		cond = handleQueryCondition(cond, col, param)
	}

	cnt, err := o.QueryTable("fangyuanxinxi").SetCond(cond).Count()
	if err != nil {
		c.JSON(500, models.ReturnMsg{
			Code: 500,
			Msg:  "查询总数失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	res := models.ReturnMsg{
		Code: 0,
		Msg:  "请求成功！",
		Data: cnt,
	}
	c.JSON(200, res)
}
