package gin

import (
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
)

type Download struct {
}

func NewDownload() Download {
	return Download{}
}

func (p *Download) handle(c *gin.Context) {
	platform := c.Param("platform")
	file := c.Param("filePath")
	fileDir := "files/" + platform + "/" + file
	fmt.Println("fileDir:", fileDir)
	if exist := checkFileIsExist(fileDir); !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not exist:" + fileDir, "status": "StatusNotFound"})
		return
	}
	http.ServeFile(c.Writer, c.Request, fileDir)
	return
}

func (p *Download) handle1(c *gin.Context) {
	response, err := http.Get("http://192.168.0.130:4000/files/egin.cfg")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.cfg"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
