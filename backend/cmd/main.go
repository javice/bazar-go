
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/javice/bazar-go/backend/docs"  // Importa la documentación generada (¡nueva ruta!)
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/javice/bazar-go/backend/internal/routes"
	"github.com/javice/bazar-go/backend/pkg/config"
)

// @title Bazar API
// @version 1.0
// @description API para gestionar productos, clientes y ventas de un bazar.
// @host bazar-go.onrender.com
// @BasePath /api/v1
// @schemes https
func main() {
	// Inicializa la base de datos
	config.InitDB()
	defer config.CloseDB()

	// Configura el router de Gin
	r := gin.Default()

	// Middleware CORS (para desarrollo)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Next()
	})

	// Servir archivos estáticos del frontend (desde la raíz del proyecto)
	frontendPath := filepath.Join("/app", "frontend")
	r.Static("/static", filepath.Join(frontendPath, "static")) // CSS/JS
	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(frontendPath, "index.html"))
	})

	// Configuración de Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"), // Ruta al JSON generado
	))

	// Registra las rutas de la API
	routes.SetupRoutes(r)

	// Configuración del servidor HTTP
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    // Canal para capturar señales del sistema
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

    // Goroutine para iniciar el servidor
    go func() {
        log.Println("Servidor iniciado en http://localhost:8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error al iniciar el servidor: %v", err)
        }
    }()

    // Esperar señal de cierre
    <-quit
    log.Println("Iniciando el apagado del servidor...")

    // Contexto con timeout para el apagado
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Apagar el servidor de forma limpia
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Error al apagar el servidor: %v", err)
    }

    log.Println("Servidor detenido correctamente")

	// Inicia el servidor
	/* log.Println("Servidor iniciado en http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	} */
}