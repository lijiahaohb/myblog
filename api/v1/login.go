package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBind(&data)
	var token string 
	var code int

	code = model.CheckLogin(data.UserName, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.UserName)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrorMsg(code),
		"token": token,
	})
}
