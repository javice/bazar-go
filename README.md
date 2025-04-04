# 🛍️ Bazar API + Frontend

Aplicación completa para gestión de un bazar, con backend en **Go (Gin + SQLite)** y frontend sencillo en **HTML/JavaScript**.

## 🌟 Características

- **Backend**:
  - API REST con CRUD para `productos`, `clientes` y `ventas`.
  - Swagger integrado (`/swagger/index.html`).
  - SQLite como base de datos.
- **Frontend**:
  - Interfaz sencilla para listar, agregar productos y crear ventas.
  - Conexión directa a la API.

## 🚀 Instalación  

1. Clona el repositorio:  

   ```bash  
   git clone https://github.com/tu-usuario/bazar-go.git  
   ```

2. Inicia el backend en local:

   ```bash
   cd backend
   go mod download
   make run
   ```

3. Abre el frontend en local:

   Navega a [http://localhost:8080](http://localhost:8080) en tu navegador.

---

## 🏎️ Pruebas de Performance con Locust

   Este proyecto incluye pruebas de rendimiento para la API utilizando Locust, una herramienta de carga basada en Python. Estas pruebas permiten simular múltiples usuarios interactuando con la API para medir su rendimiento bajo diferentes condiciones.

### 📋 Requisitos previos

1. Python: Asegúrate de tener Python 3.7 o superior instalado en tu máquina.
2. Instalar Locust: Instala Locust ejecutando el siguiente comando:

```bash
pip install locust
```

### 🚀 Cómo ejecutar las pruebas de rendimiento

1. Navega al directorio del proyecto:

```bash
cd /ruta/a/bazar-go
```

2. Ejecuta Locust:

```bash
make locust
```

Esto ejecutará Locust utilizando el archivo locustfile.py ubicado en el directorio performance.

3. Accede a la interfaz web de Locust:
   Abre tu navegador y navega a http://localhost:8089.
   Configura el número de usuarios simultáneos y la tasa de inicio (por ejemplo, 10 usuarios con 1 usuario por segundo).
   Haz clic en "Start Swarming" para iniciar las pruebas.

### 🌐 Configuración de entorno

Local:
   Si estás probando en tu entorno local, asegúrate de que el backend esté corriendo en http://localhost:8080.
   Locust usará esta URL como base para las pruebas.
Producción:
   Si deseas probar en producción, asegúrate de cambiar la URL base en el archivo locustfile.py o al ejecutar Locust:

```bash
locust -f performance/locustfile.py --host=https://bazar-go.onrender.com
```

### 📂 Estructura de las pruebas

Las pruebas de rendimiento incluyen operaciones CRUD para las entidades principales de la API:

Productos:

Listar productos (GET /api/v1/products)
Crear un producto (POST /api/v1/products)
Actualizar un producto (PUT /api/v1/products/:id)
Eliminar un producto (DELETE /api/v1/products/:id)

Clientes:

Listar clientes (GET /api/v1/clients)
Crear un cliente (POST /api/v1/clients)
Actualizar un cliente (PUT /api/v1/clients/:id)
Eliminar un cliente (DELETE /api/v1/clients/:id)

Ventas:

Listar ventas (GET /api/v1/sales)
Crear una venta (POST /api/v1/sales)

### 📊 Resultados

Locust genera métricas en tiempo real, incluyendo:

Tiempos de respuesta promedio y percentiles.
Solicitudes por segundo.
Tasas de error.
Al finalizar las pruebas, puedes descargar un informe detallado desde la interfaz web.

### 📝 Notas adicionales

Pruebas en producción:

Ten cuidado al realizar pruebas de carga en producción, ya que podrían afectar a los usuarios reales.
Asegúrate de coordinar con tu equipo antes de realizar estas pruebas.
Datos dinámicos:

Las pruebas están configuradas para almacenar dinámicamente los IDs de productos y clientes creados, asegurando que las operaciones dependientes (como crear ventas) utilicen datos válidos.

## 🌐 Producción

El servicio está desplegado en **Render** y accesible en la siguiente URL:

- **Frontend**: [https://bazar-go.onrender.com](https://bazar-go.onrender.com)
- **Swagger**: [https://bazar-go.onrender.com/swagger/index.html](https://bazar-go.onrender.com/swagger/index.html)

---

## 🧪 Pruebas del Proyecto

Este proyecto incluye pruebas unitarias y de integración para garantizar la calidad del código y el correcto funcionamiento de la API.

### 📋 Requisitos previos para tests

Go instalado: Asegúrate de tener Go instalado en tu máquina.
Dependencias del proyecto: Instala las dependencias necesarias ejecutando:

```bash
make deps
```

### 🚀 Cómo ejecutar las pruebas

Ejecutar todas las pruebas: Ejecuta el siguiente comando para ejecutar todas las pruebas del proyecto:

```bash
make test
```

Ejecutar pruebas específicas: Si deseas ejecutar pruebas específicas, navega al directorio correspondiente y usa el comando:

```bash
go test -v ./ruta/al/paquete
```

Verificar la cobertura de las pruebas: Puedes generar un informe de cobertura ejecutando:

```bash
go test -cover ./...
```

### 📂 Estructura de todas las pruebas

Las pruebas están organizadas en los siguientes paquetes:

internal/handlers:

Pruebas para los controladores de la API (por ejemplo, products, clients, sales).
Verifica que las rutas y las respuestas sean correctas.

internal/repositories:

Pruebas para las operaciones con la base de datos.
Asegura que las consultas y actualizaciones funcionen correctamente.

pkg/config:

Pruebas para la configuración del proyecto (por ejemplo, conexión a la base de datos).

### 🛠️ Ejemplo de un test unitario

Aquí tienes un ejemplo de cómo se ve un test unitario en Go:

```go
package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/javice/bazar-go/backend/internal/handlers"
)

func TestGetProducts(t *testing.T) {
    // Configurar el router de prueba
    router := gin.Default()
    router.GET("/api/v1/products", handlers.GetProducts)

    // Crear una solicitud de prueba
    req, _ := http.NewRequest("GET", "/api/v1/products", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Verificar la respuesta
    if w.Code != http.StatusOK {
        t.Errorf("Se esperaba el código de estado %d, pero se obtuvo %d", http.StatusOK, w.Code)
    }
}
```

### 📝 Notas adicionales

Pruebas en local:

Asegúrate de que la base de datos esté configurada correctamente antes de ejecutar las pruebas.
Puedes usar una base de datos de prueba separada para evitar modificar datos reales.

Pruebas en CI/CD:

Integra las pruebas en tu pipeline de CI/CD para garantizar que el código pase todas las pruebas antes de ser desplegado.


## 📂 Estructura del Proyecto

```
bazar-go/
├── backend/          # Código Go (API)
├── frontend/         # Interfaz web (HTML/JS)
├── Makefile          # Comandos útiles
└── README.md
```

---

## 🔌 Endpoints clave (API)

| Método | Ruta               | Descripción              |
|--------|--------------------|--------------------------|
| GET    | /api/v1/products   | Listar todos los productos |
| POST   | /api/v1/products   | Crear un producto        |
| PUT    | /api/v1/products/:id   | Modificar un producto |
| DELETE | /api/v1/products/:id | Eliminar un producto    |

---

## 🛠️ Contribuir

Si deseas contribuir al proyecto, sigue estos pasos:

1. **Pruebas locales**:
   - Realiza tus pruebas en un entorno local utilizando `localhost`. Esto asegura que los cambios no afecten el servicio en producción.
   - Usa la URL `http://localhost:8080` para acceder al backend y frontend en local.

2. **Pruebas en producción**:
   - **No realices pruebas directamente en el entorno de producción** (`https://bazar-go.onrender.com`), ya que esto podría afectar a los usuarios finales.

3. **Enviar cambios**:
   - Crea un fork del repositorio.
   - Realiza tus cambios en una nueva rama.
   - Envía un pull request con una descripción detallada de los cambios.

---

## 🚀 Despliegue

### Local

Ejecuta el siguiente comando en la carpeta `backend` para iniciar el servidor:

```bash
make run
```

### Producción

El despliegue en producción utiliza Docker y Render. Consulta la [wiki](https://github.com/tu-usuario/bazar-go/wiki) para más detalles.

---

## 📝 Notas adicionales importantes

- **Swagger**:
  - La documentación Swagger está disponible en:
    - Local: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    - Producción: [https://bazar-go.onrender.com/swagger/index.html](https://bazar-go.onrender.com/swagger/index.html)

- **Base de datos**:
  - En local, se utiliza SQLite como base de datos. Asegúrate de que el archivo de base de datos esté correctamente configurado en el entorno local.
