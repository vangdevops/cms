package memory_test

import (
	"cms/internal/entity"
	"cms/internal/repository/memory"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestProductAddSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product Add Return Error")
	}
}

func TestProductAddNameError(t *testing.T) {
	ProductName := gofakeit.Name()
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, 0)
	if err == nil {
		t.Errorf("Test Product Add Return nil")
	}
}

func TestProductAddWeightError(t *testing.T) {
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add("", ProductWeight)
	if err == nil {
		t.Errorf("Test Product Add Return nil")
	}
}

func TestProductSetDescriptionSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductDescription := gofakeit.Extension()
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd SetDescription return Error")
	}
	err = repo.SetDescription(ProductName, ProductWeight, ProductDescription)
	if err != nil {
		t.Errorf("Test Product SetDescription return Error")
	}
}

func TestProductSetDescriptionError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductDescription := gofakeit.Extension()
	repo := memory.NewProductRepository()
	err := repo.SetDescription(ProductName, ProductWeight, ProductDescription)
	if err == nil {
		t.Errorf("Test Product SetDescription return Error")
	}
}

func TestProductSetImageSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductImage := gofakeit.Extension()
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd SetImage return Error")
	}
	err = repo.SetImage(ProductName, ProductWeight, ProductImage)
	if err != nil {
		t.Errorf("Test Product SetImage return Error")
	}
}

func TestProductSetImageError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductImage := gofakeit.Extension()
	repo := memory.NewProductRepository()
	err := repo.SetImage(ProductName, ProductWeight, ProductImage)
	if err == nil {
		t.Errorf("Test Product SetImage return nil")
	}
}

func TestProductSetPriceSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductPrice := gofakeit.Extension()
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd SetPrice return Error")
	}
	err = repo.SetImage(ProductName, ProductWeight, ProductPrice)
	if err != nil {
		t.Errorf("Test Product SetPrice return Error")
	}
}

func TestProductSetPriceError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductPrice := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.SetPrice(ProductName, ProductWeight, ProductPrice)
	if err == nil {
		t.Errorf("Test Product SetPrice return nil")
	}
}

func TestProductDeleteSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd Delete return Error")
	}
	err = repo.Delete(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product Delete return Error")
	}
}

func TestProductDeleteError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Delete(ProductName, ProductWeight)
	if err == nil {
		t.Errorf("Test Product Delete return nil")
	}
}

func TestProductGetSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd Get return Error")
	}
	_, err = repo.Get(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product Get return Error")
	}
}

func TestProductGetError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}
	repo := memory.NewProductRepository()
	testproduct, err := repo.Get(ProductName, ProductWeight)
	if err == nil {
		t.Errorf("Test Product Get return nil")
	}
	if testproduct == product {
		t.Errorf("Test Product Get is equal product")
	}
}

func TestProductListSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd List return Error")
	}
	_, err = repo.List()
	if err != nil {
		t.Errorf("Test Product List return Error")
	}
}

func TestProductListError(t *testing.T) {
	repo := memory.NewProductRepository()
	_, err := repo.List()
	if err == nil {
		t.Errorf("Test Product List return nil")
	}
}
