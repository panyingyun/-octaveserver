package main

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	Server_Version  = "v1.1.1"
	Server_UpdateAt = "2022-12-28"
)

// inputs 为 输入文件数组
type ConvertReq struct {
	Inputs []*multipart.FileHeader `form:"inputs"  json:"inputs"`
}

type DatFile struct {
	DATName    string `json:"datname"`
	DATContent string `json:"datcontent"`
}

// 返回结果
type ConvertResp struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	DatFiles []DatFile `json:"datfiles"`
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
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		resp.Code = 10000
		resp.Msg = "请求参数错误"
		c.JSON(http.StatusOK, resp)
	}

	// 解析多输入csv文件
	var filename string
	form, _ := c.MultipartForm()
	files := form.File["inputs"]
	for _, file := range files {
		filename = file.Filename
		log.Println("req filename = ", filename)
		err := c.SaveUploadedFile(file, filepath.Join("convert", filename))
		if err != nil {
			resp.Code = 10001
			resp.Msg = "文件保存错误"
			c.JSON(http.StatusOK, resp)
		}
	}

	// 算法运行
	cmdstr := "octave-cli  Main.m"
	ret, err := ExecString("/app/convert", cmdstr)
	log.Println("ret = ", ret)
	if err != nil {
		log.Println("err = ", err)
		resp.Code = 10002
		resp.Msg = "算法运行错误"
		c.JSON(http.StatusOK, resp)
	}

	// 扫描当前目录下的Dat文件并使用数组返回
	var DatFiles []DatFile
	filenames, err := FindDatName("convert")
	if err != nil {
		log.Println("err = ", err)
		resp.Code = 10003
		resp.Msg = "文件读取错误"
		c.JSON(http.StatusOK, resp)
	}
	for _, name := range filenames {
		datFilename := filepath.Join("convert", name)
		dat, err := os.ReadFile(datFilename)
		if err != nil {
			log.Println("err = ", err)
			resp.Code = 10003
			resp.Msg = "文件读取错误"
			c.JSON(http.StatusOK, resp)
		}
		var datfile DatFile
		datfile.DATName = name
		datfile.DATContent = string(dat)
		DatFiles = append(DatFiles, datfile)
	}
	resp.Code = 0
	resp.Msg = "Success"
	resp.DatFiles = DatFiles
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

// ExecString
func ExecString(workDir string, cmd string) (string, error) {
	log.Println("cmd = ", cmd)
	command := exec.Command("sh", "-c", cmd)
	command.Dir = workDir
	bytes, err := command.CombinedOutput()
	return string(bytes), err
}

// Find all DAT files
func FindDatName(dirname string) ([]string, error) {
	var findNames []string
	f, err := os.Open(dirname)
	if err != nil {
		return findNames, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return findNames, err
	}

	for _, name := range names {
		if strings.HasSuffix(name, ".DAT") {
			findNames = append(findNames, name)
			break
		}
	}
	return findNames, nil
}
