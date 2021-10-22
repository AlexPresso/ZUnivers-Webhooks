package utils

import (
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func LoadConfig() {
	newConfig := false
	dir, err := os.Getwd()
	workingPath := dir + "/config.json"
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(workingPath)
	err = viper.ReadInConfig()
	if err != nil {
		if _, notFound := err.(*fs.PathError); notFound {
			Log("Missing config file, creating an empty one for you.")
			viper.SetConfigFile(basepath + "/../config.default.json")
			err = viper.ReadInConfig()
			err = viper.WriteConfigAs(workingPath)

			newConfig = true

			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	if newConfig {
		Log("Please edit your config.json file before you run the program again (check the repository example).")
		os.Exit(0)
	}

	Log("Configuration loaded.")
}
