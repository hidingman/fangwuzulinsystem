package controllers

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"

	"bytes"
	"go-schema/models"
	"go-schema/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// 获取某表的某个字段列表接口
func CommonControllerOption(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	columnName := c.Param("columnName")
	conditionColumn := c.Query("conditionColumn")
	conditionValue := c.Query("conditionValue")

	where := " where 1 = 1 "
	if conditionColumn != "" && conditionValue != "" {
		where = fmt.Sprintf("%s AND %s = '%s'", where, conditionColumn, conditionValue)
	}

	sql := fmt.Sprintf("select %s from %s %s", columnName, tableName, where)

	var list orm.ParamsList
	o.Raw(sql).ValuesFlat(&list)
	res := models.ReturnMsg{
		0, "请求成功！", list,
	}

	c.JSON(http.StatusOK, res)
}

// 获取某表的单行记录接口
func CommonControllerFollow(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	columnName := c.Param("columnName")
	columnValue := c.Query("columnValue")

	where := " where 1 = 1 "
	if columnName != "" && columnValue != "" {
		where = fmt.Sprintf("%s AND %s = '%s'", where, columnName, columnValue)
	}

	sql := fmt.Sprintf("select * from %s %s", tableName, where)

	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	if len(maps) == 0 {
		c.JSON(http.StatusNotFound, models.ReturnMsg{
			404, "未找到记录", nil,
		})
		return
	}
	res := models.ReturnMsg{
		0, "请求成功！", maps[0],
	}

	c.JSON(http.StatusOK, res)
}

// 修改某表的sfsh状态接口
func CommonControllerSh(c *gin.Context) {
	o := orm.NewOrm()

	var mp map[string]interface{}
	if err := c.ShouldBindJSON(&mp); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "请求参数解析失败", err,
		})
		return
	}

	tableName := c.Param("tableName")
	sfsh, ok := mp["sfsh"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "sfsh 参数类型错误", nil,
		})
		return
	}
	id, ok := mp["id"].(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "id 参数类型错误", nil,
		})
		return
	}

	if sfsh == "是" {
		sfsh = "否"
	} else {
		sfsh = "是"
	}

	sql := fmt.Sprintf(`UPDATE %s SET sfsh = '%s' WHERE id = ?`, tableName, sfsh)

	if _, err := o.Raw(sql, id).Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "编辑失败", err,
		})
		return
	}
	res := models.ReturnMsg{
		0, "编辑成功！", nil,
	}

	c.JSON(http.StatusOK, res)
}

// 获取需要提醒的记录数接口
func CommonControllerRemind(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	columnName := c.Param("columnName")
	typeValueStr := c.Param("type")
	typeValue, err := strconv.Atoi(typeValueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "type 参数类型错误", err,
		})
		return
	}
	remindstart := c.Query("remindstart")
	remindend := c.Query("remindend")

	sql := "SELECT 0 AS count"

	if typeValue == 1 {
		if remindstart != "" {
			sql = fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s WHERE %s >= %s`, tableName, columnName, remindstart)
		}
		if remindend != "" {
			sql = fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s WHERE %s <= %s`, tableName, columnName, remindend)
		}
	}

	if typeValue == 2 {
		if remindstart != "" {
			days, _ := strconv.Atoi(remindstart)
			remindStart := utils.GetDateTimeFormat(0-days, "yyyy-MM-dd")
			sql = fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s WHERE %s >= '%s'`, tableName, columnName, remindStart)
		}
		if remindend != "" {
			days, _ := strconv.Atoi(remindend)
			remindEnd := utils.GetDateTimeFormat(days, "yyyy-MM-dd")
			sql = fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s WHERE %s <= '%s'`, tableName, columnName, remindEnd)
		}
	}

	type Result struct {
		Count int
	}
	var result Result
	if err := o.Raw(sql).QueryRow(&result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "查询失败", err,
		})
		return
	}

	res := models.ReturnCount{
		0, result.Count,
	}

	c.JSON(http.StatusOK, res)
}

