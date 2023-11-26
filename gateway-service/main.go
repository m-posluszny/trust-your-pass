package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Password struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Strength int    `json:"strength"`
}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// temporary collection that replaces db
var passwords = []Password{
	{ID: 1, Password: "intel1", Strength: 0},
	{ID: 2, Password: "elyass15@ajilent-ci", Strength: 2},
	{ID: 3, Password: "hodygid757#$!23w", Strength: 1},
}

func main() {

	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	/*conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}*/

	router := gin.Default()
	router.GET("/api/v1/passwords", getPasswords)
	router.POST("/api/v1/passwords", postPassword)

	router.Run(config.ServerAddress)
}

func getPasswords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, passwords)
}

func postPassword(c *gin.Context) {
	var newPassword Password
	if err := c.BindJSON(&newPassword); err != nil {
		return
	}
	newPassword.ID = passwords[len(passwords)-1].ID + 1
	passwords = append(passwords, newPassword)
	c.IndentedJSON(http.StatusCreated, newPassword)
}
