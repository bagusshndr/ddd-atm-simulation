package repository

// import (
// 	"ddd-atm-simulation/internal/aggregate"
// 	"ddd-atm-simulation/internal/repository"
// 	"errors"
// 	"regexp"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/suite"
// )

// type productRepositoryMysqlTest struct {
// 	suite.Suite
// 	mock              sqlmock.Sqlmock
// 	productRepository repository.ProductRepository
// }

// func (t *productRepositoryMysqlTest) TestGetProduct() {
// 	product := aggregate.RebuildProduct(1, "bagus@bagus.com", 5)
// 	query := `SELECT id, name, price FROM products`
// 	t.Run("success", func() {
// 		rows := sqlmock.NewRows([]string{
// 			"id",
// 			"name",
// 			"price",
// 		}).AddRow(
// 			product.ID,
// 			product.Name,
// 			product.Price,
// 		)

// 		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

// 		actualProduct, err := t.productRepository.GetProduct()

// 		t.NotNil(actualProduct)
// 		t.NoError(err)
// 	})
// 	t.Run("failed", func() {
// 		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

// 		actualActivity, err := t.productRepository.GetProduct()

// 		t.Nil(actualActivity)
// 		t.Error(err)
// 	})
// }

// func (t *productRepositoryMysqlTest) TestGetProductByID() {
// 	product := aggregate.RebuildProduct(1, "bagus@bagus.com", 5)
// 	query := `SELECT id, name, price FROM products WHERE id = ? LIMIT 1`
// 	rows := sqlmock.NewRows([]string{
// 		"id",
// 		"name",
// 		"price",
// 	}).AddRow(
// 		product.ID,
// 		product.Name,
// 		product.Price,
// 	)
// 	t.Run("success", func() {

// 		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(product.ID).WillReturnRows(rows)

// 		actualProduct, _ := t.productRepository.GetProductByID(product.ID)

// 		t.NotNil(actualProduct)
// 	})

// 	t.Run("failed", func() {
// 		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(product.ID).WillReturnError(errors.New("error"))

// 		actualProduct, err := t.productRepository.GetProductByID(product.ID)

// 		t.Nil(actualProduct)
// 		t.Error(err)
// 	})

// }

// func (t *productRepositoryMysqlTest) TestDelete() {
// 	review := aggregate.RebuildReviews(1, 1, 5)
// 	query := `DELETE FROM products WHERE id = ?`
// 	t.Run("success", func() {
// 		t.mock.ExpectExec(query).WithArgs(review.ID).WillReturnResult(sqlmock.NewResult(1, 1))

// 		err := t.productRepository.DeleteProduct(review.ID)

// 		t.NoError(err)
// 	})

// 	t.Run("failed", func() {
// 		t.mock.ExpectExec(query).WithArgs(review.ID).WillReturnError(errors.New("error"))

// 		err := t.productRepository.DeleteProduct(review.ID)

// 		t.Error(err)
// 	})
// }

// func TestProductRepositoryMySQL(t *testing.T) {
// 	db, mock, _ := sqlmock.New()

// 	suite.Run(t, &productRepositoryMysqlTest{
// 		mock:              mock,
// 		productRepository: NewMysqlProductRepository(db),
// 	})
// }