// 计算规则接口
func CommonControllerCal(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	columnName := c.Param("columnName")

	sql := fmt.Sprintf(`SELECT SUM(%s) AS sum, MAX(%s) AS max, MIN(%s) AS min, AVG(%s) AS avg FROM %s`, columnName, columnName, columnName, columnName, tableName)

	type Result struct {
		Sum float64 `json:"sum"`
		Max float64 `json:"max"`
		Min float64 `json:"min"`
		Avg float64 `json:"avg"`
	}
	var result Result
	if err := o.Raw(sql).QueryRow(&result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "查询失败", err,
		})
		return
	}

	res := models.ReturnMsg{
		0, "请求成功！", result,
	}

	c.JSON(http.StatusOK, res)
}

// 类别统计接口
func CommonControllerGroup(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	columnName := c.Param("columnName")

	sql := fmt.Sprintf(`SELECT COUNT(*) AS total, %s FROM %s GROUP BY %s`, columnName, tableName, columnName)

	var result []orm.Params
	if _, err := o.Raw(sql).Values(&result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "查询失败", err,
		})
		return
	}

	res := models.ReturnMsg{
		0, "请求成功！", result,
	}

	c.JSON(http.StatusOK, res)
}

// 按值统计接口
func CommonControllerValue(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	xColumnName := c.Param("xColumnName")
	yColumnName := c.Param("yColumnName")

	sql := fmt.Sprintf(`SELECT %s, SUM(%s) AS total FROM %s GROUP BY %s`, xColumnName, yColumnName, tableName, xColumnName)

	var result []orm.Params
	if _, err := o.Raw(sql).Values(&result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "查询失败", err,
		})
		return
	}

	res := models.ReturnMsg{
		0, "请求成功！", result,
	}

	c.JSON(http.StatusOK, res)
}

// 时间统计接口
func CommonControllerValueTime(c *gin.Context) {
	o := orm.NewOrm()

	tableName := c.Param("tableName")
	xColumnName := c.Param("xColumnName")
	yColumnName := c.Param("yColumnName")
	timeStatType := c.Param("timeStatType")

	var sql string

	if timeStatType == "日" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d') " + xColumnName + ", sum(" + yColumnName + ") total FROM " + tableName + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m-%d')"
	}
	if timeStatType == "月" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y-%m') " + xColumnName + ", sum(" + yColumnName + ") total FROM " + tableName + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y-%m')"
	}
	if timeStatType == "年" {
		sql = "SELECT DATE_FORMAT(" + xColumnName + ", '%Y') " + xColumnName + ", sum(" + yColumnName + ") total FROM " + tableName + " GROUP BY DATE_FORMAT(" + xColumnName + ", '%Y')"
	}

	var result []orm.Params
	if _, err := o.Raw(sql).Values(&result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "查询失败", err,
		})
		return
	}

	res := models.ReturnMsg{
		0, "请求成功！", result,
	}

	c.JSON(http.StatusOK, res)
}

