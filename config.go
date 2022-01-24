package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	DSN        string `default:"" usage:"Database server name"`
	ConfigPath string `default:"/etc/goscript" usage:"folder path to all config files"`
	PathToWeb  string `default:"/local/files" usage:"the path to the generated page"`
	LogLevel   string `default:"info" usage:"Log level: panic, fatal, error, warn, info, debug, trace"`
}

func NewConfig() *Config {
	// fix for https://github.com/cristalhq/aconfig/issues/82
	args := []string{}
	for _, a := range os.Args {
		if !strings.HasPrefix(a, "-test.") {
			args = append(args, a)
		}
	}
	// fix for https://github.com/cristalhq/aconfig/issues/82

	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		// feel free to skip some steps :)
		// MergeFiles: true,
		SkipEnv:            false,
		SkipFiles:          false,
		AllowUnknownFlags:  true,
		AllowUnknownFields: true,
		SkipDefaults:       false,
		SkipFlags:          false,
		FailOnFileNotFound: false,
		EnvPrefix:          "GOSCRIPT",
		FlagPrefix:         "",
		FileFlag:           "config",
		Files: []string{
			"./config.yaml",
			"./config/config.yaml",
			"/etc/goscript/config.yaml",
			"/etc/goscript/config/config.yaml",
			"/usr/local/goscript/config.yaml",
			"/usr/local/goscript/config/config.yaml",
			"/opt/goscript/config.yaml",
			"/opt/goscript/config/config.yaml",
		},
		FileDecoders: map[string]aconfig.FileDecoder{
			// from `aconfigyaml` submodule
			// see submodules in repo for more formats
			".yaml": aconfigyaml.New(),
		},
		Args: args[1:], // [1:] важно, см. доку к FlagSet.Parse
	})
	if err := loader.Load(); err != nil {
		fmt.Println(err)
	}

	return &cfg
}
