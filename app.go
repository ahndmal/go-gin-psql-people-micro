package main

import (
	"fmt"
	"github.com/AndriiMaliuta/go-people-micro/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
	"net/http"
	"os"
)

func main() {
	config := pgx.ConnConfig{User: "hostman", Host: os.Getenv("DB_HOST"), Port: 5439, Database: "database", Password: os.Getenv("DB_PASS")}
	conn, err := pgx.Connect(config)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	defer conn.Close()
	var cats []model.Cat
	rows, err := conn.Query("select * from cats")
	if err != nil {
		log.Println(err)
	}

	// values
	values, err := rows.Values()
	if err != nil {
		return
	}
	fmt.Println(values)

	// scanning rows
	for rows.Next() {
		cat := model.Cat{}

		err := rows.Scan(&cat.Id, &cat.Name)
		cats = append(cats, cat)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println(cats)

	// GIN
	route := gin.Default()
	route.GET("/people", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"people": "test",
		})
	})
	err = route.Run()
	if err != nil {
		return
	}
}
