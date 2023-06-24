package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	mackerelspeedtest "mgithub.com/Arthur1/mackerel-speedtest"
)

type config struct {
	MackerelAPIKey string `env:"MACKEREL_API_KEY"`
	MackerelServiceName string `env:"MACKEREL_SERVICE_NAME"`
	SpeedtestCommandPath string `env:"SPEEDTEST_COMMAND_PATH"`
	SpeedtestServerID int64 `env:"SPEEDTEST_SERVER_ID"`
}

func main() {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	executor := mackerelspeedtest.NewSpeedTestExecutor(cfg.SpeedtestCommandPath, cfg.SpeedtestServerID)
	result, err := executor.Execute()
	if err != nil {
		log.Fatal(err)
	}
	exporter := mackerelspeedtest.NewMackerelExporter(cfg.MackerelAPIKey, cfg.MackerelServiceName)
	if err := exporter.Export(result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
