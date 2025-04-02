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

2. Inicia el backend:

```bash
cd backend
go mod download
make run
```

3. Abre el frontend:

Navega a http://localhost:8080 en tu navegador.

### 📂 Estructura del Proyecto

```
bazar-go/
├── backend/          # Código Go (API)
├── frontend/         # Interfaz web (HTML/JS)
├── Makefile          # Comandos útiles
└── README.md
```

### 🔌 Endpoints clave (API)

| Método | Ruta               | Descripción              |
|--------|--------------------|--------------------------|
| GET    | /api/v1/products   | Listar todos los productos |
| POST   | /api/v1/products   | Crear un producto        |
| PUT    | /api/v1/products/:id   | Modificar un producto |
| DELETE | /api/v1/products/:id | Eliminar un producto    |

### 🚀 Despliegue

Local: Ejecuta make run en la carpeta backend.
Producción: Usa Docker (ver wiki).
