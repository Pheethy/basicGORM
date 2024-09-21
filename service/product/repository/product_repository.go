package repository

import (
    "basicGORM/models"
    "basicGORM/service/product"
    "context"
    "database/sql"
    "errors"

    "github.com/Microsoft/go-winio/pkg/guid"
    "gorm.io/gorm"
)

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.IProductRepository {
    return &productRepository{db: db}
}

func (r productRepository) FetchAllProducts(ctx context.Context) ([]*models.Product, error) {
    products := make([]*models.Product, 0)
    result := r.db.WithContext(ctx).Preload("Images", func(db *gorm.DB) *gorm.DB {
        return db.Raw(`
      SELECT
       images.id,
       images.filename,
       images.url,
       images.product_id,
       images.created_at,
       images.updated_at
      FROM
        images
      INNER JOIN
        products
      ON
        images.product_id = products.id
      `)
    }).Preload("Categories", func(db *gorm.DB) *gorm.DB {
        return db.Raw(`
      SELECT
        categories.id,
        categories.title,
        categories.product_id
      FROM
        (
          SELECT
            categories.*,
            products_categories.product_id "product_id"
          FROM
            categories
          LEFT JOIN
            products_categories
          ON
            products_categories.category_id = categories.id
        ) AS categories
  `)
    }).Raw(`
      SELECT
        products.id,
        products.title,
        products.price,
        products.description,
        products.created_at,
        products.updated_at
      FROM
        products
    `).Find(&products)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
    }

    return products, nil
}

func (r productRepository) FetchOneProduct(ctx context.Context, productId guid.GUID) (*models.Product, error) {
    productItem := new(models.Product)
    sqlProduct := `
      SELECT
        products.id,
        products.title,
        products.price,
        products.description,
        products.created_at,
        products.updated_at
      FROM
        products
      WHERE
        products.id = @productId
  	`
    sqlCategories := `
      SELECT
        categories.id,
        categories.title,
        categories.product_id
      FROM
        (
          SELECT
            categories.*,
            products_categories.product_id "product_id"
          FROM
            categories
          LEFT JOIN
            products_categories
          ON
            products_categories.category_id = categories.id
		  WHERE
			products_categories.product_id = @productId
        ) AS categories
	`
    sqlImages := `
      SELECT
       images.id,
       images.filename,
       images.url,
       images.product_id,
       images.created_at,
       images.updated_at
      FROM
        images
      INNER JOIN
        products
      ON
        images.product_id = products.id
	  WHERE
		products.id = @productId
	`
    result := r.db.WithContext(ctx).Preload("Categories", func(db *gorm.DB) *gorm.DB {
        return db.Raw(sqlCategories, sql.Named("productId", productId.String()))
    }).Preload("Images", func(db *gorm.DB) *gorm.DB {
        return db.Raw(sqlImages, sql.Named("productId", productId.String()))
    }).Raw(sqlProduct,
        sql.Named("productId", productId.String()),
    ).Find(productItem)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
    }

    return productItem, nil
}

func (r productRepository) CreateProducts(ctx context.Context, product *models.Product) error {
    tx := r.db.Begin()
    if err := tx.Error; err != nil {
        return err
    }
    sqlc := `
    INSERT INTO products (
      products.id,
      products.title,
      products.description,
      products.price,
      products.created_at,
      products.updated_at
    ) VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
  `
    result := tx.WithContext(ctx).Raw(sqlc,
        product.Id,
        product.Title,
        product.Description,
        product.Price,
        product.CreatedAt,
        product.UpdatedAt,
    )

    if result.Error != nil {
        tx.Rollback()
        return result.Error
    }

    if result.RowsAffected == 0 {
        tx.Rollback()
        return errors.New("create products failed")
    }

    return tx.Commit().Error
}

func (r productRepository) UpdateProducts(ctx context.Context, product *models.Product) error {
    tx := r.db.Begin()
    if err := tx.Error; err != nil {
        return err
    }
    sqlc := `
    UPDATE
      products
    SET
      products.title = ?,
      products.description = ?,
      products.price = ?,
      products.updated_at = ?
    WHERE
      products.id = ?
  `

    result := tx.WithContext(ctx).Raw(sqlc,
        product.Title,
        product.Description,
        product.Price,
        product.UpdatedAt,
    )

    if result.Error != nil {
        tx.Rollback()
        return result.Error
    }
    if result.RowsAffected == 0 {
        tx.Rollback()
        return errors.New("updated failed")
    }

    return tx.Commit().Error
}

func (r productRepository) DeleteProduct(ctx context.Context, productId string) error {
    tx := r.db.Begin()
    sqlc := `
    DELETE FROM
      products
    WHERE
      products.id = ?
  `
    result := tx.WithContext(ctx).Raw(sqlc, productId)
    if result.Error != nil {
        tx.Rollback()
        return result.Error
    }
    if result.RowsAffected == 0 {
        tx.Rollback()
        return errors.New("delete failed")
    }

    return tx.Commit().Error
}
