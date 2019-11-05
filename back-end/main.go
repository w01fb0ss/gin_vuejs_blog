package main

import (
	"fmt"
	"os"

	"github.com/w01fb0ss/gin_vuejs_blog/back-end/config"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/models"
	"github.com/w01fb0ss/gin_vuejs_blog/back-end/routers"
)

func init()  {
	if err := config.InitConfig(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := models.InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	router := routers.InitRouter()
	router.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))
}
