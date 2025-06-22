package new

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server		Server		`yaml:"server"`
	Site		Site		`yaml:"site"`
	Paths		Paths		`yaml:"paths"`
}

type Site struct {
	Title		string		`yaml:"title"`
	Description	string		`yaml:"description"`
	Author		string		`yaml:author`
}

type Server struct {
	Host		string		`yaml:"host"`
	Port		int 		`yaml:"port"`
	UseTLS		bool 		`yaml:"use_tls"`
}

type Paths struct {
	Input		string		`yaml:"input"`
	Output		string		`yaml:"output"`
	Assets		string		`yaml:"assets"`
	Templates	string		`yaml:"templates"`
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
		Server: Server{
			Host:			"localhost",
			Port:			8080,
			UseTLS:			true,
		},
		Site: Site{
			Title:			"My Website",
			Description:	"A description of my website",
			Author:			"John Doe",
		},
		Paths: Paths{
			Input:			"/path/to/input",
			Output:			"/path/to/output",
			Assets:			"/path/to/assets",
			Templates:		"/path/to/templates",
		},
	}

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(yamlData)
}
