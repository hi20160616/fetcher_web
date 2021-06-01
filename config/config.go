package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var RootPath = "./"

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
	Title     string   `json:"title"`
	Domain    string   `json:"domain"`
	URL       []string `json:"url"`
	Addr      string   `json:"addr"`
	Timeout   string   `json:"timeout"`
	Heartbeat string   `json:"heartbeat"`
}

var Data = &configuration{}

func setRootPath() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	RootPath = root
	return nil
}

func get() error {
	f, err := os.ReadFile(filepath.Join(RootPath, "config/config.json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(f, Data)
}

func init() {
	if err := setRootPath(); err != nil {
		log.Printf("config init error: %v", err)
	}
	if err := get(); err != nil {
		log.Printf("config init error: %v", err)
	}
}
