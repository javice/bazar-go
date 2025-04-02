// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/clients": {
            "get": {
                "description": "Listar todos los clientes de la base de datos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clientes"
                ],
                "summary": "Listar clientes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Client"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Crear un nuevo cliente en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clientes"
                ],
                "summary": "Crear un cliente",
                "parameters": [
                    {
                        "description": "Datos del cliente",
                        "name": "sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "400": {
                        "description": "Error de validación",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/clients/{id}": {
            "get": {
                "description": "Obtener un cliente por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clientes"
                ],
                "summary": "Obtener cliente por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Actualizar un cliente existente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clientes"
                ],
                "summary": "Actualizar cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos del cliente",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    }
                }
            },
            "delete": {
                "description": "Eliminamos un cliente existente",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clientes"
                ],
                "summary": "Eliminar cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Obtiene una lista de todos los productos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Productos"
                ],
                "summary": "Listar productos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Crear un nuevo producto en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Productos"
                ],
                "summary": "Crear producto",
                "parameters": [
                    {
                        "description": "Datos del producto",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Error de validación",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Obtenemos un producto por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Productos"
                ],
                "summary": "Obtener producto por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del producto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualizamos un producto por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Productos"
                ],
                "summary": "Actualizar producto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del producto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos del producto",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Eliminamos un producto de la base de datos por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Productos"
                ],
                "summary": "Eliminar producto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del producto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/sales": {
            "get": {
                "description": "Listar todas las ventas de la base de datos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ventas"
                ],
                "summary": "Listar ventas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Sale"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Crear una nueva venta en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ventas"
                ],
                "summary": "Crear una venta",
                "parameters": [
                    {
                        "description": "Datos de la venta",
                        "name": "sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    },
                    "400": {
                        "description": "Error de validación",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/sales/{id}": {
            "get": {
                "description": "Listamos una venta por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ventas"
                ],
                "summary": "Obtener venta por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la venta",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualizar una venta existente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ventas"
                ],
                "summary": "Actualizar venta",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la venta",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos de la venta actualizada",
                        "name": "sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                }
            },
            "delete": {
                "description": "Eliminamos una venta de nuestra base de datos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ventas"
                ],
                "summary": "Eliminar venta",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la venta",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Client": {
            "description": "Modelo de cliente",
            "type": "object",
            "properties": {
                "address": {
                    "description": "Dirección del cliente",
                    "type": "string"
                },
                "client_id": {
                    "description": "ID del cliente",
                    "type": "integer"
                },
                "email": {
                    "description": "Correo electrónico del cliente",
                    "type": "string"
                },
                "name": {
                    "description": "Nombre del cliente",
                    "type": "string"
                },
                "phone": {
                    "description": "Teléfono del cliente",
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "description": "Modelo de producto",
            "type": "object",
            "properties": {
                "category": {
                    "description": "Categoría del producto",
                    "type": "string"
                },
                "description": {
                    "description": "Descripción del producto",
                    "type": "string"
                },
                "name": {
                    "description": "Nombre del producto",
                    "type": "string"
                },
                "price": {
                    "description": "Precio del producto",
                    "type": "number"
                },
                "product_id": {
                    "description": "ID del producto",
                    "type": "integer"
                },
                "stock": {
                    "description": "Stock del producto",
                    "type": "integer"
                }
            }
        },
        "models.Sale": {
            "description": "Modelo de venta",
            "type": "object",
            "properties": {
                "client": {
                    "description": "Cliente asociado a la venta",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Client"
                        }
                    ]
                },
                "client_id": {
                    "description": "ID del cliente",
                    "type": "integer"
                },
                "date": {
                    "description": "Fecha de la venta",
                    "type": "string"
                },
                "products": {
                    "description": "Productos asociados a la venta",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                },
                "quantity": {
                    "description": "Cantidad de productos vendidos",
                    "type": "integer"
                },
                "sale_id": {
                    "type": "integer"
                },
                "total_amount": {
                    "description": "Monto total de la venta",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "bazar-go.onrender.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Bazar API",
	Description:      "API para gestionar productos, clientes y ventas de un bazar.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
