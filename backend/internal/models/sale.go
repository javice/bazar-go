package models

// Sale representa una venta en la base de datos
// @Description Modelo de venta
type Sale struct {
	ID          uint      `gorm:"primaryKey" json:"sale_id"`
	ClientID    uint      `gorm:"not null" json:"client_id"` // ID del cliente
    Client      Client    `gorm:"foreignKey:ClientID" json:"client"` // Cliente asociado a la venta
    Products    []Product `gorm:"many2many:sale_products;" json:"products"` // Productos asociados a la venta
    Quantity    int       `gorm:"not null" json:"quantity"` // Cantidad de productos vendidos
    TotalAmount float64   `gorm:"not null" json:"total_amount"` // Monto total de la venta
    Date        string    `gorm:"not null" json:"date"` // Fecha de la venta
}