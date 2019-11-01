package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinGetHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context))
	{
		handler := NewVerificationCode()
		ret["/verificationCode"] = handler.handle
	}
	{
		handler := NewHotUpdateUI()
		ret["hotUpdate/ui/:platform"] = handler.handle
	}
	return ret
}
