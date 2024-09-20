package product

import (
	"basicGORM/models"
	"context"

	"github.com/Microsoft/go-winio/pkg/guid"
)

type IProductRepository interface {
	FetchOneProduct(ctx context.Context, productId guid.GUID) (*models.Product, error)
	FetchAllProducts(ctx context.Context) ([]*models.Product, error)
	CreateProducts(ctx context.Context, product *models.Product) error
	UpdateProducts(ctx context.Context, product *models.Product) error
}
