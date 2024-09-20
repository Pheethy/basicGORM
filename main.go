package main

import (
	"basicGORM/config"
	_product_repo "basicGORM/service/product/repository"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig(envPath())

	query := url.Values{}
	query.Add("database", cfg.Db().DataBase())

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(cfg.Db().Username(), cfg.Db().Password()),
		Host:     fmt.Sprintf("%s:%d", cfg.Db().Host(), cfg.Db().Port()),
		RawQuery: query.Encode(),
	}

	db, err := gorm.Open(sqlserver.Open(u.String()),
		&gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
		return
	}

	gormDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer gormDB.Close()

	/* Ping DB */
	gormDB.Ping()
	log.Println("ping pass")
	productRepo := _product_repo.NewProductRepository(db)
	// id, _ := guid.FromString("6b5cd101-7d72-4570-bda8-17eddd9a1547")
	product, err := productRepo.FetchAllProducts(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	for _, item := range product {
		log.Println("product:", item)
		for _, img := range item.Images {
			log.Println("Images:", img)
		}
		log.Println("Categories", item.Categories)
	}
}
