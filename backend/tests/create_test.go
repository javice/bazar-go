package tests

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
    "strings"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/javice/bazar-go/backend/internal/models"
    "github.com/javice/bazar-go/backend/internal/handlers"
    "github.com/javice/bazar-go/backend/pkg/config"
)

var testDB *gorm.DB

func setupTestDB() *gorm.DB {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&models.Sale{}, &models.Client{}, &models.Product{})
    testDB = db
    config.DB = db
    return db
}

func TestCreateClient(t *testing.T) {
    db := setupTestDB()
    assert.NotNil(t, db)

    // Crear cliente
    client := models.Client{ID: 1, Name: "Juan Lopez", Email: "juanlopez@test.com", Phone: "654321587", Address: "Calle del Test 11"}
    err := db.Create(&client).Error
    assert.NoError(t, err)
}

func TestCreateProduct(t *testing.T) {
    db := setupTestDB()
    assert.NotNil(t, db)

    // Crear producto
    product := models.Product{ID: 1, Name: "Lámpara", Description: "Lámpara de mesa", Category: "Muebles", Price: 15.99, Stock: 10}
    err := db.Create(&product).Error
    assert.NoError(t, err)
}

func TestCreateSale(t *testing.T) {
    db := setupTestDB()
    assert.NotNil(t, db)

    // Crear cliente necesario para la venta
    client := models.Client{ID: 1, Name: "Juan Lopez", Email: "juanlopez@test.com", Phone: "654321587", Address: "Calle del Test 11"}
    err := db.Create(&client).Error
    assert.NoError(t, err)

    // Crear productos necesarios para la venta
    products := []models.Product{
        {ID: 10, Name: "Producto 10", Description: "Descripción del producto 10", Category: "Categoría A", Price: 10.99, Stock: 20},
        {ID: 9, Name: "Producto 9", Description: "Descripción del producto 9", Category: "Categoría B", Price: 5.99, Stock: 15},
    }
    for _, product := range products {
        err = db.Create(&product).Error
        assert.NoError(t, err)
    }

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simular JSON
    c.Request = httptest.NewRequest(
        "POST",
        "/sales",
        strings.NewReader(`{
            "client_id": 1,
            "date": "2025-04-01",
            "products": [
                {"product_id": 10, "stock": 2},
                {"product_id": 9, "stock": 5}
            ],
            "quantity": 7
        }`),
    )
    c.Request.Header.Set("Content-Type", "application/json")

    handlers.CreateSale(c)
    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetClients(t *testing.T) {
    db := setupTestDB()
    assert.NotNil(t, db)

    // Crear cliente
    client := models.Client{ID: 1, Name: "Juan Lopez", Email: "juanlopez@test.com", Phone: "654321587", Address: "Calle del Test 11"}
    err := db.Create(&client).Error
    assert.NoError(t, err)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simular solicitud
    c.Request = httptest.NewRequest("GET", "/clients", nil)

    handlers.GetAllClients(c)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSales(t *testing.T) {
    db := setupTestDB()
    assert.NotNil(t, db)

    // Crear venta
    sale := models.Sale{ID: 1, ClientID: 1, Date: "2025-04-01", Quantity: 7}
    err := db.Create(&sale).Error
    assert.NoError(t, err)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simular solicitud
    c.Request = httptest.NewRequest("GET", "/sales", nil)

    handlers.GetAllSales(c)
    assert.Equal(t, http.StatusOK, w.Code)
}


