package config

type config struct {
	ServerPort                int
	ServerTls                 bool
	ServerTimeout             int
	ServerDialTimeout         int
	ServerTlsHandshakeTimeout int
	ServerFileCrt             string
	ServerFileKey             string
	AuthApiEndpoint           string
	AuthApiType               string
	FilesRootDirectory        string
}

func Init() *config {
	c := new(config)
	// TODO - read this from YAML file
	c.ServerPort = 8080
	c.ServerTls = false
	c.ServerTimeout = 5
	c.ServerDialTimeout = 10
	c.ServerTlsHandshakeTimeout = 5
	c.ServerFileCrt = ""
	c.ServerFileKey = ""
	c.AuthApiEndpoint = "http://localhost/jwt.php"
	c.AuthApiType = "JWT"
	c.FilesRootDirectory = "./files"
	return c
}
