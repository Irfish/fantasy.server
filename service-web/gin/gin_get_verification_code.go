package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/mojocn/base64Captcha"
)

type VerificationCode struct {
}

func NewVerificationCode() VerificationCode {
	return VerificationCode{}
}

func (p *VerificationCode) handle(c *gin.Context) {
	log.Logf("Verification Code")
	id, dig := base64Captcha.GenerateCaptcha("", getConfigCaptchaDigit())
	d := dig.(*base64Captcha.CaptchaImageDigit)
	img := base64Captcha.CaptchaWriteToBase64Encoding(d)
	c.JSON(http.StatusOK, gin.H{"img": img, "id": id, "val": d.VerifyValue})
	//c.Writer.Write([]byte("Verification Code 123456"))
}
