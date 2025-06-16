package new

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server		string	`yaml:"server"`
	Port		int		`yaml:"port"`
	UseTLS		bool	`yaml:"use_tls"`
	Username	string	`yaml:"username"`
	Template	string	`yaml:"template"`
}

func NewProj(projectRoot string) error {
	if err := genDir(projectRoot); err != nil {
		return err
	}
	if err := genSubDir(projectRoot); err != nil {
		return err
	}
	return nil
}

func genDir(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	fmt.Println("[+] Directory created:", path)
	return nil
}

func genSubDir(base string) error {
	subdir := []string{
		"_layouts", "_posts", "_site", "css",
	}

	for _, d := range subdir {
		fullPath := filepath.Join(base, d)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return err
		}
	}; return nil
}

// below is to generate and populate "config.yml"
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
