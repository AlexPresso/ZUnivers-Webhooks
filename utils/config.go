package utils

import (
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"os"
)

func LoadConfig() {
	newConfig := false
	dir, err := os.Getwd()
	path := dir + "/config.json"
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		if _, notFound := err.(*fs.PathError); notFound {
			Log("Missing config file, creating an empty one for you.")
			viper.SetConfigFile("config.default.json")
			err = viper.ReadInConfig()
			err = viper.WriteConfigAs(path)

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
