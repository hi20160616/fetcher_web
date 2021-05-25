package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type configuration struct {
	Title     string    `json:"title"`
	WebServer webserver `json:"webserver"`
	API       struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	MS map[string]MicroService `json:"ms"`
}

type webserver struct {
	Addr, Tmpl string
}

type api struct {
	Addr, Timeout string
}

type MicroService struct {
	Title   string `json:"title"`
	Domain  string `json:"domain"`
	URL     string `json:"url"`
	Addr    string `json:"addr"`
	Timeout string `json:"timeout"`
}

var Data = &configuration{}

func get() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	// root = "../" // for config test
	cfgFile := filepath.Join(root, "config/config.toml")
	cfg, err := toml.LoadFile(cfgFile)
	if err != nil {
		return err
	}
	return cfg.Unmarshal(Data)
}

func init() {
	if err := get(); err != nil {
		log.Printf("config init error: %#v", err)
	}
}
