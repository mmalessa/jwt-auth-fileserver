package config

type config struct {
	ServerPort         int
	JwtApiEndpoint     string
	FilesRootDirectory string
}

func Init() *config {
	c := new(config)
	// TODO - read from file
	c.ServerPort = 8080
	c.JwtApiEndpoint = "http://localhost/jwt.php"
	c.FilesRootDirectory = "./files"
	return c
}
