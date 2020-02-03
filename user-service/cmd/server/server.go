package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/mesment/microservice/user-service/pkg/cmd"
	"github.com/mesment/microservice/user-service/pkg/logger"
	"github.com/mesment/microservice/user-service/pkg/util"
)

var (
	cfg            *util.Config
	configPath                     = flag.String("f", "../../config/config.toml", "config file")
	version                        = flag.Bool("v", false, "version")
)

func init()  {
	flag.Parse()
	cfg = util.InitConfig(*configPath)
}

func main()  {
	if *version {
		fmt.Println(util.GetVersion())
		return
	}
	fmt.Println("cfg",cfg)
	logger.Init(cfg.Log.LogLevel, cfg.Log.LogFile)
	log := logger.NewSugardLogger()
	if err := cmd.RunServer(cfg); err != nil {
		log.Errorf( "%s", err.Error())
		os.Exit(1)
	}
}
