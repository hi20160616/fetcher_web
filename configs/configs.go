package configs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	Data = &configuration{ProjectName: "fetchnews"}
)

type configuration struct {
	RootPath    string
	ProjectName string
	Debug       bool      `json:"debug"`
	Verbose     bool      `json:"verbose"`
	Gist        string    `json:"gist"`
	Title       string    `json:"title"`
	WebServer   webserver `json:"webserver"`
	API         struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	MS  map[string]MicroService `json:"microservice"`
	Err error
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

func init() {
	if err := setRootPath(); err != nil {
		log.Printf("config init error: %v", err)
	}
	if err := load(); err != nil {
		log.Printf("config get() error: %v", err)
	}
}

func setRootPath() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	Data.RootPath = root
	if strings.Contains(os.Args[0], ".test") {
		rootPath4Test()
	}
	return nil
}

func load() error {
	f, err := os.ReadFile(filepath.Join(Data.RootPath, "configs/configs.json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(f, Data); err != nil {
		return err
	}

	// test gist
	// Data.Gist = "https://gist.github.com/hi20160616/d932caa9c0c905c07ee4f773fea7c850/raw/configs.json"
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
	// Write back to configs.json to check if only debug is true
	if Data.Debug {
		return os.WriteFile(filepath.Join(Data.RootPath,
			"configs/configs.json"),
			data, 0755)
	}
	return nil
}

func rootPath4Test() string {
	ps := strings.Split(Data.RootPath, string(Data.ProjectName))

	n := 0
	s := string(os.PathSeparator)
	if len(ps) > 1 {
		n = strings.Count(ps[1], s)
	}
	for i := 0; i < n; i++ {
		Data.RootPath = filepath.Join(".."+s, "."+s)
	}
	return Data.RootPath
}
