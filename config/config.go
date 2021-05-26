package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type configuration struct {
	Title     string    `json:"title"`
	WebServer webserver `json:"webserver"`
	API       struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	MS map[string]MicroService `json:"microservice"`
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
	f, err := os.ReadFile(filepath.Join(root, "config/config.json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(f, Data)
}

func init() {
	if err := get(); err != nil {
		log.Printf("config init error: %#v", err)
	}
}
