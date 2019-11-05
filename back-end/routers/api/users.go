/*
 * Copyright (c) 2019.
 * Author:Alex	xiaoshenlong1910@gmail.com

   .--,       .--,
 ( (  \.---./  ) )
  '.__/o   o\__.'
     {=  ^  =}
      >  -  <
     /       \
    //       \\
   //|   .   |\\
   "'\       /'"_.-~^`'-.
      \  _  /--'         `
    ___)( )(___
   (((__) (__)))    高山仰止,景行行止.虽不能至,心向往之。

*/

package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/models"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/pkg/e"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.GetUser(username, password)
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
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type register struct {
	Username string `valid:"Required; MaxSize(50)"`
	Email string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetRegister(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	valid := validation.Validation{}
	r := register{
		Username: username,
		Email:    email,
		Password: password,
	}
	ok, _ := valid.Valid(&r)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.IsExist(username, email, password)
		if isExist {
			code = e.ERROR_REGISTER
		} else {
			// TODO: 将信息写入数据库
			data["username"] = username
			data["email"] = email
			code = e.SUCCESS
		}

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}