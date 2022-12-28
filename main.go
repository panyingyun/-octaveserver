package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	Server_Version  = "v1.0.2"
	Server_UpdateAt = "2022-12-28"
)

// ExecString
func ExecString(workDir string, cmd string) (string, error) {
	log.Println("cmd = ", cmd)
	command := exec.Command("sh", "-c", cmd)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

// Type =1 静力分析  = 3 确定性疲劳分析  = 4 防腐面积计算
// matrix 为 输入矩阵文件
// Effective 为输入csv控制参数文件
type ConvertReq struct {
	Type      int64                 `form:"type", json:"type"`
	Matrix    *multipart.FileHeader `form:"matrix"  json:"matrix"`
	Effective *multipart.FileHeader `form:"effective"  json:"effective"`
}

// 返回结果
type ConvertResp struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	DATName    string `json:"datname"`
	DATContent string `json:"datcontent"`
}

// 定义版本号
type Version struct {
	Version  string `json:"version"`
	UpdateAt string `json:"updateat"`
}

// 处理算法调用
func convertHandler(c *gin.Context) {
	var req ConvertReq
	var resp ConvertResp
	// 参数解析
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 10000
		resp.Msg = "请求参数错误"
		c.JSON(http.StatusOK, resp)
	}
	err := c.SaveUploadedFile(req.Matrix, filepath.Join("convert", req.Matrix.Filename))
	if err != nil {
		resp.Code = 10001
		resp.Msg = "文件保存错误"
		c.JSON(http.StatusOK, resp)
	}
	err = c.SaveUploadedFile(req.Effective, filepath.Join("convert", req.Effective.Filename))
	if err != nil {
		resp.Code = 10001
		resp.Msg = "文件保存错误"
		c.JSON(http.StatusOK, resp)
	}
	log.Println("req type = ", req.Type)
	log.Println("req matrix = ", req.Matrix.Filename)
	log.Println("req Effective = ", req.Effective.Filename)
	if req.Type != 1 && req.Type != 3 && req.Type != 4 {
		resp.Code = 10000
		resp.Msg = "请求参数错误, Type 必须为1或者3或4"
		c.JSON(http.StatusOK, resp)
	}
	// 算法运行
	cmdstr := fmt.Sprintf("octave-cli  Main.m  %v", req.Type)
	ret, err := ExecString("/app/convert", cmdstr)
	log.Println("ret = ", ret)
	log.Println("err = ", err)
	// Dat文件读取
	datFilename := filepath.Join("convert", "Static PSI", "JCNINP.DAT")
	dat, err := os.ReadFile(datFilename)
	log.Println("err = ", err)
	resp.DATName = "JCNINP.DAT"
	resp.DATContent = string(dat)
	resp.Code = 0
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
	r.POST("/convert", convertHandler)
	r.GET("/version", versionHandler)
	r.Run(":8630")
}
