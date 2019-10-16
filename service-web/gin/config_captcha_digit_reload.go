package gin

import (
	"sync"

	"github.com/Irfish/component/util"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/mojocn/base64Captcha"
)

var configCaptchaDigitMutex *sync.RWMutex
var configCaptchaDigit *base64Captcha.ConfigDigit

func init() {
	configCaptchaDigitMutex = new(sync.RWMutex)
}

func getConfigCaptchaDigit() base64Captcha.ConfigDigit {
	configCaptchaDigitMutex.RLock()
	defer func() { configCaptchaDigitMutex.RUnlock() }()
	cfg1 := &pb.CaptchaDigitCfg{}
	util.LoadConfigFile("captcha_digit.cfg", cfg1)
	configCaptchaDigit = &base64Captcha.ConfigDigit{
		Height:     int(cfg1.GetHeight()),
		Width:      int(cfg1.GetWidth()),
		CaptchaLen: int(cfg1.GetCaptchaLen()),
		MaxSkew:    float64(cfg1.GetMaxSkew()) / 100,
		DotCount:   int(cfg1.GetDotCount()),
	}
	return *configCaptchaDigit
}
