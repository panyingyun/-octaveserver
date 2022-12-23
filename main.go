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
	Result float32 `json:"result"`
}

// 定义版本号
type Version struct {
	Version  string `json:"version"`
	UpdateAt string `json:"updateat"`
}

// 处理算法调用
func octaveHandler(c *gin.Context) {

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
