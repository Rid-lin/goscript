package main

import (
	"embed"

	"github.com/sirupsen/logrus"
)

type App struct {
	config *Config
	logger *logrus.Logger
	store  *Store
	fs     embed.FS
}

func NewApp(store *Store, cfg *Config, fs embed.FS) *App {
	return &App{
		store:  store,
		config: cfg,
		logger: logrus.New(),
		fs:     fs,
	}
}

func (s *App) configureLogger() {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		logrus.Errorf("Error parse the level of logs (%v). Installed by default = Info", s.config.LogLevel)
		lvl, _ = logrus.ParseLevel("info")
	}
	logrus.SetLevel(lvl)
}

func (s *App) Start() error {

	s.configureLogger()
	s.logger.Debugf("with config: %v\n", s.config)

	artiles, err := s.GetAllArticles()
	if err != nil {
		return err
	}

	err = s.SaveArticlesToFile(artiles, s.config.PathToWeb)
	if err != nil {
		return err
	}

	return nil

}
