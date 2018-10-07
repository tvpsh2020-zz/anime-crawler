package main

import (
	"github.com/tvpsh2020/anime-crawler/config"
	"github.com/tvpsh2020/anime-crawler/modules"
)

func init() {
	config.Initialize()
}
func main() {
	go modules.FetchDmhy()
	router := modules.InitRouter()
	router.Run(":" + config.Server.Setting.Port)
}
