package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinPostHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context))
	handler := NewLoginByAccount()
	ret["/post/login"] = handler.handle
	return ret
}
