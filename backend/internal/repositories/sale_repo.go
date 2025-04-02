package repositories

import (
    "github.com/javice/bazar-go/backend/internal/models"
    "github.com/javice/bazar-go/backend/pkg/config"
)

func CreateSale(sale *models.Sale) error {
    return config.DB.Create(sale).Error
}

func GetAllSales() ([]models.Sale, error) {
    var sales []models.Sale
    err := config.DB.Preload("Client").Preload("Products").Find(&sales).Error
    return sales, err
}

func GetSaleById(id uint) (*models.Sale, error) {
    var sale models.Sale
    err := config.DB.First(&sale, id).Error
    return &sale, err
}

func UpdateSale(sale *models.Sale) error {
	
	return config.DB.Save(sale).Error
}

func DeleteSale(id uint) (*models.Sale, error) {
    var sale models.Sale
    err := config.DB.First(&sale, id).Error
    if err != nil {
        return nil, err
    }

    err = config.DB.Delete(&sale).Error
    if err != nil {
        return nil, err
    }

    return &sale, nil
}