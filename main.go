package main

import (
    "basicGORM/config"
    "basicGORM/route"
    productHandler "basicGORM/service/product/handlers"
    productRepository "basicGORM/service/product/repository"
    productUsecase "basicGORM/service/product/usecase"
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/sirupsen/logrus"
    "log"
    "net/url"
    "os"

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
    cfg := config.LoadConfig(envPath())
    query := url.Values{}
    query.Add("database", cfg.Db().DataBase())
    u := &url.URL{
        Scheme:   "sqlserver",
        User:     url.UserPassword(cfg.Db().Username(), cfg.Db().Password()),
        Host:     fmt.Sprintf("%s:%d", cfg.Db().Host(), cfg.Db().Port()),
        RawQuery: query.Encode(),
    }
    db, err := gorm.Open(sqlserver.Open(u.String()))
    if err != nil {
        log.Fatal(err)
        return
    }

    gormDB, err := db.DB()
    if err != nil {
        panic(err)
    }
    defer func(gormDB *sql.DB) {
        err := gormDB.Close()
        if err != nil {
            logrus.Println(err.Error())
        }
    }(gormDB)
    /* Init Repository */
    productRepo := productRepository.NewProductRepository(db)
    /* Init Usecase */
    productUs := productUsecase.NewProductUsecase(productRepo)
    /* Init Handler */
    productHand := productHandler.NewProductHandler(productUs)

    app := fiber.New(fiber.Config{
        AppName:      cfg.App().Name(),
        BodyLimit:    cfg.App().BodyLimit(),
        ReadTimeout:  cfg.App().ReadTimeOut(),
        WriteTimeout: cfg.App().WriteTimeOut(),
        JSONEncoder:  json.Marshal,
        JSONDecoder:  json.Unmarshal,
    })
    router := app.Group("/v1")
    routing := route.NewRoute(router)
    routing.RegisterProduct(productHand)

    if err := app.Listen(cfg.App().Url()); err != nil {
        log.Println(err.Error())
    }
}
