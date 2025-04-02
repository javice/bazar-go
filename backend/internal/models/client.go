package models


// Client representa un cliente en la base de datos
// @Description Modelo de cliente
type Client struct {
    ID      uint   `gorm:"primaryKey" json:"client_id"` // ID del cliente
	Name    string `gorm:"not null" json:"name"` // Nombre del cliente
    Email   string `gorm:"unique" json:"email"` // Correo electrónico del cliente
    Phone   string `json:"phone"` // Teléfono del cliente
    Address string `json:"address"` // Dirección del cliente
}