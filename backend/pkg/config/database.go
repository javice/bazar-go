package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/javice/bazar-go/backend/internal/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("bazar.db"), &gorm.Config{})
    if err != nil {
        panic("Error al conectar a SQLite: " + err.Error())
    }
    // Migraciones
    DB.AutoMigrate(&models.Product{}, &models.Client{}, &models.Sale{})
}

func CloseDB() {
	// Cerramos la conexión con la base de datos.
    sqlDB, err := DB.DB()
    if err != nil {
        panic("Error al obtener la conexión de la base de datos: " + err.Error())
    }
    sqlDB.Close()
}