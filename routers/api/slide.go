package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huchengbei/for-my-girl/models"
	"github.com/huchengbei/for-my-girl/pkg/e"
	"github.com/unknwon/com"
)

func GetSlides(c *gin.Context) {
	data := models.GetSlides()

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func AddSlide(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	imageLink := c.PostForm("image_link")
	isShowStr := c.PostForm("is_show")
	sequence := com.StrTo(c.PostForm("sequence")).MustInt()

	var isShow bool
	if isShowStr == "true" {
		isShow = true
	} else {
		isShow = false
	}

	flag := models.AddSlide(title, description, imageLink, isShow, sequence)

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

func EditSlide(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := c.PostForm("title")
	description := c.PostForm("description")
	imageLink := c.PostForm("image_link")
	isShowStr := c.PostForm("is_show")
	sequence := com.StrTo(c.PostForm("sequence")).MustInt()

	var isShow bool
	if isShowStr == "true" {
		isShow = true
	} else {
		isShow = false
	}
	flag := models.EditSlide(id, title, description, imageLink, isShow, sequence)

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

func DeleteSlide(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	flag := models.DeleteSlide(id)

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
