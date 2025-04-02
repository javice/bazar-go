package handlers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/javice/bazar-go/backend/internal/models"
    "github.com/javice/bazar-go/backend/internal/repositories"
)

// @Summary Crear una venta
// @Description Crear una nueva venta en la base de datos
// @Tags Ventas
// @Accept json
// @Produce json
// @Param sale body models.Sale true "Datos de la venta"
// @Success 201 {object} models.Sale
// @Failure 400 {object} map[string]string "Error de validación"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /sales [post]
func CreateSale(c *gin.Context) {
    var sale models.Sale
    if err := c.ShouldBindJSON(&sale); err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    // Calcular el total y actualizar el stock de los productos
    for _, product := range sale.Products {
        // Buscar el producto en la base de datos
        dbProduct, err := repositories.GetProductByID(product.ID)
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
            return
        }

        // Verificar si hay suficiente stock
        if dbProduct.Stock < product.Stock {
            c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Stock insuficiente para el producto: " + dbProduct.Name,
            })
            return
        }

        // Restar del stock
        dbProduct.Stock -= product.Stock

        // Actualizar el producto en la base de datos
        if err := repositories.UpdateProduct(dbProduct); err != nil {
            c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el stock del producto"})
            return
        }

        // Calcular el total de la venta
        sale.TotalAmount += float64(product.Stock) * dbProduct.Price
    }

    // Crear la venta en la base de datos
    if err := repositories.CreateSale(&sale); err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, sale)
}


// @Summary Listar ventas
// @Description Listar todas las ventas de la base de datos
// @Tags Ventas
// @Produce json
// @Success 200 {array} models.Sale
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /sales [get]
func GetAllSales(c *gin.Context) {
	sales, err := repositories.GetAllSales()
    if err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, map[string]interface{}{"data": sales})
}

// @Summary Obtener venta por ID
// @Description Listamos una venta por su ID
// @Tags Ventas
// @Produce json
// @Param id path int true "ID de la venta"
// @Success 200 {object} models.Sale
// @Router /sales/{id} [get]
func GetSaleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repositories.GetSaleById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venta no encontrada"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Actualizar venta
// @Description Actualizar una venta existente
// @Tags Ventas
// @Accept json
// @Produce json
// @Param id path int true "ID de la venta"
// @Param sale body models.Sale true "Datos de la venta actualizada"
// @Success 200 {object} models.Sale
// @Router /sales/{id} [put]
func UpdateSale(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sale models.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sale.ID = uint(id)
	if err := repositories.UpdateSale(&sale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sale)
}



// @Summary Eliminar venta
// @Description Eliminamos una venta de nuestra base de datos
// @Tags Ventas
// @Produce json
// @Param id path int true "ID de la venta"
// @Success 204
// @Router /sales/{id} [delete]
func DeleteSale(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, err := repositories.DeleteSale(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}