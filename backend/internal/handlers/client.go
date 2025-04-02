package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/javice/bazar-go/backend/internal/models"
    "strconv"
    "github.com/javice/bazar-go/backend/internal/repositories"
    
)

// Mensajes de ERROR comúnes
const (
    InvalidClientIDFormat   = "Formato de ID cliente NO válido"
    ClientNotFoundMessage   = "Cliente NO encontrado"
    ClientDeletedMessage    = "Cliente eliminado correctamente"
    InternalServerErrMsg    = "Error interno del servidor"
)

// @Summary Crear un cliente
// @Description Crear un nuevo cliente en la base de datos
// @Tags Clientes
// @Accept json
// @Produce json
// @Param sale body models.Client true "Datos del cliente"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]string "Error de validación"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /clients [post]
func CreateClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    if err := repositories.CreateClient(&client); err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, client)
}

// @Summary Listar clientes
// @Description Listar todos los clientes de la base de datos
// @Tags Clientes
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /clients [get]
func GetAllClients(c *gin.Context) {
    clients, err := repositories.GetAllClients()
    if err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, map[string]interface{}{"data": clients})
}

// @Summary Obtener cliente por ID
// @Description Obtener un cliente por su ID
// @Tags Clientes
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} models.Product
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /clients/{id} [get]
func GetClientByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repositories.GetClientById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ClientNotFoundMessage})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Actualizar cliente
// @Description Actualizar un cliente existente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Param client body models.Client true "Datos del cliente"
// @Success 200 {object} models.Client
// @Router /clients/{id} [put]
func UpdateClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.ID = uint(id)
	if err := repositories.UpdateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}


// @Summary Eliminar cliente
// @Description Eliminamos un cliente existente
// @Tags Clientes
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 204
// @Router /clients/{id} [delete]
func DeleteClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, err := repositories.DeleteClient(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

