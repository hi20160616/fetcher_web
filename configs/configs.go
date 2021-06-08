package configs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var RootPath = ""

type configuration struct {
	Gist      string    `json:"gist"`
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

func load() error {
	f, err := os.ReadFile(filepath.Join(RootPath, "configs/configs.json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(f, Data); err != nil {
		return err
	}

	resp, err := http.Get(Data.Gist)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &Data); err != nil {
		return err
	}
	data, err := json.MarshalIndent(Data, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(
		filepath.Join(RootPath, "configs/configs.json"),
		data, 0755); err != nil {
		return err
	}
	return nil
}

func init() {
	if err := setRootPath(); err != nil {
		log.Printf("config init error: %v", err)
	}
	if err := load(); err != nil {
		log.Printf("config get() error: %v", err)
	}
}

// Reset is for test to reset RootPath and invoke get()
func Reset(pwd string) error {
	RootPath = pwd
	return load()
}
