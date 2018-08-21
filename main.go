package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vimsucks/resmtp/config"
	"github.com/vimsucks/resmtp/model"
	"github.com/vimsucks/resmtp/route"
)

func main() {
	cliConf := config.ParseCli()
	conf := config.InitConfig(cliConf)
	model.InitDB(conf)
	defer model.DB.Close()

	gin.SetMode(conf.Mode)

	r := gin.Default()

	route.RegisterRoute(r)

	r.Run(fmt.Sprintf("%s:%d", conf.Listen, conf.Port))
}
