package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_HttpDelivery "ddd-to-do-list/internal/delivery/handler"
	routers "ddd-to-do-list/internal/delivery/router"
	"ddd-to-do-list/internal/infrastructure/database/mysql/model"
	_Repo "ddd-to-do-list/internal/infrastructure/database/mysql/repository"
	_Ucase "ddd-to-do-list/internal/usecase"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		if err = godotenv.Load(".env"); err != nil {
			return
		}
	}
	dbHost := os.Getenv(`MYSQL_HOST`)
	dbPort := os.Getenv(`MYSQL_PORT`)
	dbUser := os.Getenv(`MYSQL_USER`)
	dbPass := os.Getenv(`MYSQL_PASSWORD`)
	dbName := os.Getenv(`MYSQL_DBNAME`)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	ds := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.TransactionDTO{}, &model.UserDTO{})

	e := echo.New()
	middL := _HttpDelivery.InitMiddleware()
	e.Use(middL.CORS)
	ar := _Repo.NewMysqlUserRepository(dbConn)
	tr := _Repo.NewMysqlTransactionRepository(dbConn)
	au := _Ucase.NewUserUsecase(ar)
	tu := _Ucase.NewTransactionUsecase(ar, tr)
	_HttpDelivery.NewHandler(au, tu)

	routers.Router(e, au, tu)

	log.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))
}
