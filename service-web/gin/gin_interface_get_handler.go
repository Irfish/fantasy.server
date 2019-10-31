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
		ret["/test/win/files.txt"] = handler.handle
		ret["/test/win/lobby.unity3d"] = handler.handle
		ret["/test/win/lobby.unity3d.manifest"] = handler.handle
		ret["/test/win/lobby.unity3d.manifest.meta"] = handler.handle
		ret["/test/win/lobby.unity3d.meta"] = handler.handle
	}
	return ret
}
