package build

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host		string		`yaml:"host"`
		Port		int			`yaml:"port"`
		UseTLS		bool		`yaml:"use_tls"`
	} `yaml:"server"`
	Paths struct {
		Input		string		`yaml:"input"`
		Output		string		`yaml:"output"`
		Assets		string		`yaml:"assets"`
		Templates	string		`yaml:"templates"`
	} `yaml:"paths"`
}

func loadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ConfigServer() {
	cfg, err := loadConfig("myproj/config.yml")
	if err != nil {
		log.Fatalf("[-] Error loading config: %v", err)
	}

	// create a server with the configuration values
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	if cfg.Server.UseTLS {
		log.Printf("[+] Starting server with TLS on %s", serverAddr)
		err := http.ListenAndServeTLS(serverAddr, "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatalf("[-] Error starting TLS server: %v", err)
		}
	} else {
		log.Printf("[+] Starting server on %s", serverAddr)
		err := http.ListenAndServe(serverAddr, nil)
		if err != nil {
			log.Fatalf("[-] Error starting server: %v", err)
		}
	}
}

func serverHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Server settings: Host - %s, Port - %d", r.Host , 8080)
}

func pathsHandle(w http.ResponseWriter, r *http.Request) {}

var handlers = map[string]http.HandlerFunc{
	"host": 		serverHandle,
	"port": 		serverHandle,
	"use_tls:":		serverHandle,

	"input":		pathsHandle,
	"output":		pathsHandle,
	"assets":		pathsHandle,
	"templates":	pathsHandle,
}
