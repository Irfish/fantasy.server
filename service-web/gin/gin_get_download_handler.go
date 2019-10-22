package gin

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Download struct {
}

func NewDownload() Download {
	return Download{}
}

func (p *Download) handle(c *gin.Context) {
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
