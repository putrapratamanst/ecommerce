package repositories

import (
	"github.com/putrapratamanst/ecommerce/shop-service/models"
	"gorm.io/gorm"
)

type ShopRepository struct {
    db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
    return &ShopRepository{db: db}
}

func (r *ShopRepository) Create(shop *models.Shop) error {
    return r.db.Create(shop).Error
}

func (r *ShopRepository) FindByID(shopID int) (*models.Shop, error) {
    var shop models.Shop
    err := r.db.First(&shop, shopID).Error
    return &shop, err
}

func (r *ShopRepository) Update(shop *models.Shop) error {
    return r.db.Save(shop).Error
}

func (r *ShopRepository) Delete(shopID int) error {
    return r.db.Delete(&models.Shop{}, shopID).Error
}

func (r *ShopRepository) FindAll() ([]models.Shop, error) {
    var shops []models.Shop
    err := r.db.Find(&shops).Error
    return shops, err
}