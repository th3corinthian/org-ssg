package new

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/th3corinthian/org-ssg/internal/models"
	"gopkg.in/yaml.v3"
)

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
	config := models.Config{
		Server: models.Server {
			Host:			"localhost",
			Port:			8080,
			UseTLS:			true,
		},
		Site: models.Site {
			Title:			"My Website",
			Description:	"A description of my website",
			Author:			"John Doe",
		},
		Paths: models.Paths {
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
