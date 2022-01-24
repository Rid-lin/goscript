package main

import (
	"embed"
	"os"

	"github.com/sirupsen/logrus"
)

//go:embed assets
var assets embed.FS

func main() {
	cfg := NewConfig()
	db, err := NewDB(cfg.DSN)
	if err != nil {
		logrus.Error("sql created error:", err)
		os.Exit(1)
	}
	store := NewStore(db)

	app := NewApp(store, cfg, assets)

	err = app.Start()
	if err != nil {
		logrus.Error("server returned error:", err)
	}

}
