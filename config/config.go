package config

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/komkom/toml"
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
	f, err := os.ReadFile(filepath.Join(root, "config/config.toml"))
	if err != nil {
		return err
	}

	dec := json.NewDecoder(toml.New(bytes.NewBuffer(f)))
	err = dec.Decode(Data)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	if err := get(); err != nil {
		log.Printf("config init error: %#v", err)
	}
}
