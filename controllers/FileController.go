package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go-schema/models"
)

// Upload 文件上传接口
func FileUpload(c *gin.Context) {
	// 读取文件信息
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ReturnMsg{
			Code: -1,
			Msg:  "获取文件失败: " + err.Error(),
			Data: nil,
		})
		return
	}
	defer file.Close()

	// 生成文件名
	filename := strconv.Itoa(int(time.Now().Unix())) + filepath.Ext(header.Filename)
	if strings.Contains(c.Query("type"), "_template") {
		filename = c.Query("type") + filepath.Ext(header.Filename)
	}

	// 保存文件
	err = c.SaveUploadedFile(header, "views/upload/"+filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ReturnMsg{
			Code: -1,
			Msg:  "保存文件失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 假设 ReturnFile 结构体有必要的字段，按实际情况调整
	// 这里简单返回文件名作为示例，可根据 ReturnFile 实际结构修改
	res := models.ReturnFile{
		Code: 0,
		File: filename,
	}

	c.JSON(200, res)
}

// Download 文件下载接口
func FileDownload(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := "views/upload/" + fileName
	c.FileAttachment(filePath, fileName)
	// 设置响应状态码为 200
	c.Status(http.StatusOK)
}
