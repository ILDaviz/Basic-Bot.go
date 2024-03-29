package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/* global variables */
var (
	Prefix string
	Token  string
)

type config struct {
	Prefix string
	Token  string
}

func init() {
	config := &config{}
	body, err := ioutil.ReadFile("./data/config.json")

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, config)

	Prefix = config.Prefix
	Token = config.Token
	/* Before starting each variable is assigned with its proper value */
}
