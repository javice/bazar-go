package tests

import (
	"testing"

	
)


func TestMainSuite(t *testing.T) {
    // Ejecutar los tests en el orden deseado
    t.Run("Client Tests // Create Client", TestCreateClient)
    t.Run("Client Tests // Get Client", TestGetClients)
    t.Run("Product Tests // Create Product", TestCreateProduct)
    t.Run("Sale Tests // Create Sale", TestCreateSale)

}