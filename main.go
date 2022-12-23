package main

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Server_Version  = "v1.0.1"
	Server_UpdateAt = "2022-12-23"
)

// calc =1 矩阵求和  = 2 求矩阵平方和
// matrix 为 输入矩阵文件
type OctaveReq struct {
	Calc   int64                 `form:"calc", json:"calc"`
	Matrix *multipart.FileHeader `form:"matrix"  json:"matrix"`
}

// 返回结果
type OctaveResp struct {
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Result float32 `json:"result"`
}

// 定义版本号
type Version struct {
	Version  string `json:"version"`
	UpdateAt string `json:"updateat"`
}

// 处理算法调用
func octaveHandler(c *gin.Context) {
	var req OctaveReq
	var resp OctaveResp
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 10000
		resp.Msg = "请求参数错误"
		c.JSON(http.StatusOK, resp)
	}
	err := c.SaveUploadedFile(req.Matrix, req.Matrix.Filename)
	if err != nil {
		resp.Code = 10001
		resp.Msg = "文件保存错误"
		c.JSON(http.StatusOK, resp)
	}
	resp.Code = 0
	resp.Result = 12.8 //TODO
	c.JSON(http.StatusOK, resp)
}

// 版本号定义
func versionHandler(c *gin.Context) {
	var v Version
	v.Version = Server_Version
	v.UpdateAt = Server_UpdateAt
	c.JSON(http.StatusOK, v)
}

func main() {
	r := gin.Default()
	r.POST("/octave", octaveHandler)
	r.GET("/version", versionHandler)
	r.Run(":8630")
}
