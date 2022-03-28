package main

import (
	"fmt"
	"github.com/AndriiMaliuta/go-people-micro/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=hostman password=%s dbname=database port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PASS"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	var persons []model.Person
	db.Table("persons").Where("gender = ?", "male").Limit(100).Find(&persons)
	//fmt.Println(persons)

	// ================
	// GIN
	route := gin.Default()
	route.GET("/persons", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"persons": persons,
		})
	})
	err = route.Run()
	if err != nil {
		return
	}
}

func DbPsql() {
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
	rows, err := conn.Query("select id, name from public.cats")
	if err != nil {
		log.Println(err)
	}
	// scanning rows
	for rows.Next() {
		cat := model.Cat{}

		err := rows.Scan(&cat.Id, &cat.Name)
		cats = append(cats, cat)
		if err != nil {
			log.Println(err)
		}
	}

	//fmt.Println(cats)
}
