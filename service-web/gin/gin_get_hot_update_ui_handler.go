package gin

import (
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
)

type HotUpdateUI struct {
}

func NewHotUpdateUI() HotUpdateUI {
	return HotUpdateUI{}
}

func (p *HotUpdateUI) handle(c *gin.Context) {
	platform := c.Param("platform")
	file := c.Query("f")
	v := c.Query("v")
	fileDir := "files/" + platform + "/" + file
	fmt.Println("fileDir:", fileDir, v)
	if exist := checkFileIsExist(fileDir); !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not exist:" + fileDir, "status": "StatusNotFound"})
		return
	}
	http.ServeFile(c.Writer, c.Request, fileDir)
	return
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//
//func (p *Download) handle1(c *gin.Context) {
//	response, err := http.Get("http://192.168.0.130:4000/files/egin.cfg")
//	if err != nil || response.StatusCode != http.StatusOK {
//		c.Status(http.StatusServiceUnavailable)
//		return
//	}
//	reader := response.Body
//	contentLength := response.ContentLength
//	contentType := response.Header.Get("Content-Type")
//
//	extraHeaders := map[string]string{
//		"Content-Disposition": `attachment; filename="gopher.cfg"`,
//	}
//
//	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
//}
