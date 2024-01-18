package main

import (
	"cron-request/internal/config"
	executor2 "cron-request/internal/executor"
	"github.com/google/logger"
	"github.com/spf13/viper"
	"io"
	"time"
)

func readConfig() (config.Config, error) {
	viper.SetConfigName("Config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return config.Config{}, err
	}
	var requestConfigs []config.RequestConfig
	err = viper.UnmarshalKey("requests", &requestConfigs)

	readConfig := config.Config{
		MetaData: config.MetaData{
			Name: viper.GetString("metadata.name"),
		},
		ExecutionConfig: config.ExecutionConfig{
			Interval: viper.GetInt("execution.interval"),
		},
		Requests: requestConfigs,
	}

	return readConfig, nil
}

func executeRequest(requestConfig config.RequestConfig) {
	startTime := time.Now()
	logger.Info("Executing request: ", requestConfig.Name)
	executor := executor2.New(requestConfig)
	body, err := executor.Execute()
	if body != "" {
		logger.Info("Response body: ", body)
	}
	if err != nil {
		logger.Error("Error executing request: ", err)
	}
	endTime := time.Now()
	logger.Info("Execution time: ", endTime.Sub(startTime))
}

func main() {
	logger.Init("Logger", true, false, io.Discard)
	readConfig, err := readConfig()
	if err != nil {
		logger.Error("Error reading config: ", err)
		return
	}

	ticker := time.NewTicker(time.Duration(readConfig.ExecutionConfig.Interval) * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				for _, requestConfig := range readConfig.Requests {
					go executeRequest(requestConfig)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	<-quit
}
