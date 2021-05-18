package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huchengbei/for-my-girl/backend/models"
	"github.com/huchengbei/for-my-girl/backend/pkg/e"
	"github.com/unknwon/com"
	"net/http"
)

func GetMoments(c *gin.Context) {
	data := models.GetMoments()

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,

	})
}

func AddMoment(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	datetimeStr := c.PostForm("datetime")
	link := c.PostForm("link")
	momentType := c.PostForm("moment_type")
	imageLink := c.PostForm("image_link")
	albumLink := c.PostForm("album_link")
	musicLink := c.PostForm("music_link")
	videoLink := c.PostForm("video_link")

	flag := models.AddMoment(title, content, datetimeStr, link, momentType, imageLink, albumLink, musicLink, videoLink)

	var code int
	if flag {
		code = e.SUCCESS
	} else {
		code = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditMoment(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := c.PostForm("title")
	content := c.PostForm("content")
	datetimeStr := c.PostForm("datetime")
	link := c.PostForm("link")
	momentType := c.PostForm("moment_type")
	imageLink := c.PostForm("image_link")
	albumLink := c.PostForm("album_link")
	musicLink := c.PostForm("music_link")
	videoLink := c.PostForm("video_link")

	flag := models.EditMoment(id, title, content, datetimeStr, link, momentType, imageLink, albumLink, musicLink, videoLink)

	var code int
	if flag {
		code = e.SUCCESS
	} else {
		code = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteMoment(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	flag := models.DeleteMoment(id)

	var code int
	if flag {
		code = e.SUCCESS
	} else {
		code = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}