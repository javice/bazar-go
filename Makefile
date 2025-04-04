# Variables
BINARY_NAME = bazar-api
SWAGGER_CMD = swag init -g cmd/main.go

.PHONY: run-backend test swagger migrate run-backend open-frontend build-backend verify-binary  locust all help

# Iniciar servidor
run-backend:
	@echo "Iniciando backend..."
	@cd backend && go run cmd/main.go

# Ejecutar tests
test:
	@echo "Ejecutando pruebas..."
	@cd backend && go test -v ./...

# Generar documentación Swagger
swagger:
	@echo "Generando documentación Swagger..."
	@cd backend/tests && ${SWAGGER_CMD}

# Migrar base de datos
migrate:
	@echo "Ejecutando migraciones de base de datos..."
	@go run cmd/main.go --migrate

# Limpiar binarios
clean:
	@echo "Limpiando binarios..."
	@rm -f backend/build/${BINARY_NAME}

# Instalar dependencias
deps:
	@echo "Descargando dependencias..."
	@cd backend && go mod download

# Construir el binario del backend
build-backend:
	@echo "Construyendo el binario del backend..."
	@cd backend/cmd && go build -v -o build/${BINARY_NAME} cmd/main.go

# Verificar que el binario existe
verify-binary:
	@echo "Verificando que el binario existe..."
	@test -f backend/build/${BINARY_NAME} && echo "El binario existe: backend/build/${BINARY_NAME}" || echo "El binario no existe"

# Ejecutar Locust
locust:
	@echo "Ejecutando Locust..."
	@cd locust && locust -f locustfile.py --host=https://bazar-go.onrender.com

# Ejecutar todo el flujo
all: deps test build-backend verify-binary
	@echo "Flujo completo ejecutado correctamente"

# Mostrar ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  run-backend    Iniciar servidor"
	@echo "  test           Ejecutar pruebas"
	@echo "  swagger        Generar documentación Swagger"
	@echo "  migrate        Migrar base de datos"
	@echo "  clean          Limpiar binarios"
	@echo "  deps           Instalar dependencias"
	@echo "  build-backend  Construir el binario del backend"
	@echo "  verify-binary  Verificar que el binario existe"
	@echo "  all            Ejecutar todo el flujo"
	@echo "  help           Mostrar esta ayuda"