package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Password struct {
	id       int    `json:"id"`
	password string `json:"password"`
	strength int    `json:"strength"`
}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var db *sql.DB

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
/*var passwords = []Password{
	{id: 1, password: "intel1", strength: 0},
	{id: 2, password: "elyass15@ajilent-ci", strength: 2},
	{id: 3, password: "hodygid757#$!23w", strength: 1},
}*/

func main() {

	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	router := gin.Default()
	router.GET("/api/v1/passwords", getPasswords)
	router.POST("/api/v1/passwords", postPassword)

	router.Run(config.ServerAddress)
}

func getPasswords(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	rows, err := db.Query("SELECT id, password, strength FROM passwords")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var passwords []Password
	for rows.Next() {
		var a Password
		err := rows.Scan(&a.id, &a.password, &a.strength)
		if err != nil {
			log.Fatal(err)
		}
		passwords = append(passwords, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, passwords)
}

func postPassword(c *gin.Context) {
	var newPassword Password
	if err := c.BindJSON(&newPassword); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO passwords (password, strength) VALUES ($2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newPassword.id, newPassword.password, newPassword.strength); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newPassword)
}
