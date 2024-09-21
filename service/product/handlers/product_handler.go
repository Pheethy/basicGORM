package handlers

import (
    "basicGORM/service/product"
    "github.com/gofiber/fiber/v2"
    "net/http"

    "github.com/Microsoft/go-winio/pkg/guid"
)

type productHandler struct {
    productUs product.IProductUsecase
}

func NewProductHandler(productUs product.IProductUsecase) product.IProductHandler {
    return &productHandler{
        productUs: productUs,
    }
}

func (p *productHandler) FetchOneProduct(c *fiber.Ctx) error {
    ctx := c.UserContext()
    productId, _ := guid.FromString(c.Params("product_id"))
    productItem, err := p.productUs.FetchOneProduct(ctx, productId)
    if err != nil {
        return fiber.NewError(http.StatusInternalServerError, err.Error())
    }

    resp := map[string]interface{}{
        "product": productItem,
    }
    return c.Status(http.StatusOK).JSON(resp)
}

func (p *productHandler) FetchAllProducts(c *fiber.Ctx) error {
    ctx := c.UserContext()
    products, err := p.productUs.FetchAllProducts(ctx)
    if err != nil {
        return fiber.NewError(http.StatusInternalServerError, err.Error())
    }

    resp := map[string]interface{}{
        "products": products,
    }
    return c.Status(http.StatusOK).JSON(resp)
}
