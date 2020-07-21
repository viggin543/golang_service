package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type Fruit struct {
	Id  int `json:"id"`
	Name string `json:"name"`
}

var con *sql.DB


func SetupServer() *gin.Engine {
	host := getDBHost()
	con = connectToMysql(host)
	r := gin.Default()
	r.GET("/fruits", fruits)
	r.GET("/ping", pingEndpoint)
	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func connectToMysql(host string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root:password@tcp(%s:3306)/frutas", host))
	if err != nil {
		panic("failed to open a mysql connection")
	}
	return db
}

func getDBHost() string {
	host, isSet := os.LookupEnv("DB_HOST")
	if !isSet {
		host = "host.docker.internal"
	}
	return host
}

func fruits(c *gin.Context) {
	fruits := getFruits()
	c.JSON(http.StatusOK, fruits)
}

func getFruits() []Fruit {
	rows, _ := con.Query("SELECT * FROM fruits")
	var fruits []Fruit
	for rows.Next() {
		var r Fruit
		_ = rows.Scan(&r.Id, &r.Name)
		fruits = append(fruits, r)
	}
	return fruits
}

