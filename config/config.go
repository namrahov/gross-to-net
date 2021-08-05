package config

import "github.com/alexflint/go-arg"

const ReportHandlerPath = "/internal/v1/salary-converter"

type Config struct {
	Port int `arg:"env:PORT"`
}

var EnvVars Config

func LoadConfig() {
	_ = arg.Parse(&EnvVars)
}
