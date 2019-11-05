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

package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/middleware/jwt"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", api.Ping)
	}
	return r
}
