package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	userRepositoryMock := repository.NewUserRepositoryMock()
	_ = userRepositoryDB
	_ = userRepositoryMock

	userService := service.NewUserService(userRepositoryDB)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{userID:[0-9]+}", userHandler.GetUser).Methods(http.MethodGet)

	log.Printf("Banking servicer started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
