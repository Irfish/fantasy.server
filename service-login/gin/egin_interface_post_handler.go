package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinPostHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context), 0)
	{
		handler := NewLoginByAccount()
		ret["/post/login"] = handler.handle
	}
	{
		handler := NewUserRegister()
		ret["/post/user_register"] = handler.handle
	}
	return ret
}
