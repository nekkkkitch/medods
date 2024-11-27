package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Router struct {
	App           *fiber.App
	Config        *Config
	TokensService TokensService
}

type Config struct {
	Host string `yaml:"router_host" env-prefix:"ROUTERHOST"`
	Port string `yaml:"router_port" env-prefix:"ROUTERPORT"`
}

type TokensService interface {
	GetTokens(id uuid.UUID, ip string) (string, string, error)
	RefreshTokens(accessToken, refreshToken string, ip string) (string, string, error)
}

func New(cfg *Config, tkns TokensService) (*Router, error) {
	app := fiber.New()
	router := Router{App: app, Config: cfg, TokensService: tkns}
	router.App.Get("/tokens", router.GetTokens())
	router.App.Get("/refreshtokens", router.RefreshTokens())
	return &router, nil
}

func (r *Router) Listen() error {
	err := r.App.Listen(r.Config.Port)
	if err != nil {
		log.Println("Cannoct listen router:", err)
		return err
	}
	return nil
}

func (r *Router) GetTokens() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Query("id"))
		if err != nil {
			log.Println("Failed to get id from query:", err)
			return err
		}
		ip := c.IP()
		access, refresh, err := r.TokensService.GetTokens(id, ip)
		if err != nil {
			log.Println("Failed to get new tokens")
			return err
		}
		c.Response().Header.Add("X-Access-Token", access)
		c.Response().Header.Add("X-Refresh-Token", refresh)
		return nil
	}
}

func (r *Router) RefreshTokens() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		access, refresh, err := r.TokensService.RefreshTokens(c.GetReqHeaders()["X-Access-Token"][0], c.GetReqHeaders()["X-Refresh-Token"][0], ip)
		if err != nil {
			log.Println("Failed to get new tokens")
			return err
		}
		c.Response().Header.Add("X-Access-Token", access)
		c.Response().Header.Add("X-Refresh-Token", refresh)
		return nil
	}
}
