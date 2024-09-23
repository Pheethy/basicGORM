package main

import (
	"basicGORM/config"
	"basicGORM/route"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	productHandler "basicGORM/service/product/handlers"
	productRepository "basicGORM/service/product/repository"
	productUsecase "basicGORM/service/product/usecase"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	/* Get Config */
	cfg := config.LoadConfig(envPath())
	/* Connecting Database */
	db, err := gorm.Open(sqlserver.Open(cfg.Db().Url()))
	if err != nil {
		log.Fatal(err)
		return
	}
	gormDB, err := db.DB()
	if err != nil {
		logrus.Panic(err)
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			logrus.Panic(err)
		}
	}(gormDB)
	gormDB.Ping()
	log.Println("ping Pass")

	/* Init Repository */
	productRepo := productRepository.NewProductRepository(db)
	/* Init Usecase */
	productUs := productUsecase.NewProductUsecase(productRepo)
	/* Init Handler */
	productHand := productHandler.NewProductHandler(productUs)

	/* Init Fiber Server */
	app := fiber.New(fiber.Config{
		AppName:      cfg.App().Name(),
		BodyLimit:    cfg.App().BodyLimit(),
		ReadTimeout:  cfg.App().ReadTimeOut(),
		WriteTimeout: cfg.App().WriteTimeOut(),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	router := app.Group("/v1")
	r := route.NewRoute(router)
	r.RegisterProduct(productHand)

	if err := app.Listen(cfg.App().Url()); err != nil {
		log.Println(err.Error())
	}
}
