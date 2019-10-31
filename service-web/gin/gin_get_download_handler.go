package gin

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Download struct {
}

func NewDownload() Download {
	return Download{}
}

func (p *Download) handle(c *gin.Context) {
	s := strings.Split(c.Request.URL.Path, "/")
	if len(s) > 2 {
		//file:="files/files.txt /"
		file := "files/" + s[len(s)-2] + "/" + s[len(s)-1]
		if exist := checkFileIsExist(file); !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "file not exist ","status":"StatusNotFound"})
			return
		}
		http.ServeFile(c.Writer, c.Request, file)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "file path error","status":"StatusNotFound"})
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
