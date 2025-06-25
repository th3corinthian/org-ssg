package build

import (
	"log"
	"os"

	"github.com/th3corinthian/org-ssg/internal/models"
	"github.com/th3corinthian/org-ssg/server"
	"gopkg.in/yaml.v3"
)

func loadConfig(filepath string) (*models.Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg models.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

//func configTLS(cfg *models.Config) {
//	if cfg.Server.UseTLS {
//		server.ListenAndPut()
//	} else {
//		fmt.Println("error")
//	}
//}

func StartServer() {
	cfg, err := loadConfig("myproj/config.yml")
	if err != nil {
		log.Fatal("[-] Error loading config: %v\n", err)
	}

	server.Start(cfg)
}



//type Config struct {
//	Actions []string	`yaml:"actions"`
//}

//func ConfigServer() {
//	cfg, err := loadConfig("myproj/config.yml")
//	if err != nil {
//		log.Fatalf("[-] Error loading config: %v", err)
//	}
//
//	for functionName, fn := range cfg.FunctionsMap {
//		fn, exists := actionMap[action]
//		if exists {
//			fn()
//		} else {
//			fmt.Println("No handler action", action)
//		}
//	}
//
//	// create a server with the configuration values
//	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
//
//	if cfg.Server.UseTLS {
//		log.Printf("[+] Starting server with TLS on %s", serverAddr)
//		err := http.ListenAndServeTLS(serverAddr, "cert.pem", "key.pem", nil)
//		if err != nil {
//			log.Fatalf("[-] Error starting TLS server: %v", err)
//		}
//	} else {
//		log.Printf("[+] Starting server on %s", serverAddr)
//		err := http.ListenAndServe(serverAddr, nil)
//		if err != nil {
//			log.Fatalf("[-] Error starting server: %v", err)
//		}
//	}
//}
