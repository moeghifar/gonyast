package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Env struct {
	Port        int    `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	APIToken    string `mapstructure:"API_TOKEN"`
}

func ReadEnvViper() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.BindEnv("port")

	fmt.Println(os.Getenv("PORT") + " <> " + os.Getenv("SERVICE_NAME"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	envData := &Env{}

	fmt.Println(viper.GetViper().AllSettings())

	err = viper.Unmarshal(envData)
	if err != nil {
		panic(fmt.Errorf("failed unmarshal config file: %w", err))
	}

	return envData
}

func main() {
	// env := ReadEnv()
	env := ReadEnvViper()
	srv := fiber.New()

	srv.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"healthy": true,
		})
	})

	fmt.Printf("run %s service with http server in port %d\n", env.ServiceName, env.Port)

	srv.Listen(fmt.Sprintf(":%d", env.Port))
}