// 人脸对比接口
func CommonControllerMatchFace(c *gin.Context) {
	o := orm.NewOrm()
	APIKeyInfo := models.Config{Name: "APIKey"}
	SecretKeyInfo := models.Config{Name: "SecretKey"}
	if err := o.Read(&APIKeyInfo, "Name"); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取 APIKey 失败", err,
		})
		return
	}
	if err := o.Read(&SecretKeyInfo, "Name"); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取 SecretKey 失败", err,
		})
		return
	}
	APIKey := APIKeyInfo.Value
	SecretKey := SecretKeyInfo.Value

	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", APIKey, SecretKey)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取 access_token 失败", err,
		})
		return
	}
	defer resp.Body.Close()
	resultStrBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "读取响应失败", err,
		})
		return
	}
	resultStr := string(resultStrBytes)

	var maps orm.Params
	if err := json.Unmarshal([]byte(resultStr), &maps); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "解析 access_token 失败", err,
		})
		return
	}
	accessToken, ok := maps["access_token"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取 access_token 失败", nil,
		})
		return
	}

	face1 := c.PostForm("face1")
	face2 := c.PostForm("face2")
	imgFile1, err := os.Open("views/upload/" + face1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "打开图片 1 失败", err,
		})
		return
	}
	defer imgFile1.Close()
	imgFile2, err := os.Open("views/upload/" + face2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "打开图片 2 失败", err,
		})
		return
	}
	defer imgFile2.Close()

	fInfo1, err := imgFile1.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取图片 1 信息失败", err,
		})
		return
	}
	fInfo2, err := imgFile2.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取图片 2 信息失败", err,
		})
		return
	}
	var size1 int64 = fInfo1.Size()
	var size2 int64 = fInfo2.Size()
	buf1 := make([]byte, size1)
	buf2 := make([]byte, size2)
	fReader1 := bufio.NewReader(imgFile1)
	fReader2 := bufio.NewReader(imgFile2)
	if _, err := fReader1.Read(buf1); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "读取图片 1 失败", err,
		})
		return
	}
	if _, err := fReader2.Read(buf2); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "读取图片 2 失败", err,
		})
		return
	}
	imgBase64Str1 := base64.StdEncoding.EncodeToString(buf1)
	imgBase64Str2 := base64.StdEncoding.EncodeToString(buf2)

	var m1 = map[string]string{
		"image":            imgBase64Str1,
		"image_type":       "BASE64",
		"face_type":        "LIVE",
		"quality_control":  "LOW",
		"liveness_control": "NONE",
	}
	var m2 = map[string]string{
		"image":            imgBase64Str2,
		"image_type":       "BASE64",
		"face_type":        "LIVE",
		"quality_control":  "LOW",
		"liveness_control": "NONE",
	}

	var reqParams = [2]map[string]string{m1, m2}
	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "请求参数编码失败", err,
		})
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprintf(`https://aip.baidubce.com/rest/2.0/face/v3/match?access_token=%s`, accessToken), bytes.NewBuffer(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "创建请求失败", err,
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "人脸对比请求失败", err,
		})
		return
	}
	var resultMap orm.Params
	if err := json.NewDecoder(resp.Body).Decode(&resultMap); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "人脸对比响应解析失败", err,
		})
		return
	}

	var score interface{}
	result, ok := resultMap["result"].(map[string]interface{})
	if ok {
		score, ok = result["score"]
		if !ok {
			c.JSON(http.StatusInternalServerError, models.ReturnMsg{
				500, "获取对比分数失败", nil,
			})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "获取对比结果失败", nil,
		})
		return
	}

	res := models.ReturnScore{
		0, score,
	}

	c.JSON(http.StatusOK, res)
}

// 爬取
func CommonControllerSpider(c *gin.Context) {
	tableName := c.PostForm("tableName")
	str := "cd /yykj/python/9999/spider${spiderSchemaName} && scrapy crawl " + tableName + "Spider -a databaseName=fangwuzulinsystem"
	if _, err := exec.Command("/bin/bash", "-c", str).Output(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "爬取失败", err,
		})
		return
	}
	res := models.ReturnMsg{
		0, "爬取成功！", nil,
	}

	c.JSON(http.StatusOK, res)
}

// 数据备份
func CommonControllerMysqldump(c *gin.Context) {
	filepath := "views/upload/mysql.dmp"
	str := "/usr/bin/mysqldump -h127.0.0.1 -uroot -P3306 -p123456 fangwuzulinsystem > " + filepath
	if _, err := exec.Command("/bin/bash", "-c", str).Output(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "备份失败", err,
		})
		return
	}
	c.FileAttachment(filepath, "mysql.dmp")
}

func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", "8tOAUqbKuKnr71TPMyFEZekn", "iqRDT7HOJgyZcVQy2MtWODY27trwDlrt")
	resp, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

// 向 AI 提问
func CommonControllerAskAI(c *gin.Context) {
	formData := map[string]string{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			400, "请求参数解析失败", err,
		})
		return
	}
	reqUrl := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions?access_token=" + GetAccessToken()
	// 创建要发送的 JSON 数据
	payload := map[string]interface{}{
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": formData["ask"],
			},
		},
		"temperature":     0.8,
		"top_p":           0.8,
		"penalty_score":   1,
		"disable_search":  false,
		"enable_citation": false,
		"collapsed":       true,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "JSON 数据编码失败", err,
		})
		return
	}
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "创建请求失败", err,
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "请求失败", err,
		})
		return
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "读取响应失败", err,
		})
		return
	}
	type TranslationResult struct {
		Result    string `json:"result"`
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}
	var result TranslationResult
	if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			500, "解析响应失败", err,
		})
		return
	}

	if result.ErrorCode > 0 {
		c.JSON(http.StatusOK, models.ReturnMsg{
			result.ErrorCode, result.ErrorMsg, "",
		})
	} else {
		c.JSON(http.StatusOK, models.ReturnMsg{
			0, "请求成功！", result.Result,
		})
	}
}
