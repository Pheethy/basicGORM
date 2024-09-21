package usecase

import (
    "basicGORM/models"
    "basicGORM/service/product"
    "context"

    "github.com/Microsoft/go-winio/pkg/guid"
)

type productUsecase struct {
    productRepo product.IProductRepository
}

func NewProductUsecase(productRepo product.IProductRepository) product.IProductUsecase {
    return &productUsecase{
        productRepo: productRepo,
    }
}

func (p *productUsecase) FetchOneProduct(ctx context.Context, productId guid.GUID) (*models.Product, error) {
    return p.productRepo.FetchOneProduct(ctx, productId)
}

func (p *productUsecase) FetchAllProducts(ctx context.Context) ([]*models.Product, error) {
    return p.productRepo.FetchAllProducts(ctx)
}
