package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinGetHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context))
	handler := NewVerificationCode()
	ret["/verificationCode"] = handler.handle
	return ret
}
