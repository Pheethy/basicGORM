package route

import (
    "basicGORM/service/product"

    "github.com/gofiber/fiber/v2"
)

type Route struct {
    e fiber.Router
}

func NewRoute(e fiber.Router) *Route {
    return &Route{
        e: e,
    }
}

func (r *Route) RegisterProduct(handler product.IProductHandler) {
    r.e.Get("/product/list", handler.FetchAllProducts)
    r.e.Get("/product/:product_id", handler.FetchOneProduct)
}
