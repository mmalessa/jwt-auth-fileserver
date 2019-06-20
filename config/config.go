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
	// c.ServerTls = true
	// c.ServerPort = 443
	c.ServerTls = false
	c.ServerPort = 8080
	c.ServerTimeout = 5
	c.ServerDialTimeout = 10
	c.ServerTlsHandshakeTimeout = 5
	c.ServerFileCrt = "ssl/test.crt"
	c.ServerFileKey = "ssl/test.key"
	c.AuthApiType = "dev" // dev | noauth | jwt
	c.AuthApiEndpoint = "http://localhost/jwt.php"
	c.FilesRootDirectory = "./files"
	return c
}
