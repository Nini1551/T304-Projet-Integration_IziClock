package alarms

import (
	"server/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	login := route.Group("/login")
	{
		login.GET("", login_oauth)
	}
}

func login_oauth(c *gin.Context) {
	state := config.RandToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + config.getLoginURL() + "'><button>Login with Google!</button> </a> </body></html>"))
}
