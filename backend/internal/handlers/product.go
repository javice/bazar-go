package handlers

import (
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/javice/bazar-go/backend/internal/models"
    "github.com/javice/bazar-go/backend/internal/repositories"
)

// @Summary Listar productos
// @Description Obtiene una lista de todos los productos
// @Tags Productos
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
    products, err := repositories.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, map[string]interface{}{"data": products})

}

// @Summary Crear producto
// @Description Crear un nuevo producto en la base de datos
// @Tags Productos
// @Accept json
// @Produce json
// @Param product body models.Product true "Datos del producto"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string "Error de validación"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /products [post]
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    if err := repositories.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

// @Summary Obtener producto por ID
// @Description Obtenemos un producto por su ID
// @Tags Productos
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repositories.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Actualizar producto
// @Description Actualizamos un producto por su ID
// @Tags Productos
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Param product body models.Product true "Datos del producto"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)
	if err := repositories.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Eliminar producto
// @Description Eliminamos un producto de la base de datos por su ID
// @Tags Productos
// @Produce json
// @Param id path int true "ID del producto"
// @Success 204
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repositories.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

