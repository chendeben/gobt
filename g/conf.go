package g

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/btlike/repository"
	"github.com/xgfone/gobt/logger"
)

var (
	Repository repository.Repository
	Conf       Config
)

type Config struct {
	Database string `json:"db"`
	LogFile  string `json:"logfile"`
	LogLevel string `json:"loglevel"`
}

func initConfig(filename string) {
	if f, err := os.Open(filename); err != nil {
		panic(err)
	} else if data, err := ioutil.ReadAll(f); err != nil {
		panic(err)
	} else if err = json.Unmarshal(data, &Conf); err != nil {
		panic(err)
	}
}

func Init(config_file string) {
	initConfig(config_file)

	if repo, err := repository.NewMysqlRepository(Conf.Database, 256, 256); err != nil {
		panic(err)
	} else {
		Repository = repo
	}

	if _logger, err := logger.NewLogger(Conf.LogLevel, Conf.LogFile); err != nil {
		panic(err)
	} else {
		logger.Logger = _logger
	}
}
