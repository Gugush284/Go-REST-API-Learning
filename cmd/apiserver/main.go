package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Gugush284/Go-server.git/internal/app/apiserver"
	_ "github.com/go-sql-driver/mysql"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "comfig-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(nil)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
