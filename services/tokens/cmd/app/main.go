package main

import (
	"log"
	"medods/pkg/jwt"
	"medods/services/tokens/internal/db"
	rtr "medods/services/tokens/internal/router"
	"medods/services/tokens/internal/services/tokensService"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWTConfig *jwt.Config `yaml:"jwt"`
	RTRConfig *rtr.Config `yaml:"rtr"`
	DBConfig  *db.Config  `yaml:"db"`
}

func readConfig(filename string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	cfg, err := readConfig("./cfg.yml")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Config file read successfully")
	jwt, err := jwt.New(cfg.JWTConfig)
	if err != nil {
		log.Fatalln("Failed to create jwt: " + err.Error())
	}
	log.Println("JWT created successfully")
	db, err := db.New(cfg.DBConfig)
	if err != nil {
		log.Fatalln("Failed to connect to db: " + err.Error())
	}
	log.Println("DB connected successfully")
	tokensService, _ := tokensService.New(db, &jwt)
	router, err := rtr.New(cfg.RTRConfig, tokensService)
	if err != nil {
		log.Fatalln("Failed to host router:", err.Error())
	}
	err = router.Listen()
	if err != nil {
		log.Fatalln("Failed listen to router:", err)
	}
	log.Printf("Router is listening on %v:%v\n", router.Config.Host, router.Config.Port)
}
