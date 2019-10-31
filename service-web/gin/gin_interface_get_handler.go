package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinGetHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context))
	{
		handler := NewVerificationCode()
		ret["/verificationCode"] = handler.handle
	}
	{
		handler := NewDownload()
		ret["ui/:platform/:filePath"] = handler.handle //下载资源路由 如何多层匹配 ? ? ?
	}
	return ret
}
