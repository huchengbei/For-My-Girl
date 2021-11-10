package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huchengbei/for-my-girl/models"
	"github.com/huchengbei/for-my-girl/pkg/e"
	"github.com/huchengbei/for-my-girl/pkg/logging"
	"github.com/huchengbei/for-my-girl/pkg/util"
)

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if username != "" && password != "" {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		logging.Info(code, e.GetMsg(code))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
