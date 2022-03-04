package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cristalhq/aconfig"
)

type RevoltSettings struct {
	Token string `default:""`
	Uid string `default:""`
}


type DiruConfig struct {
	Revolt RevoltSettings
	DeeplToken string `default:""`
}

func GetConfig(fileName string) DiruConfig {

	if _, err := os.Stat("config/" + fileName + ".json"); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile("config/" + fileName + ".json", []byte("{\n    \"revolt\": {\n        \"token\": \"\",\n        \"uid\": \"\"\n    },\n    \"deepl_token\": \"\"\n}"), 0644)

		fmt.Printf("config/%s.json not found. It has been created for you, you must enter your values for revolt_token and deepl_token.", fileName)
		os.Exit(1)
	}

	var cfg DiruConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipDefaults: true,
		SkipEnv: true,
		SkipFlags: true,
		Files: []string{"config/" + fileName + ".json"},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	config := DiruConfig{
		Revolt: cfg.Revolt,
		DeeplToken:   cfg.DeeplToken,
	}

	return config
}