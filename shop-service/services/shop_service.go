package services

import (
	"github.com/putrapratamanst/ecommerce/shop-service/models"
	"github.com/putrapratamanst/ecommerce/shop-service/repositories"
)

type ShopService interface {
	CreateShop(shop *models.Shop) error
	GetShopByID(shopID int) (*models.Shop, error)
    UpdateShop(shop *models.Shop) error
    DeleteShop(shopID int) error
    GetAllShops() ([]models.Shop, error)
}

type shopService struct {
    shopRepo *repositories.ShopRepository
}

func NewShopService(shopRepo *repositories.ShopRepository) ShopService {
    return &shopService{shopRepo: shopRepo}
}

func (s *shopService) CreateShop(shop *models.Shop) error {
    return s.shopRepo.Create(shop)
}

func (s *shopService) GetShopByID(shopID int) (*models.Shop, error) {
    return s.shopRepo.FindByID(shopID)
}

func (s *shopService) UpdateShop(shop *models.Shop) error {
    return s.shopRepo.Update(shop)
}

func (s *shopService) DeleteShop(shopID int) error {
    return s.shopRepo.Delete(shopID)
}


func (s *shopService) GetAllShops() ([]models.Shop, error) {
    return s.shopRepo.FindAll()
}