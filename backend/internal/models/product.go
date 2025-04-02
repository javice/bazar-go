package models


// Product representa un producto en la base de datos
// @Description Modelo de producto
type Product struct {
	ID          uint    `gorm:"primaryKey" json:"product_id"` // ID del producto
    Name        string  `gorm:"not null" json:"name"` // Nombre del producto
    Description string  `json:"description"`          // Descripción del producto
    Price       float64 `gorm:"not null" json:"price"` // Precio del producto
    Stock       int     `gorm:"default:0" json:"stock"` // Stock del producto
    Category    string  `json:"category"` // Categoría del producto
}