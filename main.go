package main

import (
	"log"
	"mime/multipart"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

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
	Result float64 `json:"result"`
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
	log.Println("req calc = ", req.Calc)
	log.Println("req matrix = ", req.Matrix.Filename)
	if req.Calc < 1 || req.Calc > 2 {
		resp.Code = 10000
		resp.Msg = "请求参数错误, calc 必须为1或者2"
		c.JSON(http.StatusOK, resp)
	}
	if req.Calc == 1 {
		ret, err := ExecString("/app/appsum", "octave-cli  main.m")
		log.Println("err = ", err)
		resp.Result = ParserResult(ret)
	} else {
		ret, err := ExecString("/app/appsquare", "octave-cli  main.m")
		log.Println("err = ", err)
		resp.Result = ParserResult(ret)
	}
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
	//ParserResult("sum = 642")
	r := gin.Default()
	r.POST("/octave", octaveHandler)
	r.GET("/version", versionHandler)
	r.Run(":8630")
}

func ParserResult(ret string) float64 {
	ret = strings.TrimSuffix(ret, "\n")
	rets := strings.Split(ret, " = ")
	log.Println("rets = ", rets)
	if len(rets) < 2 {
		return 0.0
	}
	log.Println("rets[1] = ", rets[1])
	num, err := strconv.ParseFloat(rets[1], 64)
	log.Println("err = ", err)
	log.Println("num = ", num)
	return num
}

// ExecString
func ExecString(workDir string, cmd string) (string, error) {
	log.Println("cmd = ", cmd)
	command := exec.Command("sh", "-c", cmd)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}
