package cfg

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/cristalhq/aconfig"
)


type DiruConfig struct {
	RevoltToken string `default:""`
	DeeplToken   string `default:""`
}

func GetConfig() DiruConfig {

	if _, err := os.Stat("config/diru.json"); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile("config/diru.json", []byte("{\n    \"revolt_token\": \"\",\n    \"deepl_token\": \"\"\n}"), 0644)

		panic("config/diru.json not found. It has been created for you, you must enter your values for revolt_token and deepl_token.")
	}

	var cfg DiruConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipDefaults: true,
		SkipEnv: true,
		SkipFlags: true,
		Files: []string{"config/Diru.json"},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	config := DiruConfig{
		RevoltToken: cfg.RevoltToken,
		DeeplToken:   cfg.DeeplToken,
	}

	return config
}