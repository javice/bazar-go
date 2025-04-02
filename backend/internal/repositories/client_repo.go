package repositories

import (
    "github.com/javice/bazar-go/backend/internal/models"
    "github.com/javice/bazar-go/backend/pkg/config"
)

func CreateClient(client *models.Client) error {
    return config.DB.Create(client).Error
}

func GetAllClients() ([]models.Client, error) {
    var clients []models.Client
    err := config.DB.Find(&clients).Error
    return clients, err
}

func GetClientById(id uint) (*models.Client, error) {
    var client models.Client
    err := config.DB.First(&client, id).Error
    return &client, err
}

func UpdateClient(client *models.Client) error {
	
	return config.DB.Save(client).Error
}

func DeleteClient(id uint) (*models.Client, error) {
    var client models.Client
    err := config.DB.First(&client, id).Error
    if err != nil {
        return nil, err
    }

    err = config.DB.Delete(&client).Error
    if err != nil {
        return nil, err
    }

    return &client, nil
}