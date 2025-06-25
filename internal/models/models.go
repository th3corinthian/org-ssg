package models

type Config struct {
	Server	Server
	Site	Site
	Paths   Paths
}

// Server struct holds information about the server configuration
type Server struct {
	Host   string
	Port   int
	UseTLS bool
}

// Site struct holds information about the site configuration
type Site struct {
	Title       string
	Description string
	Author      string
}

// Paths struct holds information about the paths for the application
type Paths struct {
	Input    string
	Output   string
	Assets   string
	Templates string
}

//func StartServer() {}

//func ConfigTLS() {}

//func OrgToHTML() {}

//type DigestConfig struct {
//	Server struct {
//		Host		string		`yaml:"host"`
//		Port		int			`yaml:"port"`
//		UseTLS		bool		`yaml:"use_tls"`
//	} `yaml:"server"`
//	Site struct {
//		Title		string		`yaml:"title"`
//		Description	string		`yaml:"description"`
//		Author		string		`yaml:"author"`
//	} `yaml:"site"`
//	Paths struct {
//		Input		string		`yaml:"input"`
//		Output		string		`yaml:"output"`
//		Assets		string		`yaml:"assets"`
//		Templates	string		`yaml:"templates"`
//	} `yaml:"paths"`
//}
