package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type configuration struct {
	Address, RootPath, TmplPath string
	Sites                       []Site
}
type Site struct {
	Title, URL, Port string
}

var Value = &configuration{}

func init() {
	if err := initJson(); err != nil {
		log.Println(err)
	}
}

func initJson() error {
	return initExt(".json")
}

func initToml() error {
	return initExt(".toml")
}

func initExt(ext string) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.ReadFile(filepath.Join(root, "config/config"+ext))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, Value); err != nil {
		return err
	}
	Value.RootPath = root
	Value.TmplPath = filepath.Join(root, Value.TmplPath)
	return nil
}
