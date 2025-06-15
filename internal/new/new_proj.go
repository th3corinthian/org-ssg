package new

import (
	"fmt"
	"os"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server		string	`yaml:"server"`
	Port		int		`yaml:"port"`
	UseTLS		bool	`yaml:"use_tls"`
	Username	string	`yaml:"username"`
	Template	string	`yaml:"template"`
}

func NewProj(projName string) {
	err := os.Mkdir(projName, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	dirPath := projName
	fmt.Println("[+] Directory created:", dirPath)
}

func GenConfig(dirPath string) {
	// generating config.yml
	configPath := dirPath + "/config.yml"
	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fmt.Println("[+] Config file created:", configPath)

	_, err = file.WriteString(genYAMLConfig())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Config file generated...")
}

func genYAMLConfig() (string) {
	config := Config{
		Server:		"localhost",
		Port:		8080,
		UseTLS:		true,
		Username:	"admin",
	}

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(yamlData)
}
