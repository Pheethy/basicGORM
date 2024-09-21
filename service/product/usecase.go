package product

import (
	"basicGORM/models"
	"context"

	"github.com/Microsoft/go-winio/pkg/guid"
)

type IProductUsecase interface {
	FetchOneProduct(ctx context.Context, productId guid.GUID) (*models.Product, error)
	FetchAllProducts(ctx context.Context) ([]*models.Product, error)
}
