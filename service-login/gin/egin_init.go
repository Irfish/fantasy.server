package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Gin) GinInit(egin *gin.Engine) {
	egin.Use(
		func(context *gin.Context) {
			var e error
			defer func() {
				if e != nil {
					fmt.Println(e)
					context.JSON(200, gin.H{
						"debug": e.Error(),
					})
					context.Abort()
				}
			}()
			context.Header("Access-Control-Allow-Origin", "*")
			if context.Request.Method == "OPTIONS" {
				//context.Header("Access-Control-Allow-Headers", "content-type,userId,token,tokenExpiredTime")
				context.Status(200)
				return
			}
		},
	)
}
